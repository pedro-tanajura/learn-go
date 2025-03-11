package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Premium  bool   `json:"premium"`
}

type ShortURL struct {
	gorm.Model
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code" gorm:"unique"`
	UserID      uint      `json:"user_id"`
	ClickCount  int       `json:"click_count"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// Auth request structures
type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type URLInput struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

// JWT claims structure
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// JWT secret key - in production, use environment variables
var jwtKey = []byte("your_secret_key")

// Hash password before saving
func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// Compare password with stored hash
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// Generate JWT token for a user
func generateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

// Register a new user
func registerHandler(c *gin.Context, db *gorm.DB) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	var existingUser User
	result := db.Where("email = ?", input.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Create user
	user := User{
		Email:    input.Email,
		Password: input.Password,
		Premium:  false, // Default to free tier
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate token
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"token":   token,
		"user": gin.H{
			"id":      user.ID,
			"email":   user.Email,
			"premium": user.Premium,
		},
	})
}

// Login existing user
func loginHandler(c *gin.Context, db *gorm.DB) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	var user User
	result := db.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if !user.CheckPassword(input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate token
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":      user.ID,
			"email":   user.Email,
			"premium": user.Premium,
		},
	})
}

// Get current user profile
func getUserProfile(c *gin.Context, db *gorm.DB) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":      user.ID,
			"email":   user.Email,
			"premium": user.Premium,
		},
	})
}

func createShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	// Create a byte slice to hold the result
	code := make([]byte, codeLength)

	// Fill with random characters
	for i := range code {
		// Use crypto/rand for better randomness in production
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

// Register a new URL
func registerURL(c *gin.Context, db *gorm.DB) {
	var input URLInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate URL format
	_, err := url.ParseRequestURI(input.OriginalURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	// Get user ID from context (set by auth middleware)
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user User
	if err := db.First(&user, userIDInterface).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var expirationTime time.Time
	if user.Premium {
		expirationTime = time.Now().Add(365 * 24 * time.Hour) // 1 year for premium
	} else {
		expirationTime = time.Now().Add(24 * time.Hour) // 1 day for free users
	}

	// Create shorten URL
	shortenURL := ShortURL{
		OriginalURL: input.OriginalURL,
		ShortCode:   "",
		UserID:      userID,
		ClickCount:  0,
		ExpiresAt:   expirationTime,
	}

	// Check for collisions and regenerate if needed
	for {
		shortCode := createShortCode()
		var existing ShortURL
		result := db.Where("short_code = ?", shortCode).First(&existing)
		if result.RowsAffected == 0 {
			// No collision, use this code
			shortenURL.ShortCode = shortCode
			break
		}
		// Collision found, try again
	}

	if err := db.Create(&shortenURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shorten URL"})
		return
	}

	baseURL := "http://yourdomain.com/" // Configure this appropriately
	c.JSON(http.StatusCreated, gin.H{
		"message":     "URL Registration successful",
		"shortURL":    baseURL + shortenURL.ShortCode,
		"originalURL": shortenURL.OriginalURL,
		"expiresAt":   shortenURL.ExpiresAt,
	})
}

// Get user URLs
func getURLs(c *gin.Context, db *gorm.DB) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var userURLs []ShortURL
	err := db.Where("user_id = ?", user.ID).Find(&userURLs)

	if err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	urlsData := make([]gin.H, len(userURLs))
	for i, url := range userURLs {
		urlsData[i] = gin.H{
			"id":           url.ID,
			"original_url": url.OriginalURL,
			"short_code":   url.ShortCode,
			"click_count":  url.ClickCount,
			"expires_at":   url.ExpiresAt,
			"created_at":   url.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"urls":  urlsData,
		"count": len(urlsData),
	})
}

// Redirect to original URL
func redirectToOriginalURL(c *gin.Context, db *gorm.DB) {
	// Get the code from URL parameter
	code := c.Param("code")

	// Find the URL in the database
	var shortURL ShortURL
	result := db.Where("short_code = ?", code).First(&shortURL)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	// Check if the URL has expired
	if time.Now().After(shortURL.ExpiresAt) {
		c.JSON(http.StatusGone, gin.H{"error": "This short URL has expired"})
		return
	}

	// Increment the click count
	db.Model(&shortURL).Update("click_count", shortURL.ClickCount+1)

	// Redirect to the original URL
	c.Redirect(http.StatusFound, shortURL.OriginalURL)
}

// Get user URLs analytics by Code
func getURLAnalyticsByCode(c *gin.Context, db *gorm.DB) {
	code := c.Param("code")

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var shortURL ShortURL
	result := db.Where("user_id = ? AND short_code = ?", user.ID, code).First(&shortURL)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Check if the URL has expired
	if time.Now().After(shortURL.ExpiresAt) {
		c.JSON(http.StatusGone, gin.H{"error": "This short URL has expired"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           shortURL.ID,
		"original_url": shortURL.OriginalURL,
		"short_code":   shortURL.ShortCode,
		"click_count":  shortURL.ClickCount,
		"expires_at":   shortURL.ExpiresAt,
		"created_at":   shortURL.CreatedAt,
	})
}

// Auth middleware to protect routes
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := authHeader[7:]

		// Parse and validate token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user ID in context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("./data/url_shortener.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Auto migrate models
	db.AutoMigrate(&User{}, &ShortURL{})

	// Initialize Gin router
	r := gin.Default()

	// User management routes
	r.POST("/register", func(c *gin.Context) {
		registerHandler(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		loginHandler(c, db)
	})

	// Protected routes
	auth := r.Group("/")
	auth.Use(authMiddleware())
	{
		auth.GET("/me", func(c *gin.Context) {
			getUserProfile(c, db)
		})

		auth.POST("/urls", func(c *gin.Context) {
			registerURL(c, db)
		})

		auth.GET("/urls", func(c *gin.Context) {
			getURLs(c, db)
		})

		auth.GET("/urls/:code", func(c *gin.Context) {
			redirectToOriginalURL(c, db)
		})

		auth.GET("/urls/:code/analytics", func(c *gin.Context) {
			getURLAnalyticsByCode(c, db)
		})
	}

	r.Run(":8080")
}
