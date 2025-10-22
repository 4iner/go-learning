package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// This example demonstrates Go's control flow structures
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Control Flow Examples ===")
	
	// Demonstrate if statements
	demonstrateIfStatements()
	
	// Demonstrate switch statements
	demonstrateSwitchStatements()
	
	// Demonstrate for loops
	demonstrateForLoops()
	
	// Demonstrate range loops
	demonstrateRangeLoops()
	
	// Demonstrate defer statements
	demonstrateDeferStatements()
	
	// Demonstrate panic and recover
	demonstratePanicRecover()
}

// demonstrateIfStatements shows different if statement patterns
func demonstrateIfStatements() {
	fmt.Println("\n1. If Statements:")
	
	// Basic if statement
	age := 25
	if age >= 18 {
		fmt.Printf("   Age %d: Adult\n", age)
	} else {
		fmt.Printf("   Age %d: Minor\n", age)
	}
	
	// If with initialization
	if err := checkFile("nonexistent.txt"); err != nil {
		fmt.Printf("   File check failed: %v\n", err)
	}
	
	// If-else if-else chain
	score := 85
	if score >= 90 {
		fmt.Printf("   Score %d: Grade A\n", score)
	} else if score >= 80 {
		fmt.Printf("   Score %d: Grade B\n", score)
	} else if score >= 70 {
		fmt.Printf("   Score %d: Grade C\n", score)
	} else {
		fmt.Printf("   Score %d: Grade F\n", score)
	}
	
	// If with short variable declaration
	if name := getName(); name != "" {
		fmt.Printf("   Hello, %s!\n", name)
	} else {
		fmt.Println("   No name provided")
	}
	
	// Multiple conditions
	x, y := 10, 20
	if x > 0 && y > 0 {
		fmt.Printf("   Both %d and %d are positive\n", x, y)
	}
	
	if x > 0 || y < 0 {
		fmt.Printf("   At least one of %d and %d is positive\n", x, y)
	}
}

// demonstrateSwitchStatements shows different switch patterns
func demonstrateSwitchStatements() {
	fmt.Println("\n2. Switch Statements:")
	
	// Basic switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("   Monday: Start of work week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("   Mid week")
	case "Friday":
		fmt.Println("   Friday: TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("   Weekend!")
	default:
		fmt.Println("   Invalid day")
	}
	
	// Switch with initialization
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("   Running on macOS")
	case "linux":
		fmt.Println("   Running on Linux")
	case "windows":
		fmt.Println("   Running on Windows")
	default:
		fmt.Printf("   Running on %s\n", os)
	}
	
	// Switch without expression (like if-else chain)
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("   Temperature: Freezing")
	case temperature < 10:
		fmt.Println("   Temperature: Cold")
	case temperature < 20:
		fmt.Println("   Temperature: Cool")
	case temperature < 30:
		fmt.Println("   Temperature: Warm")
	default:
		fmt.Println("   Temperature: Hot")
	}
	
	// Type switch
	demonstrateTypeSwitch()
	
	// Switch with fallthrough
	value := 1
	switch value {
	case 1:
		fmt.Println("   Value is 1")
		fallthrough
	case 2:
		fmt.Println("   Value is 1 or 2")
	case 3:
		fmt.Println("   Value is 3")
	}
}

// demonstrateTypeSwitch shows type switching
func demonstrateTypeSwitch() {
	fmt.Println("\n   Type Switch Examples:")
	
	// Type switch for interface{} values
	values := []interface{}{42, "hello", true, 3.14, []int{1, 2, 3}}
	
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Printf("     Integer: %d\n", v)
		case string:
			fmt.Printf("     String: %s\n", v)
		case bool:
			fmt.Printf("     Boolean: %t\n", v)
		case float64:
			fmt.Printf("     Float: %.2f\n", v)
		case []int:
			fmt.Printf("     Slice: %v\n", v)
		default:
			fmt.Printf("     Unknown type: %T\n", v)
		}
	}
}

// demonstrateForLoops shows different for loop patterns
func demonstrateForLoops() {
	fmt.Println("\n3. For Loops:")
	
	// Traditional for loop
	fmt.Println("   Traditional for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("     i = %d\n", i)
	}
	
	// Multiple variables
	fmt.Println("   Multiple variables:")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("     i = %d, j = %d\n", i, j)
	}
	
	// While-style loop
	fmt.Println("   While-style loop:")
	i := 0
	for i < 3 {
		fmt.Printf("     i = %d\n", i)
		i++
	}
	
	// Loop with break
	fmt.Println("   Loop with break:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("     Breaking at 5")
			break
		}
		fmt.Printf("     i = %d\n", i)
	}
	
	// Loop with continue
	fmt.Println("   Loop with continue (odd numbers only):")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("     i = %d\n", i)
	}
	
	// Labeled break
	fmt.Println("   Labeled break:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("     Breaking outer loop")
				break outer
			}
			fmt.Printf("     i = %d, j = %d\n", i, j)
		}
	}
}

// demonstrateRangeLoops shows range loop patterns
func demonstrateRangeLoops() {
	fmt.Println("\n4. Range Loops:")
	
	// Range over slice
	fmt.Println("   Range over slice:")
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("     Index: %d, Value: %s\n", index, value)
	}
	
	// Range over slice (value only)
	fmt.Println("   Range over slice (value only):")
	for _, value := range slice {
		fmt.Printf("     Value: %s\n", value)
	}
	
	// Range over slice (index only)
	fmt.Println("   Range over slice (index only):")
	for index := range slice {
		fmt.Printf("     Index: %d\n", index)
	}
	
	// Range over map
	fmt.Println("   Range over map:")
	m := map[string]int{"apple": 5, "banana": 3, "cherry": 8}
	for key, value := range m {
		fmt.Printf("     Key: %s, Value: %d\n", key, value)
	}
	
	// Range over string
	fmt.Println("   Range over string:")
	s := "Hello"
	for index, char := range s {
		fmt.Printf("     Index: %d, Char: %c\n", index, char)
	}
	
	// Range over string (rune by rune)
	fmt.Println("   Range over string (rune by rune):")
	s2 := "Hello, 世界"
	for index, rune := range s2 {
		fmt.Printf("     Index: %d, Rune: %c\n", index, rune)
	}
}

// demonstrateDeferStatements shows defer patterns
func demonstrateDeferStatements() {
	fmt.Println("\n5. Defer Statements:")
	
	// Basic defer
	fmt.Println("   Basic defer:")
	deferExample1()
	
	// Defer with arguments
	fmt.Println("   Defer with arguments:")
	deferExample2()
	
	// Multiple defer statements
	fmt.Println("   Multiple defer statements:")
	deferExample3()
	
	// Defer with anonymous function
	fmt.Println("   Defer with anonymous function:")
	deferExample4()
	
	// Defer for cleanup
	fmt.Println("   Defer for cleanup:")
	deferCleanupExample()
}

// deferExample1 demonstrates basic defer
func deferExample1() {
	defer fmt.Println("     This will be printed last")
	fmt.Println("     This will be printed first")
	fmt.Println("     This will be printed second")
}

// deferExample2 demonstrates defer with arguments
func deferExample2() {
	i := 0
	defer fmt.Printf("     Deferred: %d\n", i)  // i is 0
	i++
	fmt.Printf("     Current: %d\n", i)  // i is 1
}

// deferExample3 demonstrates multiple defer statements
func deferExample3() {
	defer fmt.Println("     First defer")
	defer fmt.Println("     Second defer")
	defer fmt.Println("     Third defer")
	fmt.Println("     Function body")
}

// deferExample4 demonstrates defer with anonymous function
func deferExample4() {
	i := 0
	defer func() {
		fmt.Printf("     Deferred: %d\n", i)  // i is current value
	}()
	i++
	fmt.Printf("     Current: %d\n", i)
}

// deferCleanupExample demonstrates defer for cleanup
func deferCleanupExample() {
	// Simulate file operations
	fmt.Println("     Opening file...")
	defer fmt.Println("     Closing file...")
	
	fmt.Println("     Reading file...")
	defer fmt.Println("     Releasing resources...")
	
	fmt.Println("     Processing data...")
}

// demonstratePanicRecover shows panic and recover patterns
func demonstratePanicRecover() {
	fmt.Println("\n6. Panic and Recover:")
	
	// Basic panic recovery
	fmt.Println("   Basic panic recovery:")
	panicRecoverExample()
	
	// Panic recovery in function
	fmt.Println("   Panic recovery in function:")
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("     Error: %v\n", err)
	} else {
		fmt.Printf("     Result: %d\n", result)
	}
	
	// Panic recovery in goroutine
	fmt.Println("   Panic recovery in goroutine:")
	go panicRecoverGoroutine()
	time.Sleep(100 * time.Millisecond) // Wait for goroutine
}

// panicRecoverExample demonstrates basic panic recovery
func panicRecoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("     Recovered from panic: %v\n", r)
		}
	}()
	
	fmt.Println("     About to panic...")
	panic("Something went wrong!")
	fmt.Println("     This won't execute")
}

// safeDivide demonstrates panic recovery in function
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()
	
	if b == 0 {
		panic("division by zero")
	}
	
	result = a / b
	return result, nil
}

// panicRecoverGoroutine demonstrates panic recovery in goroutine
func panicRecoverGoroutine() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("     Goroutine panicked: %v\n", r)
		}
	}()
	
	fmt.Println("     Goroutine running...")
	panic("Goroutine panic!")
}

// Helper functions
func checkFile(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return err
	}
	return nil
}

func getName() string {
	return "Alice"
}
