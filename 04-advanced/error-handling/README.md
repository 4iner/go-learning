# Go Error Handling

This comprehensive guide covers Go's error handling philosophy and patterns. Go treats errors as values, not exceptions, which leads to explicit error handling and more predictable code. Understanding Go's error handling is crucial for writing robust applications.

## Table of Contents
1. [Error Interface](#error-interface)
2. [Basic Error Handling](#basic-error-handling)
3. [Custom Error Types](#custom-error-types)
4. [Error Wrapping](#error-wrapping)
5. [Error Checking Patterns](#error-checking-patterns)
6. [Panic and Recover](#panic-and-recover)
7. [Best Practices](#best-practices)
8. [Common Patterns](#common-patterns)

## Error Interface

### The `error` Interface

```go
// Built-in error interface
type error interface {
    Error() string
}

// Basic error creation
err := errors.New("something went wrong")
fmt.Println(err.Error()) // "something went wrong"

// Errorf for formatted errors
err = fmt.Errorf("user %s not found", "alice")
fmt.Println(err) // "user alice not found"
```

### Error Values

```go
// Errors are values, not exceptions
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 5
```

## Basic Error Handling

### Function Error Returns

```go
// Standard Go pattern: return error as last value
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
    log.Printf("Failed to read file: %v", err)
    return
}
fmt.Println(string(content))
```

### Error Checking

```go
// Explicit error checking
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    
    // Process data
    return nil
}

// Multiple error checks
func complexOperation() error {
    if err := step1(); err != nil {
        return err
    }
    
    if err := step2(); err != nil {
        return err
    }
    
    if err := step3(); err != nil {
        return err
    }
    
    return nil
}
```

### Error Handling with Defer

```go
// Clean up resources even when errors occur
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Always close file
    
    // Process file
    return nil
}
```

## Custom Error Types

### Struct-based Errors

```go
// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// Usage
func validateUser(user User) error {
    if user.Name == "" {
        return ValidationError{
            Field:   "name",
            Message: "name is required",
        }
    }
    
    if user.Age < 0 {
        return ValidationError{
            Field:   "age",
            Message: "age must be positive",
        }
    }
    
    return nil
}
```

### Error with Additional Information

```go
// Error with context
type DatabaseError struct {
    Operation string
    Table     string
    Err       error
}

func (e DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s on table %s: %v", 
        e.Operation, e.Table, e.Err)
}

func (e DatabaseError) Unwrap() error {
    return e.Err
}

// Usage
func saveUser(user User) error {
    if err := db.Save(user); err != nil {
        return DatabaseError{
            Operation: "save",
            Table:     "users",
            Err:       err,
        }
    }
    return nil
}
```

### Error Codes

```go
// Error with codes
type ErrorCode int

const (
    ErrNotFound ErrorCode = iota
    ErrUnauthorized
    ErrValidation
    ErrInternal
)

type AppError struct {
    Code    ErrorCode
    Message string
    Err     error
}

func (e AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %v", e.Message, e.Err)
    }
    return e.Message
}

func (e AppError) Unwrap() error {
    return e.Err
}

// Usage
func getUser(id int) (*User, error) {
    user, err := db.GetUser(id)
    if err != nil {
        return nil, AppError{
            Code:    ErrNotFound,
            Message: "user not found",
            Err:     err,
        }
    }
    return user, nil
}
```

## Error Wrapping

### Basic Error Wrapping

```go
// Wrap errors with context
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
```

### Error Unwrapping

```go
// Unwrap errors to get original error
func handleError(err error) {
    if err == nil {
        return
    }
    
    // Check for specific error types
    var validationErr ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation error on field %s: %s\n", 
            validationErr.Field, validationErr.Message)
        return
    }
    
    // Check for wrapped errors
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File not found")
        return
    }
    
    fmt.Printf("Unknown error: %v\n", err)
}
```

### Error Chain Inspection

```go
// Inspect error chain
func inspectError(err error) {
    for err != nil {
        fmt.Printf("Error: %v\n", err)
        err = errors.Unwrap(err)
    }
}

// Usage
func main() {
    err := processUser(123)
    if err != nil {
        inspectError(err)
    }
}
```

## Error Checking Patterns

### Early Return Pattern

```go
// Good: Early return on error
func processData(data []byte) ([]byte, error) {
    if len(data) == 0 {
        return nil, errors.New("empty data")
    }
    
    if !isValid(data) {
        return nil, errors.New("invalid data")
    }
    
    // Process data
    return data, nil
}

// Avoid: Deep nesting
func processDataBad(data []byte) ([]byte, error) {
    if len(data) > 0 {
        if isValid(data) {
            // Process data
            return data, nil
        } else {
            return nil, errors.New("invalid data")
        }
    } else {
        return nil, errors.New("empty data")
    }
}
```

### Error Aggregation

```go
// Collect multiple errors
type MultiError struct {
    Errors []error
}

func (e MultiError) Error() string {
    var messages []string
    for _, err := range e.Errors {
        messages = append(messages, err.Error())
    }
    return strings.Join(messages, "; ")
}

func validateUser(user User) error {
    var errors MultiError
    
    if user.Name == "" {
        errors.Errors = append(errors.Errors, 
            errors.New("name is required"))
    }
    
    if user.Age < 0 {
        errors.Errors = append(errors.Errors, 
            errors.New("age must be positive"))
    }
    
    if user.Email == "" {
        errors.Errors = append(errors.Errors, 
            errors.New("email is required"))
    }
    
    if len(errors.Errors) > 0 {
        return errors
    }
    
    return nil
}
```

### Error Recovery

```go
// Recover from panics
func safeOperation() (result interface{}, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic occurred: %v", r)
        }
    }()
    
    // Potentially panicking operation
    result = riskyOperation()
    return result, nil
}
```

## Panic and Recover

### When to Use Panic

```go
// Panic for programming errors
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero") // Programming error
    }
    return a / b
}

// Panic for unrecoverable errors
func initialize() {
    if err := loadConfig(); err != nil {
        panic(fmt.Sprintf("failed to load config: %v", err))
    }
}
```

### When NOT to Use Panic

```go
// Don't panic for normal error conditions
func divideSafe(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Don't panic for user input errors
func validateInput(input string) error {
    if input == "" {
        return errors.New("input is required")
    }
    return nil
}
```

### Recover Pattern

```go
// Recover from panics in goroutines
func worker() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Worker panicked: %v", r)
        }
    }()
    
    // Worker logic that might panic
    riskyOperation()
}

// Recover in HTTP handlers
func httpHandler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            log.Printf("Handler panicked: %v", r)
        }
    }()
    
    // Handler logic
}
```

## Best Practices

### 1. Handle Errors Explicitly

```go
// Good: Explicit error handling
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()
    
    // Process file
    return nil
}

// Avoid: Ignoring errors
func processFileBad(filename string) {
    file, _ := os.Open(filename) // Ignoring error
    defer file.Close()
    
    // Process file
}
```

### 2. Wrap Errors with Context

```go
// Good: Wrap errors with context
func getUserData(userID int) (*UserData, error) {
    user, err := getUser(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to get user %d: %w", userID, err)
    }
    
    data, err := getData(user.ID)
    if err != nil {
        return nil, fmt.Errorf("failed to get data for user %d: %w", userID, err)
    }
    
    return data, nil
}

// Avoid: Losing error context
func getUserDataBad(userID int) (*UserData, error) {
    user, err := getUser(userID)
    if err != nil {
        return nil, err // Lost context
    }
    
    data, err := getData(user.ID)
    if err != nil {
        return nil, err // Lost context
    }
    
    return data, nil
}
```

### 3. Use Specific Error Types

```go
// Good: Specific error types
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// Usage
func validateUser(user User) error {
    if user.Name == "" {
        return ValidationError{Field: "name", Message: "required"}
    }
    return nil
}

// Avoid: Generic errors
func validateUserBad(user User) error {
    if user.Name == "" {
        return errors.New("validation error") // Too generic
    }
    return nil
}
```

### 4. Don't Panic for Normal Errors

```go
// Good: Return errors for normal conditions
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    return nil
}

// Avoid: Panicking for normal conditions
func processDataBad(data []byte) {
    if len(data) == 0 {
        panic("empty data") // Don't panic for normal errors
    }
}
```

### 5. Use errors.Is and errors.As

```go
// Good: Use errors.Is for error comparison
func handleError(err error) {
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File not found")
        return
    }
    
    if errors.Is(err, os.ErrPermission) {
        fmt.Println("Permission denied")
        return
    }
    
    fmt.Printf("Unknown error: %v\n", err)
}

// Good: Use errors.As for error type assertion
func handleValidationError(err error) {
    var validationErr ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Validation error on field %s: %s\n", 
            validationErr.Field, validationErr.Message)
    }
}
```

## Common Patterns

### 1. Error Middleware

```go
// HTTP error middleware
func errorMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                log.Printf("Panic in handler: %v", r)
            }
        }()
        
        next.ServeHTTP(w, r)
    })
}
```

### 2. Error Logging

```go
// Structured error logging
func logError(err error, context map[string]interface{}) {
    log.Printf("Error: %v, Context: %+v", err, context)
}

// Usage
func processRequest(req Request) error {
    if err := validateRequest(req); err != nil {
        logError(err, map[string]interface{}{
            "request_id": req.ID,
            "user_id":    req.UserID,
        })
        return err
    }
    return nil
}
```

### 3. Error Retry

```go
// Retry with exponential backoff
func retryOperation(operation func() error, maxRetries int) error {
    var err error
    for i := 0; i < maxRetries; i++ {
        err = operation()
        if err == nil {
            return nil
        }
        
        if i < maxRetries-1 {
            time.Sleep(time.Duration(i+1) * time.Second)
        }
    }
    return fmt.Errorf("operation failed after %d retries: %w", maxRetries, err)
}
```

### 4. Error Metrics

```go
// Error metrics collection
type ErrorMetrics struct {
    ErrorCounts map[string]int
    mu          sync.RWMutex
}

func (em *ErrorMetrics) RecordError(err error) {
    em.mu.Lock()
    defer em.mu.Unlock()
    
    errorType := fmt.Sprintf("%T", err)
    em.ErrorCounts[errorType]++
}

func (em *ErrorMetrics) GetErrorCount(errorType string) int {
    em.mu.RLock()
    defer em.mu.RUnlock()
    return em.ErrorCounts[errorType]
}
```

## References

- [Go Language Specification - Error Handling](https://golang.org/ref/spec#Errors)
- [Effective Go - Errors](https://golang.org/doc/effective_go.html#errors)
- [Go by Example - Errors](https://gobyexample.com/errors)
- [Go by Example - Panic](https://gobyexample.com/panic)
- [Go by Example - Recover](https://gobyexample.com/recover)

## Next Steps

After mastering error handling, continue with:
- [Testing](../testing/) - Learn Go's testing framework
- [Packages](../packages/) - Master Go's package system
