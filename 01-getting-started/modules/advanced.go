package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// This example demonstrates advanced module features
// including replace directives, workspace, and vendor

func main() {
	fmt.Println("=== Advanced Go Modules Features ===")
	
	// Demonstrate replace directives
	demonstrateReplaceDirectives()
	
	// Show workspace usage
	demonstrateWorkspace()
	
	// Demonstrate vendor
	demonstrateVendor()
	
	// Show proxy configuration
	demonstrateProxy()
}

// demonstrateReplaceDirectives shows how to use replace directives
func demonstrateReplaceDirectives() {
	fmt.Println("\n1. Replace Directives:")
	fmt.Println("   Replace directives allow you to substitute dependencies:")
	fmt.Println("")
	fmt.Println("   # Replace with local path:")
	fmt.Println("   replace github.com/example/package => ./local/package")
	fmt.Println("")
	fmt.Println("   # Replace with different version:")
	fmt.Println("   replace github.com/example/package => github.com/example/package v1.2.3")
	fmt.Println("")
	fmt.Println("   # Replace with fork:")
	fmt.Println("   replace github.com/example/package => github.com/myfork/package v1.2.3")
	fmt.Println("")
	fmt.Println("   Common use cases:")
	fmt.Println("   - Local development")
	fmt.Println("   - Testing patches")
	fmt.Println("   - Using forks")
}

// demonstrateWorkspace shows how to use Go workspaces
func demonstrateWorkspace() {
	fmt.Println("\n2. Go Workspaces:")
	fmt.Println("   Workspaces allow you to work with multiple modules:")
	fmt.Println("")
	fmt.Println("   # Initialize workspace:")
	fmt.Println("   go work init")
	fmt.Println("")
	fmt.Println("   # Add modules to workspace:")
	fmt.Println("   go work use ./module1")
	fmt.Println("   go work use ./module2")
	fmt.Println("")
	fmt.Println("   # List workspace modules:")
	fmt.Println("   go work use")
	fmt.Println("")
	fmt.Println("   Benefits:")
	fmt.Println("   - Work on multiple related modules")
	fmt.Println("   - Shared dependencies")
	fmt.Println("   - Simplified development workflow")
}

// demonstrateVendor shows how to use vendor directory
func demonstrateVendor() {
	fmt.Println("\n3. Vendor Directory:")
	fmt.Println("   Vendor allows you to include dependencies in your repository:")
	fmt.Println("")
	fmt.Println("   # Create vendor directory:")
	fmt.Println("   go mod vendor")
	fmt.Println("")
	fmt.Println("   # Build with vendored dependencies:")
	fmt.Println("   go build -mod=vendor")
	fmt.Println("")
	fmt.Println("   # Test with vendored dependencies:")
	fmt.Println("   go test -mod=vendor")
	fmt.Println("")
	fmt.Println("   When to use vendor:")
	fmt.Println("   - Offline builds")
	fmt.Println("   - Reproducible builds")
	fmt.Println("   - Corporate environments")
	fmt.Println("")
	fmt.Println("   Note: Vendor directory can be large, consider .gitignore")
}

// demonstrateProxy shows proxy configuration
func demonstrateProxy() {
	fmt.Println("\n4. Module Proxy Configuration:")
	fmt.Println("   Configure where Go downloads modules:")
	fmt.Println("")
	
	// Show current proxy settings
	if output, err := exec.Command("go", "env", "GOPROXY").Output(); err == nil {
		proxy := strings.TrimSpace(string(output))
		fmt.Printf("   Current GOPROXY: %s\n", proxy)
	}
	
	if output, err := exec.Command("go", "env", "GOSUMDB").Output(); err == nil {
		sumdb := strings.TrimSpace(string(output))
		fmt.Printf("   Current GOSUMDB: %s\n", sumdb)
	}
	
	fmt.Println("")
	fmt.Println("   Common proxy configurations:")
	fmt.Println("   # Use default proxy:")
	fmt.Println("   go env -w GOPROXY=https://proxy.golang.org,direct")
	fmt.Println("")
	fmt.Println("   # Use private proxy:")
	fmt.Println("   go env -w GOPROXY=https://proxy.company.com,https://proxy.golang.org,direct")
	fmt.Println("")
	fmt.Println("   # Disable proxy (direct only):")
	fmt.Println("   go env -w GOPROXY=direct")
	fmt.Println("")
	fmt.Println("   # Disable checksum database:")
	fmt.Println("   go env -w GOSUMDB=off")
}

// Additional utility functions for module management
func showModuleHelp() {
	fmt.Println("\n=== Module Management Help ===")
	fmt.Println("")
	fmt.Println("Essential commands:")
	fmt.Println("  go mod init <module-name>     - Initialize new module")
	fmt.Println("  go mod tidy                   - Clean up dependencies")
	fmt.Println("  go get <package>              - Add dependency")
	fmt.Println("  go get -u                     - Update all dependencies")
	fmt.Println("  go mod download               - Download dependencies")
	fmt.Println("  go mod verify                 - Verify dependencies")
	fmt.Println("  go mod why <package>          - Show why dependency is needed")
	fmt.Println("  go mod graph                  - Show dependency graph")
	fmt.Println("")
	fmt.Println("Environment variables:")
	fmt.Println("  GOPROXY                       - Module proxy URL")
	fmt.Println("  GOSUMDB                       - Checksum database URL")
	fmt.Println("  GOPRIVATE                     - Private module patterns")
	fmt.Println("  GONOPROXY                     - Modules to not use proxy for")
	fmt.Println("  GONOSUMDB                     - Modules to not use sumdb for")
}
