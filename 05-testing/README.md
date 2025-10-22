# Go Testing

This comprehensive guide covers Go's built-in testing framework, which provides powerful tools for writing unit tests, benchmarks, and examples. Go's testing package is designed to be simple yet effective, encouraging good testing practices.

## Table of Contents
1. [Basic Testing](#basic-testing)
2. [Test Functions](#test-functions)
3. [Test Organization](#test-organization)
4. [Table-Driven Tests](#table-driven-tests)
5. [Benchmarks](#benchmarks)
6. [Test Coverage](#test-coverage)
7. [Test Helpers](#test-helpers)
8. [Mocking and Stubbing](#mocking-and-stubbing)
9. [Integration Tests](#integration-tests)
10. [Best Practices](#best-practices)

## Basic Testing

### Test File Structure

```go
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("Subtract(5, 3) = %d; want 2", result)
    }
}
```

### Running Tests

```bash
# Run all tests in current package
go test

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run tests with coverage
go test -cover

# Run tests in all packages
go test ./...
```

### Test Functions

```go
// Basic test function
func TestFunctionName(t *testing.T) {
    // Test implementation
}

// Test with subtests
func TestMathOperations(t *testing.T) {
    t.Run("addition", func(t *testing.T) {
        result := Add(2, 3)
        if result != 5 {
            t.Errorf("Add(2, 3) = %d; want 5", result)
        }
    })
    
    t.Run("subtraction", func(t *testing.T) {
        result := Subtract(5, 3)
        if result != 2 {
            t.Errorf("Subtract(5, 3) = %d; want 2", result)
        }
    })
}
```

## Test Functions

### Basic Test Structure

```go
func TestAdd(t *testing.T) {
    // Arrange
    a, b := 2, 3
    expected := 5
    
    // Act
    result := Add(a, b)
    
    // Assert
    if result != expected {
        t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
    }
}
```

### Test with Error Cases

```go
func TestDivide(t *testing.T) {
    // Test successful division
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Divide(10, 2) returned error: %v", err)
    }
    if result != 5 {
        t.Errorf("Divide(10, 2) = %f; want 5", result)
    }
    
    // Test division by zero
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) should return error")
    }
}
```

### Test with Setup and Teardown

```go
func TestWithSetup(t *testing.T) {
    // Setup
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    // Test
    user, err := db.GetUser(1)
    if err != nil {
        t.Errorf("GetUser(1) returned error: %v", err)
    }
    if user == nil {
        t.Error("GetUser(1) returned nil user")
    }
}

func setupTestDB(t *testing.T) *Database {
    db := NewTestDatabase()
    if err := db.Migrate(); err != nil {
        t.Fatalf("Failed to setup test database: %v", err)
    }
    return db
}

func cleanupTestDB(t *testing.T, db *Database) {
    if err := db.Close(); err != nil {
        t.Errorf("Failed to cleanup test database: %v", err)
    }
}
```

## Test Organization

### Test Files

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    // Test implementation
}
```

### Test Packages

```go
// math_test.go (same package)
package math

import "testing"

func TestAdd(t *testing.T) {
    // Can access unexported functions
    result := add(2, 3)
    // ...
}

// math_test.go (separate package)
package math_test

import (
    "testing"
    "yourproject/math"
)

func TestAdd(t *testing.T) {
    // Can only access exported functions
    result := math.Add(2, 3)
    // ...
}
```

### Test Groups

```go
func TestMathOperations(t *testing.T) {
    tests := []struct {
        name string
        fn   func()
    }{
        {"addition", testAdd},
        {"subtraction", testSubtract},
        {"multiplication", testMultiply},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.fn()
        })
    }
}

func testAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

## Table-Driven Tests

### Basic Table-Driven Test

```go
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
        {"zero", 0, 5, 5},
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
```

### Table-Driven Test with Error Cases

```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name     string
        a        float64
        b        float64
        expected float64
        wantErr  bool
    }{
        {"normal division", 10, 2, 5, false},
        {"division by zero", 10, 0, 0, true},
        {"negative result", -10, 2, -5, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Divide(tt.a, tt.b)
            
            if tt.wantErr {
                if err == nil {
                    t.Errorf("Divide(%f, %f) expected error", tt.a, tt.b)
                }
                return
            }
            
            if err != nil {
                t.Errorf("Divide(%f, %f) returned error: %v", tt.a, tt.b, err)
                return
            }
            
            if result != tt.expected {
                t.Errorf("Divide(%f, %f) = %f; want %f", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Table-Driven Test with Complex Types

```go
func TestProcessUser(t *testing.T) {
    tests := []struct {
        name     string
        user     User
        expected ProcessedUser
        wantErr  bool
    }{
        {
            name: "valid user",
            user: User{Name: "Alice", Age: 30},
            expected: ProcessedUser{Name: "Alice", Age: 30, Status: "active"},
            wantErr: false,
        },
        {
            name: "invalid age",
            user: User{Name: "Bob", Age: -5},
            expected: ProcessedUser{},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := ProcessUser(tt.user)
            
            if tt.wantErr {
                if err == nil {
                    t.Errorf("ProcessUser() expected error")
                }
                return
            }
            
            if err != nil {
                t.Errorf("ProcessUser() returned error: %v", err)
                return
            }
            
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("ProcessUser() = %v; want %v", result, tt.expected)
            }
        })
    }
}
```

## Benchmarks

### Basic Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

### Benchmark with Setup

```go
func BenchmarkProcessData(b *testing.B) {
    data := generateTestData(1000)
    
    b.ResetTimer() // Reset timer after setup
    for i := 0; i < b.N; i++ {
        ProcessData(data)
    }
}
```

### Benchmark with Different Input Sizes

```go
func BenchmarkProcessData(b *testing.B) {
    sizes := []int{100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
            data := generateTestData(size)
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                ProcessData(data)
            }
        })
    }
}
```

### Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkAdd

# Run benchmark with memory profiling
go test -bench=BenchmarkAdd -benchmem

# Run benchmark multiple times
go test -bench=BenchmarkAdd -count=5
```

## Test Coverage

### Coverage Commands

```bash
# Run tests with coverage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in HTML
go tool cover -html=coverage.out

# View coverage in terminal
go tool cover -func=coverage.out
```

### Coverage Example

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

// Only Add is tested, so coverage will be 33%
```

## Test Helpers

### Basic Test Helper

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assertEqual(t, result, 5)
}
```

### Setup Helper

```go
func setupTestDB(t *testing.T) *Database {
    t.Helper()
    
    db := NewTestDatabase()
    if err := db.Migrate(); err != nil {
        t.Fatalf("Failed to setup test database: %v", err)
    }
    return db
}

func TestUserOperations(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    // Test user operations
}
```

### Assertion Helpers

```go
func assertNoError(t *testing.T, err error) {
    t.Helper()
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
}

func assertError(t *testing.T, err error) {
    t.Helper()
    if err == nil {
        t.Error("expected error, got nil")
    }
}

func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

## Mocking and Stubbing

### Interface Mocking

```go
// Define interface
type UserService interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

// Mock implementation
type MockUserService struct {
    users map[int]*User
    err   error
}

func (m *MockUserService) GetUser(id int) (*User, error) {
    if m.err != nil {
        return nil, m.err
    }
    return m.users[id], nil
}

func (m *MockUserService) SaveUser(user *User) error {
    if m.err != nil {
        return m.err
    }
    m.users[user.ID] = user
    return nil
}

// Test with mock
func TestProcessUser(t *testing.T) {
    mockService := &MockUserService{
        users: make(map[int]*User),
    }
    
    user := &User{ID: 1, Name: "Alice"}
    mockService.users[1] = user
    
    result, err := ProcessUser(mockService, 1)
    assertNoError(t, err)
    assertEqual(t, result.Name, "Alice")
}
```

### Using testify/mock

```go
import "github.com/stretchr/testify/mock"

type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) GetUser(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func TestProcessUser(t *testing.T) {
    mockService := new(MockUserService)
    user := &User{ID: 1, Name: "Alice"}
    
    mockService.On("GetUser", 1).Return(user, nil)
    
    result, err := ProcessUser(mockService, 1)
    
    assertNoError(t, err)
    assertEqual(t, result.Name, "Alice")
    mockService.AssertExpectations(t)
}
```

## Integration Tests

### Database Integration Test

```go
func TestUserIntegration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    
    db := setupIntegrationDB(t)
    defer cleanupIntegrationDB(t, db)
    
    // Test user operations
    user := &User{Name: "Alice", Email: "alice@example.com"}
    
    err := db.SaveUser(user)
    assertNoError(t, err)
    
    retrieved, err := db.GetUser(user.ID)
    assertNoError(t, err)
    assertEqual(t, retrieved.Name, "Alice")
}
```

### HTTP Integration Test

```go
func TestHTTPHandler(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(handleUser))
    defer server.Close()
    
    resp, err := http.Get(server.URL + "/user/1")
    assertNoError(t, err)
    defer resp.Body.Close()
    
    assertEqual(t, resp.StatusCode, http.StatusOK)
    
    var user User
    err = json.NewDecoder(resp.Body).Decode(&user)
    assertNoError(t, err)
    assertEqual(t, user.ID, 1)
}
```

## Best Practices

### 1. Test Naming

```go
// Good: Descriptive test names
func TestAdd_WithPositiveNumbers_ReturnsCorrectSum(t *testing.T) {
    // Test implementation
}

func TestDivide_WithZeroDenominator_ReturnsError(t *testing.T) {
    // Test implementation
}

// Avoid: Unclear test names
func TestAdd(t *testing.T) {
    // Test implementation
}
```

### 2. Test Structure

```go
// Good: Arrange-Act-Assert pattern
func TestAdd(t *testing.T) {
    // Arrange
    a, b := 2, 3
    expected := 5
    
    // Act
    result := Add(a, b)
    
    // Assert
    if result != expected {
        t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
    }
}
```

### 3. Test One Thing

```go
// Good: Test one thing per test
func TestAdd_WithPositiveNumbers(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestAdd_WithNegativeNumbers(t *testing.T) {
    result := Add(-2, -3)
    if result != -5 {
        t.Errorf("Add(-2, -3) = %d; want -5", result)
    }
}

// Avoid: Testing multiple things
func TestAdd(t *testing.T) {
    // Test positive numbers
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
    
    // Test negative numbers
    result = Add(-2, -3)
    if result != -5 {
        t.Errorf("Add(-2, -3) = %d; want -5", result)
    }
}
```

### 4. Use Table-Driven Tests

```go
// Good: Table-driven test
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
```

### 5. Test Error Cases

```go
// Good: Test both success and error cases
func TestDivide(t *testing.T) {
    // Success case
    result, err := Divide(10, 2)
    if err != nil {
        t.Errorf("Divide(10, 2) returned error: %v", err)
    }
    if result != 5 {
        t.Errorf("Divide(10, 2) = %f; want 5", result)
    }
    
    // Error case
    _, err = Divide(10, 0)
    if err == nil {
        t.Error("Divide(10, 0) should return error")
    }
}
```

## References

- [Go Testing Package](https://golang.org/pkg/testing/)
- [Effective Go - Testing](https://golang.org/doc/effective_go.html#testing)
- [Go by Example - Testing](https://gobyexample.com/testing)
- [Go by Example - Benchmarking](https://gobyexample.com/benchmarking)

## Next Steps

After mastering testing, continue with:
- [Packages](../packages/) - Learn Go's package system
- [Best Practices](../best-practices/) - Master Go's best practices
