package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// This example shows how to work with Go modules programmatically
// It demonstrates module commands and dependency management

func main() {
	fmt.Println("=== Go Modules Commands Example ===")
	
	// Show current module status
	showModuleStatus()
	
	// Demonstrate module commands
	demonstrateModuleCommands()
	
	// Show dependency information
	showDependencyInfo()
}

// showModuleStatus displays the current module status
func showModuleStatus() {
	fmt.Println("\n1. Current Module Status:")
	
	// Check if go.mod exists
	if _, err := os.Stat("go.mod"); err == nil {
		fmt.Println("   ✓ go.mod file exists")
		
		// Read and display go.mod content
		content, err := os.ReadFile("go.mod")
		if err == nil {
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) != "" {
					fmt.Printf("   %s\n", line)
				}
			}
		}
	} else {
		fmt.Println("   ✗ go.mod file not found")
	}
	
	// Check if go.sum exists
	if _, err := os.Stat("go.sum"); err == nil {
		fmt.Println("   ✓ go.sum file exists")
	} else {
		fmt.Println("   ✗ go.sum file not found")
	}
}

// demonstrateModuleCommands shows common module commands
func demonstrateModuleCommands() {
	fmt.Println("\n2. Common Module Commands:")
	
	commands := []struct {
		command string
		description string
	}{
		{"go mod init <module-name>", "Initialize a new module"},
		{"go mod tidy", "Add missing and remove unused modules"},
		{"go get <package>", "Add a dependency"},
		{"go get <package>@<version>", "Add specific version"},
		{"go get -u <package>", "Update dependency"},
		{"go get -u", "Update all dependencies"},
		{"go mod download", "Download dependencies"},
		{"go mod verify", "Verify dependencies"},
		{"go mod why <package>", "Show why a dependency is needed"},
		{"go mod graph", "Print module requirement graph"},
	}
	
	for _, cmd := range commands {
		fmt.Printf("   %-30s - %s\n", cmd.command, cmd.description)
	}
}

// showDependencyInfo displays information about dependencies
func showDependencyInfo() {
	fmt.Println("\n3. Dependency Information:")
	
	// Try to get module graph
	if output, err := exec.Command("go", "mod", "graph").Output(); err == nil {
		lines := strings.Split(string(output), "\n")
		if len(lines) > 1 && lines[0] != "" {
			fmt.Println("   Module dependency graph:")
			for i, line := range lines {
				if i >= 5 { // Limit output
					fmt.Println("   ... (truncated)")
					break
				}
				if strings.TrimSpace(line) != "" {
					fmt.Printf("   %s\n", line)
				}
			}
		} else {
			fmt.Println("   No dependencies found")
		}
	} else {
		fmt.Println("   Could not retrieve dependency graph")
	}
	
	fmt.Println("\n   To see all dependencies:")
	fmt.Println("   go mod graph | head -20")
	fmt.Println("\n   To see why a dependency is needed:")
	fmt.Println("   go mod why <package-name>")
}
