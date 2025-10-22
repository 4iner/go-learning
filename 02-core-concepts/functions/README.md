# Go Functions

This comprehensive guide covers Go's function system including function declarations, parameters, return values, closures, and advanced function concepts. Understanding functions is essential for writing modular and reusable Go code.

## Table of Contents
1. [Basic Functions](#basic-functions)
2. [Function Parameters](#function-parameters)
3. [Return Values](#return-values)
4. [Variadic Functions](#variadic-functions)
5. [Anonymous Functions](#anonymous-functions)
6. [Closures](#closures)
7. [Higher-Order Functions](#higher-order-functions)
8. [Method Receivers](#method-receivers)
9. [Function Types](#function-types)
10. [Best Practices](#best-practices)

## Basic Functions

### Function Declaration

```go
// Basic function declaration
func functionName() {
    // function body
}

// Function with parameters
func add(a int, b int) int {
    return a + b
}

// Function with multiple parameters of same type
func multiply(x, y int) int {
    return x * y
}

// Function with multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Function Calls

```go
// Calling functions
result := add(5, 3)        // 8
product := multiply(4, 6)  // 24

// Multiple return values
quotient, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(quotient)  // 5

// Ignoring return values
divide(10, 2)  // Ignore both return values
_, err := divide(10, 0)  // Ignore quotient, check error
```

## Function Parameters

### Value Parameters

```go
// Parameters are passed by value
func modifyValue(x int) {
    x = 100  // Doesn't affect original
}

func main() {
    a := 5
    modifyValue(a)
    fmt.Println(a)  // Still 5
}
```

### Pointer Parameters

```go
// Parameters passed by reference
func modifyPointer(x *int) {
    *x = 100  // Modifies original
}

func main() {
    a := 5
    modifyPointer(&a)
    fmt.Println(a)  // Now 100
}
```

### Slice and Map Parameters

```go
// Slices and maps are reference types
func modifySlice(s []int) {
    s[0] = 100  // Modifies original slice
}

func modifyMap(m map[string]int) {
    m["key"] = 100  // Modifies original map
}

func main() {
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println(slice)  // [100 2 3]
    
    m := map[string]int{"key": 1}
    modifyMap(m)
    fmt.Println(m)  // map[key:100]
}
```

### Struct Parameters

```go
type Person struct {
    Name string
    Age  int
}

// Value receiver (copies struct)
func (p Person) GetName() string {
    return p.Name
}

// Pointer receiver (modifies original)
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Function with struct parameter
func updatePerson(p *Person, name string, age int) {
    p.Name = name
    p.Age = age
}
```

## Return Values

### Single Return Value

```go
func square(x int) int {
    return x * x
}

func isEven(n int) bool {
    return n%2 == 0
}
```

### Multiple Return Values

```go
// Multiple return values
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Named return values
func divideNamed(a, b int) (result int, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return  // Returns zero value for result
    }
    result = a / b
    return  // Returns named values
}

// Multiple return values of same type
func getMinMax(numbers []int) (min, max int) {
    if len(numbers) == 0 {
        return 0, 0
    }
    
    min, max = numbers[0], numbers[0]
    for _, num := range numbers {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    return
}
```

### Error Handling Pattern

```go
// Common Go error handling pattern
func readFile(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    content, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }
    
    return content, nil
}

// Usage
content, err := readFile("example.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(content))
```

## Variadic Functions

### Basic Variadic Functions

```go
// Variadic function (variable number of arguments)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result1 := sum(1, 2, 3, 4, 5)  // 15
result2 := sum(10, 20)         // 30
result3 := sum()               // 0

// Passing slice to variadic function
numbers := []int{1, 2, 3, 4, 5}
result4 := sum(numbers...)     // 15
```

### Variadic Functions with Different Types

```go
// Variadic function with interface{}
func printValues(values ...interface{}) {
    for i, value := range values {
        fmt.Printf("Value %d: %v (type: %T)\n", i, value, value)
    }
}

// Usage
printValues(1, "hello", true, 3.14, []int{1, 2, 3})
```

### Variadic Functions with Mixed Parameters

```go
// Variadic function with regular parameters
func formatString(prefix string, values ...interface{}) string {
    var parts []string
    for _, value := range values {
        parts = append(parts, fmt.Sprintf("%v", value))
    }
    return prefix + strings.Join(parts, " ")
}

// Usage
result := formatString("Result: ", 1, 2, 3)  // "Result: 1 2 3"
```

## Anonymous Functions

### Basic Anonymous Functions

```go
// Anonymous function assigned to variable
add := func(a, b int) int {
    return a + b
}

result := add(5, 3)  // 8

// Anonymous function called immediately
result2 := func(a, b int) int {
    return a * b
}(4, 6)  // 24
```

### Anonymous Functions as Parameters

```go
// Function that takes another function as parameter
func processNumbers(numbers []int, processor func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = processor(num)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5}

// Square each number
squared := processNumbers(numbers, func(x int) int {
    return x * x
})

// Double each number
doubled := processNumbers(numbers, func(x int) int {
    return x * 2
})
```

## Closures

### Basic Closures

```go
// Closure: function that captures variables from outer scope
func createCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Usage
counter := createCounter()
fmt.Println(counter())  // 1
fmt.Println(counter())  // 2
fmt.Println(counter())  // 3

// Each closure has its own captured variables
counter2 := createCounter()
fmt.Println(counter2())  // 1
fmt.Println(counter())    // 4
```

### Closures with Parameters

```go
// Closure with parameters
func createMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Usage
double := createMultiplier(2)
triple := createMultiplier(3)

fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

### Closures in Loops

```go
// Common mistake with closures in loops
func closureMistake() {
    var funcs []func() int
    
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() int {
            return i  // All functions return 3!
        })
    }
    
    for _, f := range funcs {
        fmt.Println(f())  // Prints 3, 3, 3
    }
}

// Correct way to handle closures in loops
func closureCorrect() {
    var funcs []func() int
    
    for i := 0; i < 3; i++ {
        i := i  // Create new variable for each iteration
        funcs = append(funcs, func() int {
            return i
        })
    }
    
    for _, f := range funcs {
        fmt.Println(f())  // Prints 0, 1, 2
    }
}
```

## Higher-Order Functions

### Functions as Parameters

```go
// Higher-order function: takes function as parameter
func filter(numbers []int, predicate func(int) bool) []int {
    var result []int
    for _, num := range numbers {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Filter even numbers
evens := filter(numbers, func(x int) bool {
    return x%2 == 0
})

// Filter numbers greater than 5
greaterThan5 := filter(numbers, func(x int) bool {
    return x > 5
})
```

### Functions as Return Values

```go
// Higher-order function: returns function
func createValidator(min, max int) func(int) bool {
    return func(value int) bool {
        return value >= min && value <= max
    }
}

// Usage
validateAge := createValidator(0, 120)
validateScore := createValidator(0, 100)

fmt.Println(validateAge(25))    // true
fmt.Println(validateAge(150))    // false
fmt.Println(validateScore(85))   // true
fmt.Println(validateScore(150))  // false
```

### Function Composition

```go
// Function composition
func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}

// Usage
square := func(x int) int { return x * x }
double := func(x int) int { return x * 2 }

squareThenDouble := compose(double, square)
doubleThenSquare := compose(square, double)

fmt.Println(squareThenDouble(3))  // 18 (3^2 * 2)
fmt.Println(doubleThenSquare(3))  // 36 ((3*2)^2)
```

## Method Receivers

### Value Receivers

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Value receiver (copies struct)
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Value receiver (doesn't modify original)
func (r Rectangle) Scale(factor float64) Rectangle {
    return Rectangle{
        Width:  r.Width * factor,
        Height: r.Height * factor,
    }
}
```

### Pointer Receivers

```go
// Pointer receiver (modifies original)
func (r *Rectangle) ScaleInPlace(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Pointer receiver for methods that might modify
func (r *Rectangle) SetDimensions(width, height float64) {
    r.Width = width
    r.Height = height
}
```

### Method Sets

```go
// Method sets determine which methods are available
type T struct{}

func (T) method1() {}    // Value receiver
func (*T) method2() {}   // Pointer receiver

func main() {
    var t T
    var p *T = &t
    
    t.method1()  // OK
    t.method2()  // OK (Go automatically takes address)
    
    p.method1()  // OK (Go automatically dereferences)
    p.method2()  // OK
}
```

## Function Types

### Function Type Declarations

```go
// Function type declaration
type BinaryOp func(int, int) int
type Predicate func(int) bool
type Transformer func(int) int

// Using function types
func applyOperation(a, b int, op BinaryOp) int {
    return op(a, b)
}

// Usage
add := func(a, b int) int { return a + b }
multiply := func(a, b int) int { return a * b }

result1 := applyOperation(5, 3, add)      // 8
result2 := applyOperation(5, 3, multiply) // 15
```

### Function Type Interfaces

```go
// Interface with function type
type Processor interface {
    Process(int) int
}

// Function type implementing interface
type FuncProcessor func(int) int

func (f FuncProcessor) Process(x int) int {
    return f(x)
}

// Usage
var processor Processor = FuncProcessor(func(x int) int {
    return x * 2
})

result := processor.Process(5)  // 10
```

## Best Practices

### 1. Use Meaningful Function Names

```go
// Good: Descriptive function names
func calculateTotalPrice(items []Item) float64 {
    // implementation
}

func validateUserInput(input string) error {
    // implementation
}

// Avoid: Unclear function names
func calc(items []Item) float64 {
    // implementation
}

func check(input string) error {
    // implementation
}
```

### 2. Keep Functions Small and Focused

```go
// Good: Single responsibility
func calculateTax(amount float64, rate float64) float64 {
    return amount * rate
}

func formatCurrency(amount float64) string {
    return fmt.Sprintf("$%.2f", amount)
}

// Avoid: Multiple responsibilities
func processOrder(order Order) (float64, string, error) {
    // Calculate tax
    // Format currency
    // Validate order
    // Send email
    // Update database
    // ...
}
```

### 3. Use Error Returns Consistently

```go
// Good: Consistent error handling
func readConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("failed to open config file: %w", err)
    }
    defer file.Close()
    
    // Read and parse config
    // ...
    return config, nil
}

// Avoid: Inconsistent error handling
func readConfigBad(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err  // No context
    }
    defer file.Close()
    
    // Read and parse config
    // ...
    return config, nil
}
```

### 4. Use Pointer Receivers Appropriately

```go
// Good: Use pointer receivers for methods that modify
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Good: Use value receivers for methods that don't modify
func (p Person) GetAge() int {
    return p.Age
}

// Good: Use pointer receivers for large structs
func (p *Person) UpdateProfile(profile Profile) {
    p.Profile = profile
}
```

### 5. Use Variadic Functions Sparingly

```go
// Good: Use variadic functions when appropriate
func logf(format string, args ...interface{}) {
    fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), 
        fmt.Sprintf(format, args...))
}

// Avoid: Overusing variadic functions
func processData(data ...interface{}) {
    // Too generic, hard to understand
}
```

## References

- [Go Language Specification - Functions](https://golang.org/ref/spec#Function_declarations)
- [Effective Go - Functions](https://golang.org/doc/effective_go.html#functions)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Closures](https://gobyexample.com/closures)

## Next Steps

After mastering functions, continue with:
- [Interfaces](../interfaces/) - Learn Go's interface system
- [Concurrency](../concurrency/) - Master goroutines and channels
