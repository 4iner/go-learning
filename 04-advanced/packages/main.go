package main

import (
	"fmt"
	"yourproject/pkg/math"
	"yourproject/pkg/user"
)

// This example demonstrates Go's package system
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Packages Examples ===")
	
	// Demonstrate package usage
	demonstratePackageUsage()
	
	// Demonstrate package visibility
	demonstratePackageVisibility()
	
	// Demonstrate package initialization
	demonstratePackageInitialization()
}

func demonstratePackageUsage() {
	fmt.Println("\n1. Package Usage:")
	
	// Use math package
	result := math.Add(2, 3)
	fmt.Printf("   math.Add(2, 3) = %d\n", result)
	
	result = math.Subtract(5, 3)
	fmt.Printf("   math.Subtract(5, 3) = %d\n", result)
	
	// Use user package
	u := user.New("Alice", 30)
	fmt.Printf("   user: %+v\n", u)
	
	name := u.GetName()
	fmt.Printf("   user name: %s\n", name)
}

func demonstratePackageVisibility() {
	fmt.Println("\n2. Package Visibility:")
	
	// Can access exported functions
	result := math.Add(2, 3)
	fmt.Printf("   Exported function result: %d\n", result)
	
	// Cannot access unexported functions
	// result := math.add(2, 3)  // Compile error
	
	// Can access exported variables
	fmt.Printf("   Exported variable: %f\n", math.Pi)
	
	// Cannot access unexported variables
	// fmt.Printf("Unexported variable: %f\n", math.pi)  // Compile error
}

func demonstratePackageInitialization() {
	fmt.Println("\n3. Package Initialization:")
	
	// Package initialization happens automatically
	fmt.Println("   Packages are initialized automatically")
	fmt.Printf("   Config loaded: %+v\n", math.GetConfig())
}
