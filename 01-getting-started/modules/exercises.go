package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Module Exercises - Hands-on practice with Go modules
// Complete these exercises to master Go module management

func main() {
	fmt.Println("=== Go Modules Exercises ===")
	fmt.Println("Complete these exercises to practice Go module management:")
	fmt.Println("")
	
	showExercises()
	
	// Uncomment the exercise you want to run
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
}

// showExercises displays all available exercises
func showExercises() {
	exercises := []struct {
		number int
		title string
		description string
	}{
		{1, "Initialize a New Module", "Create a new Go module and explore its structure"},
		{2, "Add Dependencies", "Add external dependencies and understand version management"},
		{3, "Replace Directives", "Use replace directives for local development"},
		{4, "Module Workspace", "Create a workspace with multiple modules"},
	}
	
	for _, ex := range exercises {
		fmt.Printf("Exercise %d: %s\n", ex.number, ex.title)
		fmt.Printf("  %s\n", ex.description)
		fmt.Println("")
	}
}

// exercise1: Initialize a new module
func exercise1() {
	fmt.Println("=== Exercise 1: Initialize a New Module ===")
	
	// Create a temporary directory for this exercise
	tempDir := "exercise1-temp"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)
	
	// Change to temp directory
	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	fmt.Println("1. Creating a new module...")
	
	// Initialize module
	cmd := exec.Command("go", "mod", "init", "github.com/learner/exercise1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error initializing module: %v\n", err)
		return
	}
	fmt.Printf("Output: %s\n", string(output))
	
	// Show created files
	fmt.Println("\n2. Created files:")
	files, _ := os.ReadDir(".")
	for _, file := range files {
		fmt.Printf("   %s\n", file.Name())
	}
	
	// Show go.mod content
	fmt.Println("\n3. go.mod content:")
	content, _ := os.ReadFile("go.mod")
	fmt.Printf("%s\n", string(content))
	
	fmt.Println("\n4. Next steps:")
	fmt.Println("   - Add some Go code")
	fmt.Println("   - Add dependencies with 'go get'")
	fmt.Println("   - Run 'go mod tidy' to clean up")
}

// exercise2: Add dependencies
func exercise2() {
	fmt.Println("=== Exercise 2: Add Dependencies ===")
	
	// Create a temporary directory
	tempDir := "exercise2-temp"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fmt.Printf("Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)
	
	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)
	
	fmt.Println("1. Initializing module...")
	exec.Command("go", "mod", "init", "github.com/learner/exercise2").Run()
	
	// Create a simple Go file that uses external dependencies
	goFile := `package main

import (
	"fmt"
	"strings"
)

func main() {
	// This will require the strings package (built-in)
	text := "Hello, Go Modules!"
	upper := strings.ToUpper(text)
	fmt.Println(upper)
}
`
	
	if err := os.WriteFile("main.go", []byte(goFile), 0644); err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		return
	}
	
	fmt.Println("2. Created main.go with string manipulation")
	
	fmt.Println("\n3. Running 'go mod tidy'...")
	cmd := exec.Command("go", "mod", "tidy")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running go mod tidy: %v\n", err)
		return
	}
	fmt.Printf("Output: %s\n", string(output))
	
	fmt.Println("\n4. Current module state:")
	content, _ := os.ReadFile("go.mod")
	fmt.Printf("go.mod:\n%s\n", string(content))
	
	if _, err := os.Stat("go.sum"); err == nil {
		fmt.Println("go.sum file created (no external dependencies in this example)")
	}
	
	fmt.Println("\n5. Try adding external dependencies:")
	fmt.Println("   go get github.com/gin-gonic/gin")
	fmt.Println("   go get github.com/lib/pq")
	fmt.Println("   go mod tidy")
}

// exercise3: Replace directives
func exercise3() {
	fmt.Println("=== Exercise 3: Replace Directives ===")
	
	fmt.Println("Replace directives are useful for:")
	fmt.Println("1. Local development")
	fmt.Println("2. Testing patches")
	fmt.Println("3. Using forks")
	fmt.Println("")
	
	fmt.Println("Example go.mod with replace directive:")
	exampleMod := `module github.com/learner/exercise3

go 1.21

require (
	github.com/example/package v1.0.0
)

// Replace with local development version
replace github.com/example/package => ./local/package

// Replace with specific version
replace github.com/example/package => github.com/example/package v1.2.3

// Replace with fork
replace github.com/example/package => github.com/myfork/package v1.0.0
`
	
	fmt.Println(exampleMod)
	
	fmt.Println("Commands to try:")
	fmt.Println("1. Create a local package directory")
	fmt.Println("2. Add replace directive to go.mod")
	fmt.Println("3. Run 'go mod tidy'")
	fmt.Println("4. Build and test your application")
}

// exercise4: Module workspace
func exercise4() {
	fmt.Println("=== Exercise 4: Module Workspace ===")
	
	fmt.Println("Creating a workspace with multiple modules...")
	
	// Create workspace directory
	workspaceDir := "exercise4-workspace"
	if err := os.MkdirAll(workspaceDir, 0755); err != nil {
		fmt.Printf("Error creating workspace: %v\n", err)
		return
	}
	defer os.RemoveAll(workspaceDir)
	
	originalDir, _ := os.Getwd()
	os.Chdir(workspaceDir)
	defer os.Chdir(originalDir)
	
	fmt.Println("1. Initializing workspace...")
	cmd := exec.Command("go", "work", "init")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error initializing workspace: %v\n", err)
		return
	}
	fmt.Printf("Output: %s\n", string(output))
	
	// Create module1
	os.MkdirAll("module1", 0755)
	os.Chdir("module1")
	exec.Command("go", "mod", "init", "github.com/learner/module1").Run()
	os.WriteFile("main.go", []byte(`package main

import "fmt"

func main() {
	fmt.Println("Module 1")
}
`), 0644)
	os.Chdir("..")
	
	// Create module2
	os.MkdirAll("module2", 0755)
	os.Chdir("module2")
	exec.Command("go", "mod", "init", "github.com/learner/module2").Run()
	os.WriteFile("main.go", []byte(`package main

import "fmt"

func main() {
	fmt.Println("Module 2")
}
`), 0644)
	os.Chdir("..")
	
	// Add modules to workspace
	fmt.Println("\n2. Adding modules to workspace...")
	exec.Command("go", "work", "use", "./module1").Run()
	exec.Command("go", "work", "use", "./module2").Run()
	
	// Show workspace status
	fmt.Println("\n3. Workspace status:")
	cmd = exec.Command("go", "work", "use")
	output, err = cmd.CombinedOutput()
	if err == nil {
		fmt.Printf("Modules in workspace:\n%s\n", string(output))
	}
	
	// Show go.work file
	fmt.Println("4. go.work file content:")
	content, _ := os.ReadFile("go.work")
	fmt.Printf("%s\n", string(content))
	
	fmt.Println("\n5. Benefits of workspace:")
	fmt.Println("   - Work on multiple related modules")
	fmt.Println("   - Shared dependencies")
	fmt.Println("   - Simplified development workflow")
	fmt.Println("   - No need to publish modules for local development")
}

// Utility function to show module help
func showModuleHelp() {
	fmt.Println("\n=== Quick Reference ===")
	fmt.Println("")
	fmt.Println("Essential Commands:")
	fmt.Println("  go mod init <name>           - Initialize module")
	fmt.Println("  go mod tidy                  - Clean dependencies")
	fmt.Println("  go get <package>             - Add dependency")
	fmt.Println("  go get -u                    - Update dependencies")
	fmt.Println("  go mod download              - Download dependencies")
	fmt.Println("  go mod verify                - Verify dependencies")
	fmt.Println("  go mod why <package>         - Show dependency reason")
	fmt.Println("  go mod graph                  - Show dependency graph")
	fmt.Println("")
	fmt.Println("Workspace Commands:")
	fmt.Println("  go work init                 - Initialize workspace")
	fmt.Println("  go work use <path>           - Add module to workspace")
	fmt.Println("  go work use                  - List workspace modules")
	fmt.Println("")
	fmt.Println("Environment Variables:")
	fmt.Println("  GOPROXY                      - Module proxy URL")
	fmt.Println("  GOSUMDB                      - Checksum database")
	fmt.Println("  GOPRIVATE                    - Private module patterns")
}
