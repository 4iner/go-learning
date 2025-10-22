# Go Control Flow

This comprehensive guide covers Go's control flow structures including conditional statements, loops, and the unique `defer` statement. Understanding control flow is essential for writing logical and efficient Go programs.

## Table of Contents
1. [If Statements](#if-statements)
2. [Switch Statements](#switch-statements)
3. [For Loops](#for-loops)
4. [Range Loops](#range-loops)
5. [Defer Statement](#defer-statement)
6. [Panic and Recover](#panic-and-recover)
7. [Best Practices](#best-practices)

## If Statements

### Basic If Statement

```go
// Basic if statement
if condition {
    // code to execute if condition is true
}

// If-else statement
if condition {
    // code if true
} else {
    // code if false
}

// If-else if-else chain
if condition1 {
    // code if condition1 is true
} else if condition2 {
    // code if condition2 is true
} else {
    // code if all conditions are false
}
```

### If with Initialization

Go allows initialization statements in if conditions:

```go
// Initialize variable in if condition
if err := doSomething(); err != nil {
    // handle error
}

// Multiple initialization
if x, y := getValues(); x > y {
    // use x and y
}

// Common pattern with maps
if value, exists := myMap["key"]; exists {
    // use value
}
```

### If with Short Variable Declaration

```go
// Short variable declaration in if
if name := getName(); name != "" {
    fmt.Printf("Hello, %s!\n", name)
}

// Variable is scoped to the if block
if age := getAge(); age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
// age is not accessible here
```

## Switch Statements

### Basic Switch

```go
// Basic switch statement
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
case 3:
    fmt.Println("Three")
default:
    fmt.Println("Other")
}
```

### Switch with Multiple Values

```go
// Multiple values in case
switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid day")
}
```

### Switch with Initialization

```go
// Switch with initialization
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("macOS")
case "linux":
    fmt.Println("Linux")
case "windows":
    fmt.Println("Windows")
default:
    fmt.Printf("Other: %s\n", os)
}
```

### Type Switch

```go
// Type switch for interface{} values
func processValue(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Switch without Expression

```go
// Switch without expression (like if-else chain)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
case score >= 60:
    grade = "D"
default:
    grade = "F"
}
```

### Fallthrough

```go
// Use fallthrough to continue to next case
switch value {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two or One")
case 3:
    fmt.Println("Three")
}
```

## For Loops

### Basic For Loop

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Printf("i = %d\n", i)
}

// Multiple variables
for i, j := 0, 10; i < j; i, j = i+1, j-1 {
    fmt.Printf("i = %d, j = %d\n", i, j)
}
```

### While-style Loop

```go
// While-style loop (no initialization or post statement)
i := 0
for i < 10 {
    fmt.Printf("i = %d\n", i)
    i++
}

// Infinite loop
for {
    // infinite loop
    if condition {
        break
    }
}
```

### For with Range

```go
// Range over slice
slice := []string{"a", "b", "c"}
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}

// Range over map
m := map[string]int{"a": 1, "b": 2, "c": 3}
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}

// Range over string
s := "Hello"
for index, char := range s {
    fmt.Printf("Index: %d, Char: %c\n", index, char)
}

// Range over channel
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)
for value := range ch {
    fmt.Printf("Value: %d\n", value)
}
```

### Loop Control

```go
// Break statement
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // exit loop
    }
    fmt.Printf("i = %d\n", i)
}

// Continue statement
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue  // skip to next iteration
    }
    fmt.Printf("Odd: %d\n", i)
}

// Labeled break and continue
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer  // break outer loop
        }
        fmt.Printf("i = %d, j = %d\n", i, j)
    }
}
```

## Range Loops

### Range over Different Types

```go
// Range over slice (index and value)
slice := []int{1, 2, 3, 4, 5}
for i, v := range slice {
    fmt.Printf("Index: %d, Value: %d\n", i, v)
}

// Range over slice (value only)
for _, v := range slice {
    fmt.Printf("Value: %d\n", v)
}

// Range over slice (index only)
for i := range slice {
    fmt.Printf("Index: %d\n", i)
}

// Range over map
m := map[string]int{"a": 1, "b": 2, "c": 3}
for k, v := range m {
    fmt.Printf("Key: %s, Value: %d\n", k, v)
}

// Range over string (rune by rune)
s := "Hello, 世界"
for i, r := range s {
    fmt.Printf("Index: %d, Rune: %c\n", i, r)
}
```

## Defer Statement

### Basic Defer

```go
// Defer executes function when surrounding function returns
func example() {
    defer fmt.Println("This will be printed last")
    fmt.Println("This will be printed first")
    fmt.Println("This will be printed second")
}
// Output:
// This will be printed first
// This will be printed second
// This will be printed last
```

### Defer with Arguments

```go
// Arguments are evaluated when defer is called
func example() {
    i := 0
    defer fmt.Println("Deferred:", i)  // i is 0
    i++
    fmt.Println("Current:", i)  // i is 1
}
// Output:
// Current: 1
// Deferred: 0
```

### Multiple Defer Statements

```go
// Multiple defer statements execute in LIFO order
func example() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")
    fmt.Println("Function body")
}
// Output:
// Function body
// Third defer
// Second defer
// First defer
```

### Defer with Anonymous Functions

```go
// Use anonymous function to defer with current values
func example() {
    i := 0
    defer func() {
        fmt.Println("Deferred:", i)  // i is current value
    }()
    i++
    fmt.Println("Current:", i)
}
// Output:
// Current: 1
// Deferred: 1
```

### Common Defer Patterns

```go
// File operations
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // Always close file
    
    // Read file content
    // ...
    return nil
}

// Mutex locking
func safeOperation() {
    mu.Lock()
    defer mu.Unlock()  // Always unlock
    
    // Critical section
    // ...
}

// Database transactions
func updateUser(id int, name string) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()  // Rollback by default
    
    // Update user
    if err := tx.Exec("UPDATE users SET name = ? WHERE id = ?", name, id); err != nil {
        return err
    }
    
    return tx.Commit()  // Commit overrides rollback
}
```

## Panic and Recover

### Panic

```go
// Panic stops normal execution and starts panicking
func example() {
    fmt.Println("Before panic")
    panic("Something went wrong!")
    fmt.Println("After panic")  // This won't execute
}

// Panic with error
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

### Recover

```go
// Recover stops panic and returns the value passed to panic
func example() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    panic("Something went wrong!")
    fmt.Println("This won't execute")
}
// Output: Recovered from panic: Something went wrong!
```

### Panic and Recover Patterns

```go
// Safe division with recover
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic occurred: %v", r)
        }
    }()
    
    result = a / b
    return result, nil
}

// Panic recovery in goroutines
func worker() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Worker panicked: %v\n", r)
        }
    }()
    
    // Worker logic that might panic
    // ...
}
```

## Best Practices

### 1. Use Early Returns

```go
// Good: Early return reduces nesting
func processUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    }
    
    if user.Name == "" {
        return errors.New("user name is empty")
    }
    
    // Main logic
    return nil
}

// Avoid: Deep nesting
func processUser(user *User) error {
    if user != nil {
        if user.Name != "" {
            // Main logic
            return nil
        } else {
            return errors.New("user name is empty")
        }
    } else {
        return errors.New("user is nil")
    }
}
```

### 2. Use Switch for Multiple Conditions

```go
// Good: Use switch for multiple conditions
func getGrade(score int) string {
    switch {
    case score >= 90:
        return "A"
    case score >= 80:
        return "B"
    case score >= 70:
        return "C"
    case score >= 60:
        return "D"
    default:
        return "F"
    }
}

// Avoid: Long if-else chain
func getGrade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    } else if score >= 70 {
        return "C"
    } else if score >= 60 {
        return "D"
    } else {
        return "F"
    }
}
```

### 3. Use Defer for Cleanup

```go
// Good: Always use defer for cleanup
func readConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    // Parse config
    // ...
    return config, nil
}
```

### 4. Avoid Panic in Normal Code

```go
// Good: Return errors instead of panicking
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Avoid: Panic for normal error conditions
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}
```

### 5. Use Range Appropriately

```go
// Good: Use range for iteration
for i, v := range items {
    fmt.Printf("Item %d: %v\n", i, v)
}

// Good: Use range for maps
for k, v := range m {
    fmt.Printf("%s: %v\n", k, v)
}

// Avoid: Manual indexing when range is available
for i := 0; i < len(items); i++ {
    fmt.Printf("Item %d: %v\n", i, items[i])
}
```

## References

- [Go Language Specification - Control Flow](https://golang.org/ref/spec#Control_structures)
- [Effective Go - Control Structures](https://golang.org/doc/effective_go.html#control)
- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Go by Example - Switch](https://gobyexample.com/switch)
- [Go by Example - For](https://gobyexample.com/for)
- [Go by Example - Defer](https://gobyexample.com/defer)

## Next Steps

After mastering control flow, continue with:
- [Functions](../functions/) - Learn function declarations and usage
- [Data Structures](../data-structures/) - Master Go's data structures
