# Go Best Practices

This comprehensive guide covers Go's best practices and idiomatic patterns. Following these practices will help you write clean, maintainable, and efficient Go code that follows the language's philosophy and conventions.

## Table of Contents
1. [Code Style](#code-style)
2. [Naming Conventions](#naming-conventions)
3. [Error Handling](#error-handling)
4. [Concurrency](#concurrency)
5. [Performance](#performance)
6. [Testing](#testing)
7. [Package Design](#package-design)
8. [Documentation](#documentation)

## Code Style

### Formatting

```go
// Good: Use gofmt
func processData(data []byte) ([]byte, error) {
    if len(data) == 0 {
        return nil, errors.New("empty data")
    }
    
    // Process data
    return data, nil
}

// Good: Consistent indentation
func complexFunction() {
    if condition {
        for i := 0; i < 10; i++ {
            if i%2 == 0 {
                fmt.Println(i)
            }
        }
    }
}
```

### Variable Declarations

```go
// Good: Use short variable declarations
func processUser(userID int) error {
    user, err := getUser(userID)
    if err != nil {
        return err
    }
    
    return processUserData(user)
}

// Good: Group related declarations
var (
    maxRetries = 3
    timeout    = 30 * time.Second
    debug      = false
)

// Avoid: Unnecessary var keyword
func badExample() {
    var name string = "John"
    var age int = 25
}
```

### Function Design

```go
// Good: Small, focused functions
func validateUser(user User) error {
    if user.Name == "" {
        return errors.New("name is required")
    }
    
    if user.Age < 0 {
        return errors.New("age must be positive")
    }
    
    return nil
}

// Good: Single responsibility
func calculateTotal(items []Item) float64 {
    total := 0.0
    for _, item := range items {
        total += item.Price
    }
    return total
}

// Avoid: Functions that do too much
func badFunction(user User, items []Item) (float64, error) {
    // Validate user
    // Process items
    // Calculate total
    // Send email
    // Update database
    // ...
}
```

## Naming Conventions

### Package Names

```go
// Good: Short, descriptive package names
package math
package user
package http

// Good: Use lowercase, no underscores
package stringutils
package filemanager

// Avoid: Long, verbose names
package mathematicaloperations
package userdatamanagement
```

### Function Names

```go
// Good: Descriptive function names
func calculateTotal(items []Item) float64
func validateUser(user User) error
func processPayment(amount float64) error

// Good: Use verbs for functions
func getUser(id int) (*User, error)
func saveUser(user *User) error
func deleteUser(id int) error

// Avoid: Unclear function names
func calc(items []Item) float64
func check(user User) error
func do(amount float64) error
```

### Variable Names

```go
// Good: Descriptive variable names
func processOrder(order Order) error {
    customerName := order.Customer.Name
    totalAmount := calculateTotal(order.Items)
    
    if totalAmount > 1000 {
        return errors.New("amount too large")
    }
    
    return nil
}

// Good: Use meaningful names
func calculateTax(amount float64, rate float64) float64 {
    taxAmount := amount * rate
    return taxAmount
}

// Avoid: Abbreviations and unclear names
func badExample(order Order) error {
    cn := order.Customer.Name
    ta := calcTotal(order.Items)
    
    if ta > 1000 {
        return errors.New("too large")
    }
    
    return nil
}
```

### Constants

```go
// Good: UPPERCASE for exported constants
const (
    MaxRetries = 3
    DefaultTimeout = 30 * time.Second
    StatusOK = 200
)

// Good: camelCase for unexported constants
const (
    maxBufferSize = 1024
    defaultPort = 8080
)

// Good: Group related constants
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)
```

## Error Handling

### Explicit Error Handling

```go
// Good: Handle errors explicitly
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    content, err := io.ReadAll(file)
    if err != nil {
        return fmt.Errorf("failed to read file: %w", err)
    }
    
    return processContent(content)
}

// Good: Use errors.Is and errors.As
func handleError(err error) {
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File not found")
        return
    }
    
    var validationErr ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation error: %s\n", validationErr.Message)
        return
    }
    
    fmt.Printf("Unknown error: %v\n", err)
}
```

### Error Wrapping

```go
// Good: Wrap errors with context
func processUser(userID int) error {
    user, err := getUser(userID)
    if err != nil {
        return fmt.Errorf("failed to get user %d: %w", userID, err)
    }
    
    if err := validateUser(user); err != nil {
        return fmt.Errorf("failed to validate user %d: %w", userID, err)
    }
    
    return nil
}

// Avoid: Losing error context
func badProcessUser(userID int) error {
    user, err := getUser(userID)
    if err != nil {
        return err  // Lost context
    }
    
    if err := validateUser(user); err != nil {
        return err  // Lost context
    }
    
    return nil
}
```

## Concurrency

### Use Channels for Communication

```go
// Good: Use channels for communication
func processData(input <-chan int, output chan<- int) {
    for data := range input {
        result := data * 2
        output <- result
    }
}

// Good: Use WaitGroup for synchronization
func processWorkers() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Do work
        }(i)
    }
    
    wg.Wait()
}

// Avoid: Shared memory with mutexes
type BadCounter struct {
    mu    sync.Mutex
    value int
}

func (c *BadCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}
```

### Context for Cancellation

```go
// Good: Use context for cancellation
func longRunningOperation(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
            time.Sleep(100 * time.Millisecond)
        }
    }
}

// Good: Pass context through call chain
func processRequest(ctx context.Context, req Request) error {
    if err := validateRequest(ctx, req); err != nil {
        return err
    }
    
    if err := processData(ctx, req.Data); err != nil {
        return err
    }
    
    return nil
}
```

## Performance

### Memory Allocation

```go
// Good: Pre-allocate slices when size is known
func processItems(items []Item) []Result {
    results := make([]Result, 0, len(items))
    
    for _, item := range items {
        result := processItem(item)
        results = append(results, result)
    }
    
    return results
}

// Good: Use sync.Pool for frequently allocated objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func processData(data []byte) {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    
    // Use buffer
}
```

### String Operations

```go
// Good: Use strings.Builder for string concatenation
func buildString(parts []string) string {
    var builder strings.Builder
    builder.Grow(100)  // Pre-allocate capacity
    
    for _, part := range parts {
        builder.WriteString(part)
    }
    
    return builder.String()
}

// Avoid: String concatenation in loops
func badBuildString(parts []string) string {
    result := ""
    for _, part := range parts {
        result += part  // Inefficient
    }
    return result
}
```

### Avoid Unnecessary Allocations

```go
// Good: Reuse slices
func processData(data []byte) []byte {
    result := data[:0]  // Reuse underlying array
    
    for _, b := range data {
        if b != 0 {
            result = append(result, b)
        }
    }
    
    return result
}

// Good: Use range for iteration
func sumNumbers(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}
```

## Testing

### Test Structure

```go
// Good: Use table-driven tests
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed signs", -2, 3, 1},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Good: Use test helpers
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

### Test Coverage

```go
// Good: Aim for high test coverage
func TestUserService(t *testing.T) {
    service := NewUserService()
    
    // Test successful operations
    user, err := service.GetUser(1)
    assertNoError(t, err)
    assertEqual(t, user.ID, 1)
    
    // Test error cases
    _, err = service.GetUser(999)
    assertError(t, err)
    
    // Test edge cases
    user, err = service.GetUser(0)
    assertError(t, err)
}
```

## Package Design

### Single Responsibility

```go
// Good: Single responsibility packages
// math/arithmetic.go
package math

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

// user/user.go
package user

type User struct {
    ID   int
    Name string
}

func (u *User) GetName() string {
    return u.Name
}
```

### Package Interfaces

```go
// Good: Define interfaces in consumer packages
package user

type Service interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

// Good: Use interfaces for testing
func TestUserHandler(t *testing.T) {
    mockService := &MockUserService{}
    handler := NewUserHandler(mockService)
    
    // Test handler
}
```

### Package Initialization

```go
// Good: Use init() for package setup
package config

func init() {
    loadConfig()
    setupLogging()
}

// Avoid: Complex logic in init()
package bad

func init() {
    // Don't do complex business logic
    // Don't start servers
    // Don't make network calls
}
```

## Documentation

### Package Documentation

```go
// Package math provides mathematical operations.
//
// This package includes basic arithmetic operations
// and more advanced mathematical functions.
package math

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}
```

### Function Documentation

```go
// ProcessUser processes a user and returns the result.
//
// The function validates the user data, applies business rules,
// and returns a processed user object.
//
// Example:
//   user := User{Name: "Alice", Age: 30}
//   result, err := ProcessUser(user)
//   if err != nil {
//       log.Fatal(err)
//   }
func ProcessUser(user User) (*ProcessedUser, error) {
    // Implementation
}
```

### Type Documentation

```go
// User represents a user in the system.
//
// User contains basic information about a person
// and provides methods for user management.
type User struct {
    // Name is the user's full name.
    Name string
    
    // Age is the user's age in years.
    Age int
    
    // Email is the user's email address.
    Email string
}
```

## Common Anti-Patterns

### 1. Panic for Normal Errors

```go
// Bad: Panic for normal errors
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

// Good: Return errors for normal conditions
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### 2. Ignoring Errors

```go
// Bad: Ignoring errors
func badExample() {
    file, _ := os.Open("file.txt")  // Ignoring error
    defer file.Close()
    
    // Process file
}

// Good: Handle errors
func goodExample() error {
    file, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Process file
    return nil
}
```

### 3. Deep Nesting

```go
// Bad: Deep nesting
func badExample(user User) error {
    if user.Name != "" {
        if user.Age > 0 {
            if user.Email != "" {
                // Process user
                return nil
            } else {
                return errors.New("email required")
            }
        } else {
            return errors.New("age must be positive")
        }
    } else {
        return errors.New("name required")
    }
}

// Good: Early returns
func goodExample(user User) error {
    if user.Name == "" {
        return errors.New("name required")
    }
    
    if user.Age <= 0 {
        return errors.New("age must be positive")
    }
    
    if user.Email == "" {
        return errors.New("email required")
    }
    
    // Process user
    return nil
}
```

## References

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go by Example](https://gobyexample.com/)
- [Go Best Practices](https://github.com/golang/go/wiki/CodeReviewComments)

## Next Steps

After mastering best practices, you're ready to:
- Build real-world applications
- Contribute to open source projects
- Explore advanced Go topics
- Learn Go-specific frameworks and libraries
