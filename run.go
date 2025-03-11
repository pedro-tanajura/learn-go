package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run run.go <program-path>")
		fmt.Println("Example: go run run.go data-structures/stack")
		fmt.Println("Example: go run run.go hello")
		os.Exit(1)
	}

	programPath := os.Args[1]

	repoRoot, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("Error getting repository root: %v\n", err)
		os.Exit(1)
	}

	fullPath := filepath.Join(repoRoot, programPath)

	info, err := os.Stat(fullPath)
	if err != nil || !info.IsDir() {
		fmt.Printf("Error: %s is not a valid directory\n", programPath)
		os.Exit(1)
	}

	goModPath := filepath.Join(fullPath, "go.mod")
	_, err = os.Stat(goModPath)
	hasGoMod := err == nil

	files, err := filepath.Glob(filepath.Join(fullPath, "*.go"))
	if err != nil || len(files) == 0 {
		fmt.Printf("Error: No Go files found in %s\n", programPath)
		os.Exit(1)
	}

	fmt.Printf("Running program in %s\n", programPath)

	var cmd *exec.Cmd
	if hasGoMod {

		cmd = exec.Command("go", "run", ".")
		cmd.Dir = fullPath
	} else {

		args := []string{"run"}
		for _, file := range files {
			args = append(args, file)
		}
		cmd = exec.Command("go", args...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
