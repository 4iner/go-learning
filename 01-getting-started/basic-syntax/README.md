# Go Basic Syntax

This comprehensive guide covers Go's fundamental syntax, including variables, constants, types, operators, and basic language constructs. Understanding these concepts is essential for writing effective Go programs.

## Table of Contents
1. [Variables and Constants](#variables-and-constants)
2. [Basic Types](#basic-types)
3. [Type Conversion](#type-conversion)
4. [Operators](#operators)
5. [String Operations](#string-operations)
6. [Comments](#comments)
7. [Naming Conventions](#naming-conventions)
8. [Best Practices](#best-practices)

## Variables and Constants

### Variable Declaration

Go provides several ways to declare variables:

```go
// Method 1: var keyword with type
var name string
var age int
var isActive bool

// Method 2: var with initialization
var name string = "John"
var age int = 25
var isActive bool = true

// Method 3: Type inference (short declaration)
name := "John"
age := 25
isActive := true

// Method 4: Multiple variables
var (
    name     string = "John"
    age      int    = 25
    isActive bool   = true
)

// Method 5: Multiple assignment
var name, age = "John", 25
name, age := "John", 25
```

### Constants

Constants are declared using the `const` keyword:

```go
// Single constant
const pi = 3.14159

// Multiple constants
const (
    StatusOK    = 200
    StatusError = 500
    MaxRetries  = 3
)

// Typed constants
const pi float64 = 3.14159

// Computed constants
const (
    KB = 1024
    MB = KB * 1024
    GB = MB * 1024
)
```

### Zero Values

In Go, variables are initialized with zero values if not explicitly initialized:

```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var p *int       // nil
var slice []int  // nil
var m map[string]int // nil
```

## Basic Types

### Numeric Types

```go
// Integers
var i8  int8   = 127
var i16 int16  = 32767
var i32 int32  = 2147483647
var i64 int64  = 9223372036854775807
var i   int    = 9223372036854775807  // Platform-dependent

// Unsigned integers
var u8  uint8   = 255
var u16 uint16  = 65535
var u32 uint32  = 4294967295
var u64 uint64  = 18446744073709551615
var u   uint    = 18446744073709551615

// Floating point
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Complex numbers
var c64 complex64  = 1 + 2i
var c128 complex128 = 1 + 2i

// Byte and rune
var b byte = 255  // alias for uint8
var r rune = 'A'  // alias for int32, represents Unicode code point
```

### Boolean Type

```go
var isTrue  bool = true
var isFalse bool = false

// Boolean operations
result := isTrue && isFalse  // false
result = isTrue || isFalse   // true
result = !isTrue             // false
```

### String Type

```go
// String literals
var s1 string = "Hello, World!"
var s2 string = `Multiline
string literal`

// String operations
length := len(s1)                    // 13
char := s1[0]                        // 'H'
substring := s1[0:5]                // "Hello"
concatenated := s1 + " " + s2        // "Hello, World! Multiline..."
```

## Type Conversion

Go requires explicit type conversion:

```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var s string = string(i)    // Convert int to string (Unicode)

// Common conversions
var str string = "123"
var num int = 42

// String to number (requires strconv package)
import "strconv"
n, err := strconv.Atoi(str)  // "123" -> 123
s := strconv.Itoa(num)       // 123 -> "123"

// Number to string with formatting
s = strconv.FormatInt(int64(num), 10)  // Base 10
s = strconv.FormatFloat(3.14, 'f', 2, 64)  // "3.14"
```

## Operators

### Arithmetic Operators

```go
a, b := 10, 3

sum := a + b      // 13
diff := a - b     // 7
product := a * b  // 30
quotient := a / b // 3
remainder := a % b // 1

// Increment and decrement
a++  // a = 11
a--  // a = 10
```

### Comparison Operators

```go
a, b := 10, 20

equal := a == b      // false
notEqual := a != b   // true
less := a < b        // true
lessEqual := a <= b  // true
greater := a > b     // false
greaterEqual := a >= b // false
```

### Logical Operators

```go
p, q := true, false

and := p && q  // false
or := p || q   // true
not := !p      // false
```

### Bitwise Operators

```go
a, b := 5, 3  // Binary: 101, 011

and := a & b   // 1 (001)
or := a | b    // 7 (111)
xor := a ^ b   // 6 (110)
not := ^a      // -6 (two's complement)
leftShift := a << 1  // 10 (1010)
rightShift := a >> 1 // 2 (10)
```

### Assignment Operators

```go
a := 10

a += 5   // a = 15
a -= 3   // a = 12
a *= 2   // a = 24
a /= 4   // a = 6
a %= 5   // a = 1
a &= 3   // a = 1
a |= 2   // a = 3
a ^= 1   // a = 2
a <<= 1  // a = 4
a >>= 1  // a = 2
```

## String Operations

### String Package

```go
import "strings"

s := "Hello, World!"

// Case conversion
upper := strings.ToUpper(s)    // "HELLO, WORLD!"
lower := strings.ToLower(s)    // "hello, world!"

// Searching
contains := strings.Contains(s, "World")  // true
hasPrefix := strings.HasPrefix(s, "Hello") // true
hasSuffix := strings.HasSuffix(s, "!")    // true
index := strings.Index(s, "World")       // 7

// Manipulation
replaced := strings.Replace(s, "World", "Go", 1)  // "Hello, Go!"
trimmed := strings.TrimSpace("  hello  ")         // "hello"
splitted := strings.Split(s, ",")                 // ["Hello", " World!"]
joined := strings.Join([]string{"a", "b"}, "-")   // "a-b"
```

### String Formatting

```go
import "fmt"

name := "John"
age := 25

// Printf formatting
fmt.Printf("Name: %s, Age: %d\n", name, age)
fmt.Printf("Age in hex: %x\n", age)
fmt.Printf("Pi: %.2f\n", 3.14159)

// Sprintf (returns string)
message := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)

// Common format verbs
// %v - default format
// %T - type
// %t - boolean
// %d - integer
// %f - float
// %s - string
// %p - pointer
```

## Comments

Go supports two types of comments:

```go
// Single-line comment

/*
Multi-line comment
Can span multiple lines
*/

// Package comment (should start with package name)
// Package main demonstrates basic Go syntax.
package main

// Function comment (should start with function name)
// main is the entry point of the program.
func main() {
    // Inline comment
    fmt.Println("Hello, World!")
}
```

## Naming Conventions

### General Rules

```go
// Exported (public) identifiers start with uppercase
var PublicVariable = "accessible from other packages"
func PublicFunction() {}

// Unexported (private) identifiers start with lowercase
var privateVariable = "only accessible within package"
func privateFunction() {}

// Constants are typically UPPERCASE
const MAX_SIZE = 100
const DefaultTimeout = 30

// Acronyms should be all uppercase
var HTTPClient *http.Client
var XMLParser xml.Parser
```

### Specific Conventions

```go
// Variables: camelCase
var userName string
var maxRetries int

// Functions: camelCase
func calculateTotal() int {}
func isValidUser() bool {}

// Types: PascalCase
type User struct {}
type DatabaseConnection interface {}

// Interfaces: often end with 'er'
type Reader interface {}
type Writer interface {}
type Stringer interface {}

// Constants: UPPER_CASE or camelCase
const MAX_RETRIES = 3
const defaultTimeout = 30
```

## Best Practices

### 1. Use Short Variable Declarations

```go
// Good: Use := when possible
name := "John"
age := 25

// Avoid: Unnecessary var keyword
var name string = "John"
var age int = 25
```

### 2. Initialize Variables When Declaring

```go
// Good: Initialize with meaningful values
var config Config = loadConfig()
var logger Logger = newLogger()

// Avoid: Declare then assign separately
var config Config
config = loadConfig()
```

### 3. Use Meaningful Names

```go
// Good: Descriptive names
var userCount int
var isAuthenticated bool
var connectionTimeout time.Duration

// Avoid: Abbreviations and unclear names
var cnt int
var auth bool
var timeout time.Duration
```

### 4. Group Related Declarations

```go
// Good: Group related variables
var (
    serverHost string = "localhost"
    serverPort int    = 8080
    serverTimeout time.Duration = 30 * time.Second
)

// Good: Group related constants
const (
    StatusOK    = 200
    StatusError = 500
    StatusNotFound = 404
)
```

### 5. Use Type Inference Appropriately

```go
// Good: Use type inference for simple types
name := "John"        // string
age := 25             // int
isActive := true      // bool

// Good: Specify type for complex types
var users []User
var config map[string]string
```

## References

- [Go Language Specification](https://golang.org/ref/spec)
- [Effective Go - Variables](https://golang.org/doc/effective_go.html#variables)
- [Go by Example - Variables](https://gobyexample.com/variables)
- [Go by Example - Constants](https://gobyexample.com/constants)

## Next Steps

After mastering basic syntax, continue with:
- [Control Flow](../control-flow/) - Learn Go's control structures
- [Functions](../functions/) - Master function declarations and usage
