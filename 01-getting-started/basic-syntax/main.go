package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This example demonstrates Go's basic syntax including variables, constants, and types
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Basic Syntax Examples ===")
	
	// Demonstrate variable declarations
	demonstrateVariables()
	
	// Demonstrate constants
	demonstrateConstants()
	
	// Demonstrate basic types
	demonstrateTypes()
	
	// Demonstrate type conversion
	demonstrateTypeConversion()
	
	// Demonstrate operators
	demonstrateOperators()
	
	// Demonstrate string operations
	demonstrateStringOperations()
}

// demonstrateVariables shows different ways to declare variables
func demonstrateVariables() {
	fmt.Println("\n1. Variable Declarations:")
	
	// Method 1: var keyword with type
	var name string
	var age int
	var isActive bool
	fmt.Printf("   Zero values - name: '%s', age: %d, isActive: %t\n", name, age, isActive)
	
	// Method 2: var with initialization
	var userName string = "Alice"
	var userAge int = 30
	var userActive bool = true
	fmt.Printf("   Initialized - userName: '%s', userAge: %d, userActive: %t\n", userName, userAge, userActive)
	
	// Method 3: Type inference (short declaration)
	email := "alice@example.com"
	score := 95.5
	enabled := true
	fmt.Printf("   Inferred - email: '%s', score: %.1f, enabled: %t\n", email, score, enabled)
	
	// Method 4: Multiple variables
	var (
		firstName string = "Bob"
		lastName  string = "Smith"
		height    int    = 180
		weight    float64 = 75.5
	)
	fmt.Printf("   Multiple - %s %s, height: %dcm, weight: %.1fkg\n", firstName, lastName, height, weight)
	
	// Method 5: Multiple assignment
	var x, y = 10, 20
	a, b := "hello", "world"
	fmt.Printf("   Multiple assignment - x: %d, y: %d, a: '%s', b: '%s'\n", x, y, a, b)
	
	// Swap values
	x, y = y, x
	fmt.Printf("   After swap - x: %d, y: %d\n", x, y)
}

// demonstrateConstants shows how to declare and use constants
func demonstrateConstants() {
	fmt.Println("\n2. Constants:")
	
	// Single constant
	const pi = 3.14159
	fmt.Printf("   Pi: %.5f\n", pi)
	
	// Multiple constants
	const (
		StatusOK    = 200
		StatusError = 500
		MaxRetries  = 3
	)
	fmt.Printf("   Status codes - OK: %d, Error: %d, Max retries: %d\n", StatusOK, StatusError, MaxRetries)
	
	// Typed constants
	const piFloat64 float64 = 3.141592653589793
	fmt.Printf("   Typed pi: %.15f\n", piFloat64)
	
	// Computed constants
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)
	fmt.Printf("   Sizes - KB: %d, MB: %d, GB: %d\n", KB, MB, GB)
	
	// String constants
	const (
		AppName    = "MyApp"
		AppVersion = "1.0.0"
	)
	fmt.Printf("   App info - Name: %s, Version: %s\n", AppName, AppVersion)
}

// demonstrateTypes shows Go's basic types
func demonstrateTypes() {
	fmt.Println("\n3. Basic Types:")
	
	// Integer types
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = 9223372036854775807
	fmt.Printf("   Integers - int8: %d, int16: %d, int32: %d, int64: %d, int: %d\n", i8, i16, i32, i64, i)
	
	// Unsigned integers
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 18446744073709551615
	fmt.Printf("   Unsigned - uint8: %d, uint16: %d, uint32: %d, uint64: %d, uint: %d\n", u8, u16, u32, u64, u)
	
	// Floating point
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793
	fmt.Printf("   Floats - float32: %.5f, float64: %.15f\n", f32, f64)
	
	// Complex numbers
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i
	fmt.Printf("   Complex - complex64: %v, complex128: %v\n", c64, c128)
	
	// Byte and rune
	var b byte = 255  // alias for uint8
	var r rune = 'A'  // alias for int32
	fmt.Printf("   Byte and rune - byte: %d, rune: %c (%d)\n", b, r, r)
	
	// Boolean
	var isTrue bool = true
	var isFalse bool = false
	fmt.Printf("   Booleans - true: %t, false: %t\n", isTrue, isFalse)
	
	// String
	var s string = "Hello, 世界!"
	fmt.Printf("   String: '%s' (length: %d)\n", s, len(s))
}

// demonstrateTypeConversion shows how to convert between types
func demonstrateTypeConversion() {
	fmt.Println("\n4. Type Conversion:")
	
	// Basic conversions
	var i int = 42
	var f float64 = float64(i)
	var s string = string(i)  // This converts to Unicode character!
	fmt.Printf("   int %d -> float64 %.1f -> string '%s'\n", i, f, s)
	
	// Proper string conversion
	var num int = 123
	var str string = strconv.Itoa(num)
	fmt.Printf("   int %d -> string '%s'\n", num, str)
	
	// String to number
	var strNum string = "456"
	if n, err := strconv.Atoi(strNum); err == nil {
		fmt.Printf("   string '%s' -> int %d\n", strNum, n)
	}
	
	// Float conversions
	var floatStr string = "3.14"
	if f, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("   string '%s' -> float64 %.2f\n", floatStr, f)
	}
	
	// Formatting numbers to strings
	var formatted string = strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("   float64 3.14159 -> string '%s' (2 decimal places)\n", formatted)
}

// demonstrateOperators shows Go's operators
func demonstrateOperators() {
	fmt.Println("\n5. Operators:")
	
	// Arithmetic operators
	a, b := 15, 4
	fmt.Printf("   Arithmetic - %d + %d = %d\n", a, b, a+b)
	fmt.Printf("   Arithmetic - %d - %d = %d\n", a, b, a-b)
	fmt.Printf("   Arithmetic - %d * %d = %d\n", a, b, a*b)
	fmt.Printf("   Arithmetic - %d / %d = %d\n", a, b, a/b)
	fmt.Printf("   Arithmetic - %d %% %d = %d\n", a, b, a%b)
	
	// Comparison operators
	fmt.Printf("   Comparison - %d == %d: %t\n", a, b, a == b)
	fmt.Printf("   Comparison - %d != %d: %t\n", a, b, a != b)
	fmt.Printf("   Comparison - %d < %d: %t\n", a, b, a < b)
	fmt.Printf("   Comparison - %d > %d: %t\n", a, b, a > b)
	
	// Logical operators
	p, q := true, false
	fmt.Printf("   Logical - %t && %t: %t\n", p, q, p && q)
	fmt.Printf("   Logical - %t || %t: %t\n", p, q, p || q)
	fmt.Printf("   Logical - !%t: %t\n", p, !p)
	
	// Bitwise operators
	x, y := 5, 3  // Binary: 101, 011
	fmt.Printf("   Bitwise - %d & %d = %d (binary: %b)\n", x, y, x&y, x&y)
	fmt.Printf("   Bitwise - %d | %d = %d (binary: %b)\n", x, y, x|y, x|y)
	fmt.Printf("   Bitwise - %d ^ %d = %d (binary: %b)\n", x, y, x^y, x^y)
	fmt.Printf("   Bitwise - %d << 1 = %d (binary: %b)\n", x, x<<1, x<<1)
	fmt.Printf("   Bitwise - %d >> 1 = %d (binary: %b)\n", x, x>>1, x>>1)
	
	// Assignment operators
	z := 10
	z += 5
	fmt.Printf("   Assignment - z += 5: %d\n", z)
	z *= 2
	fmt.Printf("   Assignment - z *= 2: %d\n", z)
}

// demonstrateStringOperations shows string manipulation
func demonstrateStringOperations() {
	fmt.Println("\n6. String Operations:")
	
	text := "Hello, World!"
	fmt.Printf("   Original: '%s'\n", text)
	
	// Case conversion
	upper := strings.ToUpper(text)
	lower := strings.ToLower(text)
	fmt.Printf("   Upper: '%s'\n", upper)
	fmt.Printf("   Lower: '%s'\n", lower)
	
	// Searching
	contains := strings.Contains(text, "World")
	hasPrefix := strings.HasPrefix(text, "Hello")
	hasSuffix := strings.HasSuffix(text, "!")
	index := strings.Index(text, "World")
	fmt.Printf("   Contains 'World': %t\n", contains)
	fmt.Printf("   Has prefix 'Hello': %t\n", hasPrefix)
	fmt.Printf("   Has suffix '!': %t\n", hasSuffix)
	fmt.Printf("   Index of 'World': %d\n", index)
	
	// Manipulation
	replaced := strings.Replace(text, "World", "Go", 1)
	trimmed := strings.TrimSpace("  hello  ")
	splitted := strings.Split(text, ",")
	joined := strings.Join([]string{"Go", "is", "awesome"}, " ")
	fmt.Printf("   Replaced: '%s'\n", replaced)
	fmt.Printf("   Trimmed: '%s'\n", trimmed)
	fmt.Printf("   Split: %v\n", splitted)
	fmt.Printf("   Joined: '%s'\n", joined)
	
	// String formatting
	name := "Alice"
	age := 30
	formatted := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Printf("   Formatted: '%s'\n", formatted)
	
	// String length and indexing
	fmt.Printf("   Length: %d\n", len(text))
	fmt.Printf("   First character: '%c'\n", text[0])
	fmt.Printf("   Last character: '%c'\n", text[len(text)-1])
}
