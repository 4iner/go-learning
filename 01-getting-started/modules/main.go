package main

import (
	"fmt"
	"log"
	"os"
)

// This example demonstrates Go modules in action
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Modules Example ===")
	
	// Demonstrate module information
	showModuleInfo()
	
	// Demonstrate dependency usage
	demonstrateDependencies()
	
	// Show environment variables
	showGoEnv()
}

// showModuleInfo displays information about the current module
func showModuleInfo() {
	fmt.Println("\n1. Module Information:")
	fmt.Printf("   Current working directory: %s\n", getCurrentDir())
	fmt.Println("   This module is defined in go.mod")
	fmt.Println("   Dependencies are managed automatically")
}

// demonstrateDependencies shows how dependencies work
func demonstrateDependencies() {
	fmt.Println("\n2. Dependency Management:")
	fmt.Println("   - Dependencies are declared in go.mod")
	fmt.Println("   - Checksums are stored in go.sum")
	fmt.Println("   - Use 'go get' to add dependencies")
	fmt.Println("   - Use 'go mod tidy' to clean up")
	
	// Example of how you would add a dependency:
	fmt.Println("\n   Example commands:")
	fmt.Println("   go get github.com/gin-gonic/gin")
	fmt.Println("   go get github.com/lib/pq@v1.10.9")
	fmt.Println("   go mod tidy")
}

// showGoEnv displays relevant Go environment variables
func showGoEnv() {
	fmt.Println("\n3. Go Environment Variables:")
	
	// Get Go environment variables
	gopath := os.Getenv("GOPATH")
	goproxy := os.Getenv("GOPROXY")
	gosumdb := os.Getenv("GOSUMDB")
	
	fmt.Printf("   GOPATH: %s\n", gopath)
	fmt.Printf("   GOPROXY: %s\n", goproxy)
	fmt.Printf("   GOSUMDB: %s\n", gosumdb)
	
	fmt.Println("\n   Note: With modules, GOPATH is less important")
	fmt.Println("   Modules store dependencies in $GOPATH/pkg/mod")
}

// getCurrentDir returns the current working directory
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
