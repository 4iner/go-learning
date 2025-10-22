package main

import (
	"fmt"
	"strings"
	"time"
)

// This example demonstrates Go's function system
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Functions Examples ===")
	
	// Demonstrate basic functions
	demonstrateBasicFunctions()
	
	// Demonstrate function parameters
	demonstrateFunctionParameters()
	
	// Demonstrate return values
	demonstrateReturnValues()
	
	// Demonstrate variadic functions
	demonstrateVariadicFunctions()
	
	// Demonstrate anonymous functions
	demonstrateAnonymousFunctions()
	
	// Demonstrate closures
	demonstrateClosures()
	
	// Demonstrate higher-order functions
	demonstrateHigherOrderFunctions()
	
	// Demonstrate method receivers
	demonstrateMethodReceivers()
	
	// Demonstrate function types
	demonstrateFunctionTypes()
}

// demonstrateBasicFunctions shows basic function declarations and calls
func demonstrateBasicFunctions() {
	fmt.Println("\n1. Basic Functions:")
	
	// Simple function call
	result := add(5, 3)
	fmt.Printf("   add(5, 3) = %d\n", result)
	
	// Function with multiple parameters
	product := multiply(4, 6)
	fmt.Printf("   multiply(4, 6) = %d\n", product)
	
	// Function with multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   divide(10, 2) = %d\n", quotient)
	}
	
	// Function with error
	quotient2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("   divide(10, 0) error: %v\n", err2)
	} else {
		fmt.Printf("   divide(10, 0) = %d\n", quotient2)
	}
}

// demonstrateFunctionParameters shows different parameter types
func demonstrateFunctionParameters() {
	fmt.Println("\n2. Function Parameters:")
	
	// Value parameters
	x := 5
	fmt.Printf("   Before modifyValue: x = %d\n", x)
	modifyValue(x)
	fmt.Printf("   After modifyValue: x = %d\n", x)
	
	// Pointer parameters
	fmt.Printf("   Before modifyPointer: x = %d\n", x)
	modifyPointer(&x)
	fmt.Printf("   After modifyPointer: x = %d\n", x)
	
	// Slice parameters (reference type)
	slice := []int{1, 2, 3}
	fmt.Printf("   Before modifySlice: %v\n", slice)
	modifySlice(slice)
	fmt.Printf("   After modifySlice: %v\n", slice)
	
	// Map parameters (reference type)
	m := map[string]int{"a": 1, "b": 2}
	fmt.Printf("   Before modifyMap: %v\n", m)
	modifyMap(m)
	fmt.Printf("   After modifyMap: %v\n", m)
	
	// Struct parameters
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("   Before updatePerson: %+v\n", person)
	updatePerson(&person, "Bob", 25)
	fmt.Printf("   After updatePerson: %+v\n", person)
}

// demonstrateReturnValues shows different return value patterns
func demonstrateReturnValues() {
	fmt.Println("\n3. Return Values:")
	
	// Single return value
	squared := square(5)
	fmt.Printf("   square(5) = %d\n", squared)
	
	// Multiple return values
	min, max := getMinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	fmt.Printf("   getMinMax([3,1,4,1,5,9,2,6]) = min: %d, max: %d\n", min, max)
	
	// Named return values
	result, err := divideNamed(20, 4)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   divideNamed(20, 4) = %d\n", result)
	}
	
	// Error handling pattern
	content, err := readFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("   readFile error: %v\n", err)
	} else {
		fmt.Printf("   File content: %s\n", content)
	}
}

// demonstrateVariadicFunctions shows variadic function usage
func demonstrateVariadicFunctions() {
	fmt.Println("\n4. Variadic Functions:")
	
	// Basic variadic function
	result1 := sum(1, 2, 3, 4, 5)
	fmt.Printf("   sum(1, 2, 3, 4, 5) = %d\n", result1)
	
	result2 := sum(10, 20)
	fmt.Printf("   sum(10, 20) = %d\n", result2)
	
	result3 := sum()
	fmt.Printf("   sum() = %d\n", result3)
	
	// Passing slice to variadic function
	numbers := []int{1, 2, 3, 4, 5}
	result4 := sum(numbers...)
	fmt.Printf("   sum([]int{1,2,3,4,5}...) = %d\n", result4)
	
	// Variadic function with different types
	fmt.Println("   printValues(1, \"hello\", true, 3.14):")
	printValues(1, "hello", true, 3.14, []int{1, 2, 3})
	
	// Variadic function with mixed parameters
	result5 := formatString("Result: ", 1, 2, 3)
	fmt.Printf("   formatString(\"Result: \", 1, 2, 3) = %s\n", result5)
}

// demonstrateAnonymousFunctions shows anonymous function usage
func demonstrateAnonymousFunctions() {
	fmt.Println("\n5. Anonymous Functions:")
	
	// Anonymous function assigned to variable
	add := func(a, b int) int {
		return a + b
	}
	result := add(5, 3)
	fmt.Printf("   Anonymous add(5, 3) = %d\n", result)
	
	// Anonymous function called immediately
	result2 := func(a, b int) int {
		return a * b
	}(4, 6)
	fmt.Printf("   Immediate anonymous function (4, 6) = %d\n", result2)
	
	// Anonymous function as parameter
	numbers := []int{1, 2, 3, 4, 5}
	squared := processNumbers(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("   Squared numbers: %v\n", squared)
	
	doubled := processNumbers(numbers, func(x int) int {
		return x * 2
	})
	fmt.Printf("   Doubled numbers: %v\n", doubled)
}

// demonstrateClosures shows closure usage
func demonstrateClosures() {
	fmt.Println("\n6. Closures:")
	
	// Basic closure
	counter := createCounter()
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
	fmt.Printf("   Counter: %d\n", counter())
	
	// New counter (independent)
	counter2 := createCounter()
	fmt.Printf("   New counter: %d\n", counter2())
	fmt.Printf("   Original counter: %d\n", counter())
	
	// Closure with parameters
	double := createMultiplier(2)
	triple := createMultiplier(3)
	fmt.Printf("   double(5) = %d\n", double(5))
	fmt.Printf("   triple(5) = %d\n", triple(5))
	
	// Closure in loops (correct way)
	fmt.Println("   Closures in loops:")
	funcs := createClosures(3)
	for i, f := range funcs {
		fmt.Printf("     Function %d: %d\n", i, f())
	}
}

// demonstrateHigherOrderFunctions shows higher-order function usage
func demonstrateHigherOrderFunctions() {
	fmt.Println("\n7. Higher-Order Functions:")
	
	// Function as parameter
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	evens := filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("   Even numbers: %v\n", evens)
	
	greaterThan5 := filter(numbers, func(x int) bool {
		return x > 5
	})
	fmt.Printf("   Numbers > 5: %v\n", greaterThan5)
	
	// Function as return value
	validateAge := createValidator(0, 120)
	validateScore := createValidator(0, 100)
	
	fmt.Printf("   validateAge(25) = %t\n", validateAge(25))
	fmt.Printf("   validateAge(150) = %t\n", validateAge(150))
	fmt.Printf("   validateScore(85) = %t\n", validateScore(85))
	fmt.Printf("   validateScore(150) = %t\n", validateScore(150))
	
	// Function composition
	square := func(x int) int { return x * x }
	double := func(x int) int { return x * 2 }
	
	squareThenDouble := compose(double, square)
	doubleThenSquare := compose(square, double)
	
	fmt.Printf("   squareThenDouble(3) = %d\n", squareThenDouble(3))
	fmt.Printf("   doubleThenSquare(3) = %d\n", doubleThenSquare(3))
}

// demonstrateMethodReceivers shows method receiver usage
func demonstrateMethodReceivers() {
	fmt.Println("\n8. Method Receivers:")
	
	// Value receiver
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %.2f\n", rect.Area())
	fmt.Printf("   Perimeter: %.2f\n", rect.Perimeter())
	
	// Value receiver (doesn't modify original)
	scaled := rect.Scale(2)
	fmt.Printf("   Original after Scale: %+v\n", rect)
	fmt.Printf("   Scaled rectangle: %+v\n", scaled)
	
	// Pointer receiver (modifies original)
	rect.ScaleInPlace(2)
	fmt.Printf("   After ScaleInPlace(2): %+v\n", rect)
	
	// Method sets
	var t T
	var p *T = &t
	
	fmt.Println("   Method sets:")
	t.method1()  // OK
	t.method2()  // OK (Go automatically takes address)
	p.method1()  // OK (Go automatically dereferences)
	p.method2()  // OK
}

// demonstrateFunctionTypes shows function type usage
func demonstrateFunctionTypes() {
	fmt.Println("\n9. Function Types:")
	
	// Function type declarations
	add := func(a, b int) int { return a + b }
	multiply := func(a, b int) int { return a * b }
	
	result1 := applyOperation(5, 3, add)
	result2 := applyOperation(5, 3, multiply)
	
	fmt.Printf("   applyOperation(5, 3, add) = %d\n", result1)
	fmt.Printf("   applyOperation(5, 3, multiply) = %d\n", result2)
	
	// Function type interfaces
	var processor Processor = FuncProcessor(func(x int) int {
		return x * 2
	})
	
	result3 := processor.Process(5)
	fmt.Printf("   processor.Process(5) = %d\n", result3)
}

// Basic function implementations
func add(a, b int) int {
	return a + b
}

func multiply(x, y int) int {
	return x * y
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func square(x int) int {
	return x * x
}

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

func divideNamed(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}

func readFile(filename string) (string, error) {
	// Simulate file reading
	if filename == "nonexistent.txt" {
		return "", fmt.Errorf("file not found: %s", filename)
	}
	return "file content", nil
}

func modifyValue(x int) {
	x = 100
}

func modifyPointer(x *int) {
	*x = 100
}

func modifySlice(s []int) {
	s[0] = 100
}

func modifyMap(m map[string]int) {
	m["key"] = 100
}

func updatePerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func printValues(values ...interface{}) {
	for i, value := range values {
		fmt.Printf("     Value %d: %v (type: %T)\n", i, value, value)
	}
}

func formatString(prefix string, values ...interface{}) string {
	var parts []string
	for _, value := range values {
		parts = append(parts, fmt.Sprintf("%v", value))
	}
	return prefix + strings.Join(parts, " ")
}

func processNumbers(numbers []int, processor func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = processor(num)
	}
	return result
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func createClosures(count int) []func() int {
	var funcs []func() int
	for i := 0; i < count; i++ {
		i := i  // Create new variable for each iteration
		funcs = append(funcs, func() int {
			return i
		})
	}
	return funcs
}

func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func createValidator(min, max int) func(int) bool {
	return func(value int) bool {
		return value >= min && value <= max
	}
}

func compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func applyOperation(a, b int, op BinaryOp) int {
	return op(a, b)
}

// Type definitions
type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Width  float64
	Height float64
}

type T struct{}

type BinaryOp func(int, int) int
type Processor interface {
	Process(int) int
}
type FuncProcessor func(int) int

// Method implementations
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Scale(factor float64) Rectangle {
	return Rectangle{
		Width:  r.Width * factor,
		Height: r.Height * factor,
	}
}

func (r *Rectangle) ScaleInPlace(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (T) method1() {
	fmt.Println("     method1 called")
}

func (*T) method2() {
	fmt.Println("     method2 called")
}

func (f FuncProcessor) Process(x int) int {
	return f(x)
}
