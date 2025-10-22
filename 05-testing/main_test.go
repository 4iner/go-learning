package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// This example demonstrates Go's testing framework
// Run this with: go test -v

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

func TestAddTableDriven(t *testing.T) {
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

func TestDivideTableDriven(t *testing.T) {
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

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func BenchmarkProcessData(b *testing.B) {
	data := generateTestData(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ProcessData(data)
	}
}

func BenchmarkProcessDataSizes(b *testing.B) {
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

func TestWithHelpers(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5)
	
	_, err := Divide(10, 0)
	assertError(t, err)
	
	result2, err2 := Divide(10, 2)
	assertNoError(t, err2)
	assertEqual(t, result2, 5.0)
}

func TestWithMock(t *testing.T) {
	mockService := &MockUserService{
		users: make(map[int]*User),
	}
	
	user := &User{ID: 1, Name: "Alice"}
	mockService.users[1] = user
	
	result, err := ProcessUserWithService(mockService, 1)
	assertNoError(t, err)
	assertEqual(t, result.Name, "Alice")
}

// Helper functions
func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

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

// Type definitions
type User struct {
	ID   int
	Name string
	Age  int
}

type ProcessedUser struct {
	Name   string
	Age    int
	Status string
}

type UserService interface {
	GetUser(id int) (*User, error)
}

type MockUserService struct {
	users map[int]*User
	err   error
}

// Method implementations
func (m *MockUserService) GetUser(id int) (*User, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.users[id], nil
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func ProcessUser(user User) (ProcessedUser, error) {
	if user.Age < 0 {
		return ProcessedUser{}, errors.New("invalid age")
	}
	
	return ProcessedUser{
		Name:   user.Name,
		Age:    user.Age,
		Status: "active",
	}, nil
}

func ProcessUserWithService(service UserService, id int) (*User, error) {
	return service.GetUser(id)
}

func generateTestData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func ProcessData(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}
