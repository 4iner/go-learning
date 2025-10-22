package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// This example demonstrates Go's error handling patterns
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Error Handling Examples ===")
	
	// Demonstrate basic error handling
	demonstrateBasicErrorHandling()
	
	// Demonstrate custom error types
	demonstrateCustomErrorTypes()
	
	// Demonstrate error wrapping
	demonstrateErrorWrapping()
	
	// Demonstrate error checking patterns
	demonstrateErrorCheckingPatterns()
	
	// Demonstrate panic and recover
	demonstratePanicRecover()
	
	// Demonstrate common patterns
	demonstrateCommonPatterns()
}

// demonstrateBasicErrorHandling shows basic error handling
func demonstrateBasicErrorHandling() {
	fmt.Println("\n1. Basic Error Handling:")
	
	// Basic error creation
	err := errors.New("something went wrong")
	fmt.Printf("   Basic error: %v\n", err)
	
	// Formatted error
	err = fmt.Errorf("user %s not found", "alice")
	fmt.Printf("   Formatted error: %v\n", err)
	
	// Function error returns
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %d\n", result)
	}
	
	// Error with division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %d\n", result)
	}
	
	// Error handling with defer
	fmt.Println("   Error handling with defer:")
	if err := processFile("example.txt"); err != nil {
		fmt.Printf("     Error: %v\n", err)
	}
}

// demonstrateCustomErrorTypes shows custom error types
func demonstrateCustomErrorTypes() {
	fmt.Println("\n2. Custom Error Types:")
	
	// Validation error
	user := User{Name: "", Age: -5}
	if err := validateUser(user); err != nil {
		fmt.Printf("   Validation error: %v\n", err)
		
		// Type assertion
		if validationErr, ok := err.(ValidationError); ok {
			fmt.Printf("     Field: %s, Message: %s\n", validationErr.Field, validationErr.Message)
		}
	}
	
	// Database error
	if err := saveUser(user); err != nil {
		fmt.Printf("   Database error: %v\n", err)
		
		// Type assertion
		if dbErr, ok := err.(DatabaseError); ok {
			fmt.Printf("     Operation: %s, Table: %s\n", dbErr.Operation, dbErr.Table)
		}
	}
	
	// App error with codes
	if _, err := getUser(123); err != nil {
		fmt.Printf("   App error: %v\n", err)
		
		// Type assertion
		if appErr, ok := err.(AppError); ok {
			fmt.Printf("     Code: %d, Message: %s\n", appErr.Code, appErr.Message)
		}
	}
}

// demonstrateErrorWrapping shows error wrapping
func demonstrateErrorWrapping() {
	fmt.Println("\n3. Error Wrapping:")
	
	// Error wrapping with context
	if err := processUser(123); err != nil {
		fmt.Printf("   Wrapped error: %v\n", err)
		
		// Error unwrapping
		fmt.Println("   Error chain:")
		for err != nil {
			fmt.Printf("     %v\n", err)
			err = errors.Unwrap(err)
		}
	}
	
	// Error inspection
	fmt.Println("\n   Error inspection:")
	err := processUser(456)
	if err != nil {
		inspectError(err)
	}
}

// demonstrateErrorCheckingPatterns shows error checking patterns
func demonstrateErrorCheckingPatterns() {
	fmt.Println("\n4. Error Checking Patterns:")
	
	// Early return pattern
	data := []byte("test data")
	if result, err := processData(data); err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %s\n", string(result))
	}
	
	// Error aggregation
	user := User{Name: "", Age: -5, Email: ""}
	if err := validateUserMultiple(user); err != nil {
		fmt.Printf("   Multiple errors: %v\n", err)
	}
	
	// Error recovery
	result, err := safeOperation()
	if err != nil {
		fmt.Printf("   Recovered error: %v\n", err)
	} else {
		fmt.Printf("   Safe operation result: %v\n", result)
	}
}

// demonstratePanicRecover shows panic and recover
func demonstratePanicRecover() {
	fmt.Println("\n5. Panic and Recover:")
	
	// Panic for programming errors
	fmt.Println("   Panic for programming errors:")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("     Recovered from panic: %v\n", r)
		}
	}()
	
	// This will panic
	dividePanic(10, 0)
	
	// Panic in goroutine
	fmt.Println("   Panic in goroutine:")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("     Goroutine recovered from panic: %v\n", r)
			}
		}()
		
		panic("goroutine panic")
	}()
	
	time.Sleep(100 * time.Millisecond)
}

// demonstrateCommonPatterns shows common error patterns
func demonstrateCommonPatterns() {
	fmt.Println("\n6. Common Patterns:")
	
	// Error middleware pattern
	fmt.Println("   Error middleware pattern:")
	handler := errorMiddleware(httpHandler)
	handler.ServeHTTP(nil, nil)
	
	// Error logging
	fmt.Println("   Error logging:")
	logError(errors.New("test error"), map[string]interface{}{
		"request_id": "123",
		"user_id":    "456",
	})
	
	// Error retry
	fmt.Println("   Error retry:")
	if err := retryOperation(func() error {
		return errors.New("temporary error")
	}, 3); err != nil {
		fmt.Printf("     Retry failed: %v\n", err)
	}
	
	// Error metrics
	fmt.Println("   Error metrics:")
	metrics := &ErrorMetrics{ErrorCounts: make(map[string]int)}
	metrics.RecordError(errors.New("test error"))
	metrics.RecordError(errors.New("test error"))
	fmt.Printf("     Error count: %d\n", metrics.GetErrorCount("*errors.errorString"))
}

// Helper functions
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()
	
	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	
	fmt.Printf("     File content: %s\n", string(content))
	return nil
}

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

func saveUser(user User) error {
	// Simulate database error
	return DatabaseError{
		Operation: "save",
		Table:     "users",
		Err:       errors.New("connection failed"),
	}
}

func getUser(id int) (*User, error) {
	// Simulate user not found
	return nil, AppError{
		Code:    ErrNotFound,
		Message: "user not found",
		Err:     errors.New("database query failed"),
	}
}

func processUser(userID int) error {
	user, err := getUser(userID)
	if err != nil {
		return fmt.Errorf("failed to get user %d: %w", userID, err)
	}
	
	if err := validateUser(*user); err != nil {
		return fmt.Errorf("failed to validate user %d: %w", userID, err)
	}
	
	return nil
}

func inspectError(err error) {
	// Check for specific error types
	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("     Validation error on field %s: %s\n", 
			validationErr.Field, validationErr.Message)
		return
	}
	
	// Check for wrapped errors
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("     File not found")
		return
	}
	
	fmt.Printf("     Unknown error: %v\n", err)
}

func processData(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}
	
	if !isValid(data) {
		return nil, errors.New("invalid data")
	}
	
	return data, nil
}

func isValid(data []byte) bool {
	return len(data) > 0
}

func validateUserMultiple(user User) error {
	var errors MultiError
	
	if user.Name == "" {
		errors.Errors = append(errors.Errors, errors.New("name is required"))
	}
	
	if user.Age < 0 {
		errors.Errors = append(errors.Errors, errors.New("age must be positive"))
	}
	
	if user.Email == "" {
		errors.Errors = append(errors.Errors, errors.New("email is required"))
	}
	
	if len(errors.Errors) > 0 {
		return errors
	}
	
	return nil
}

func safeOperation() (result interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
		}
	}()
	
	// Simulate risky operation
	result = riskyOperation()
	return result, nil
}

func riskyOperation() interface{} {
	// Simulate panic
	panic("risky operation failed")
}

func dividePanic(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func errorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("     Panic in handler: %v\n", r)
			}
		}()
		
		next.ServeHTTP(w, r)
	})
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate panic
	panic("handler panic")
}

func logError(err error, context map[string]interface{}) {
	fmt.Printf("     Error: %v, Context: %+v\n", err, context)
}

func retryOperation(operation func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		
		if i < maxRetries-1 {
			time.Sleep(time.Duration(i+1) * time.Millisecond)
		}
	}
	return fmt.Errorf("operation failed after %d retries: %w", maxRetries, err)
}

// Type definitions
type User struct {
	Name  string
	Age   int
	Email string
}

type ValidationError struct {
	Field   string
	Message string
}

type DatabaseError struct {
	Operation string
	Table     string
	Err       error
}

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

type MultiError struct {
	Errors []error
}

type ErrorMetrics struct {
	ErrorCounts map[string]int
	mu          sync.RWMutex
}

type http.Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type http.ResponseWriter interface{}
type http.Request struct{}

// Method implementations
func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error during %s on table %s: %v", 
		e.Operation, e.Table, e.Err)
}

func (e DatabaseError) Unwrap() error {
	return e.Err
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

func (e MultiError) Error() string {
	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
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
