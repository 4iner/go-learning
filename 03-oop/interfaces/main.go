package main

import (
	"fmt"
	"sort"
	"strings"
)

// This example demonstrates Go's interface system
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Interfaces Examples ===")
	
	// Demonstrate basic interfaces
	demonstrateBasicInterfaces()
	
	// Demonstrate interface implementation
	demonstrateInterfaceImplementation()
	
	// Demonstrate empty interface
	demonstrateEmptyInterface()
	
	// Demonstrate interface composition
	demonstrateInterfaceComposition()
	
	// Demonstrate type assertions
	demonstrateTypeAssertions()
	
	// Demonstrate type switches
	demonstrateTypeSwitches()
	
	// Demonstrate common interface patterns
	demonstrateCommonPatterns()
	
	// Demonstrate advanced concepts
	demonstrateAdvancedConcepts()
}

// demonstrateBasicInterfaces shows basic interface declarations
func demonstrateBasicInterfaces() {
	fmt.Println("\n1. Basic Interfaces:")
	
	// Basic interface usage
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}
	
	fmt.Printf("   Rectangle area: %.2f\n", rect.Area())
	fmt.Printf("   Circle area: %.2f\n", circle.Area())
	
	// Using interface
	printShapeInfo(rect)
	printShapeInfo(circle)
}

// demonstrateInterfaceImplementation shows interface implementation
func demonstrateInterfaceImplementation() {
	fmt.Println("\n2. Interface Implementation:")
	
	// Multiple interface implementation
	point := Point{X: 0, Y: 0}
	
	fmt.Println("   Drawing point:")
	drawShape(point)
	
	fmt.Println("   Moving point:")
	moveShape(&point, 5, 10)
	drawShape(point)
	
	// Interface satisfaction
	var drawable Drawable = point
	var movable Movable = &point
	
	drawable.Draw()
	movable.Move(2, 3)
	drawable.Draw()
}

// demonstrateEmptyInterface shows empty interface usage
func demonstrateEmptyInterface() {
	fmt.Println("\n3. Empty Interface:")
	
	// Empty interface can hold any type
	values := []interface{}{42, "hello", 3.14, true, []int{1, 2, 3}}
	
	for i, value := range values {
		fmt.Printf("   Value %d: %v (type: %T)\n", i, value, value)
	}
	
	// Type assertions with empty interface
	fmt.Println("\n   Type assertions:")
	for _, value := range values {
		processValue(value)
	}
}

// demonstrateInterfaceComposition shows interface composition
func demonstrateInterfaceComposition() {
	fmt.Println("\n4. Interface Composition:")
	
	// File implementing multiple interfaces
	file := &File{name: "test.txt"}
	
	// File implements Reader, Writer, and Closer
	var reader Reader = file
	var writer Writer = file
	var closer Closer = file
	
	// File also implements composed interfaces
	var readWriter ReadWriter = file
	var readWriteCloser ReadWriteCloser = file
	
	fmt.Printf("   File name: %s\n", file.name)
	
	// Demonstrate interface methods
	data := make([]byte, 10)
	n, err := reader.Read(data)
	if err == nil {
		fmt.Printf("   Read %d bytes\n", n)
	}
	
	n, err = writer.Write([]byte("Hello"))
	if err == nil {
		fmt.Printf("   Wrote %d bytes\n", n)
	}
	
	err = closer.Close()
	if err == nil {
		fmt.Println("   File closed successfully")
	}
	
	// Using composed interfaces
	fmt.Printf("   ReadWriter: %T\n", readWriter)
	fmt.Printf("   ReadWriteCloser: %T\n", readWriteCloser)
}

// demonstrateTypeAssertions shows type assertion usage
func demonstrateTypeAssertions() {
	fmt.Println("\n5. Type Assertions:")
	
	// Basic type assertions
	var value interface{} = "hello"
	
	str := value.(string)
	fmt.Printf("   String: %s\n", str)
	
	// Safe type assertions
	value = 42
	if str, ok := value.(string); ok {
		fmt.Printf("   String: %s\n", str)
	} else {
		fmt.Printf("   Not a string: %T\n", value)
	}
	
	// Type assertions with interfaces
	value = Rectangle{Width: 5, Height: 3}
	if shape, ok := value.(Shape); ok {
		fmt.Printf("   Shape area: %.2f\n", shape.Area())
	} else {
		fmt.Printf("   Not a shape: %T\n", value)
	}
	
	// Multiple type assertions
	values := []interface{}{"hello", 42, Rectangle{Width: 2, Height: 3}}
	for _, v := range values {
		processValueMultiple(v)
	}
}

// demonstrateTypeSwitches shows type switch usage
func demonstrateTypeSwitches() {
	fmt.Println("\n6. Type Switches:")
	
	// Basic type switch
	values := []interface{}{42, "hello", 3.14, true, Rectangle{Width: 4, Height: 5}}
	
	for _, value := range values {
		processValueSwitch(value)
	}
	
	// Type switch with interfaces
	fmt.Println("\n   Shape type switch:")
	shapes := []interface{}{Rectangle{Width: 3, Height: 4}, Circle{Radius: 2}}
	
	for _, shape := range shapes {
		processShapeSwitch(shape)
	}
}

// demonstrateCommonPatterns shows common interface patterns
func demonstrateCommonPatterns() {
	fmt.Println("\n7. Common Interface Patterns:")
	
	// Stringer interface
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("   Person: %s\n", person)
	
	// Error interface
	err := validateUser(User{Name: "", Age: 25})
	if err != nil {
		fmt.Printf("   Validation error: %s\n", err)
	}
	
	// sort.Interface
	numbers := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("   Before sort: %v\n", numbers)
	sort.Sort(numbers)
	fmt.Printf("   After sort: %v\n", numbers)
	
	// Custom sorting
	people := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}
	
	fmt.Printf("   Before sort: %v\n", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("   After sort by age: %v\n", people)
}

// demonstrateAdvancedConcepts shows advanced interface concepts
func demonstrateAdvancedConcepts() {
	fmt.Println("\n8. Advanced Interface Concepts:")
	
	// Interface values
	var w Writer
	fmt.Printf("   w is nil: %t\n", w == nil)
	
	w = &File{name: "test.txt"}
	fmt.Printf("   w is nil: %t\n", w == nil)
	fmt.Printf("   w type: %T\n", w)
	
	// Interface comparison
	var w1, w2 Writer
	w1 = &File{name: "test.txt"}
	w2 = &File{name: "test.txt"}
	
	fmt.Printf("   w1 == w2: %t\n", w1 == w2)
	
	// Method sets
	var t T
	var p *T = &t
	
	fmt.Println("   Method sets:")
	t.method1()
	p.method1()
	p.method2()
	
	// Interface with value receiver
	var i1 interface{ method1() } = t
	i1.method1()
	
	// Interface with pointer receiver
	var i2 interface{ method2() } = p
	i2.method2()
}

// Helper functions
func printShapeInfo(s Shape) {
	fmt.Printf("   Shape info - Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func drawShape(d Drawable) {
	d.Draw()
}

func moveShape(m Movable, x, y float64) {
	m.Move(x, y)
}

func processValue(value interface{}) {
	if str, ok := value.(string); ok {
		fmt.Printf("     String: %s (length: %d)\n", str, len(str))
	} else if num, ok := value.(int); ok {
		fmt.Printf("     Integer: %d\n", num)
	} else if slice, ok := value.([]int); ok {
		fmt.Printf("     Slice: %v (length: %d)\n", slice, len(slice))
	} else {
		fmt.Printf("     Unknown type: %T\n", value)
	}
}

func processValueMultiple(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("     String: %s\n", v)
	case int:
		fmt.Printf("     Integer: %d\n", v)
	case Shape:
		fmt.Printf("     Shape area: %.2f\n", v.Area())
	default:
		fmt.Printf("     Unknown type: %T\n", v)
	}
}

func processValueSwitch(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("     Integer: %d\n", v)
	case string:
		fmt.Printf("     String: %s\n", v)
	case bool:
		fmt.Printf("     Boolean: %t\n", v)
	case float64:
		fmt.Printf("     Float: %.2f\n", v)
	default:
		fmt.Printf("     Unknown type: %T\n", v)
	}
}

func processShapeSwitch(value interface{}) {
	switch s := value.(type) {
	case Rectangle:
		fmt.Printf("     Rectangle: %.2f x %.2f\n", s.Width, s.Height)
	case Circle:
		fmt.Printf("     Circle: radius %.2f\n", s.Radius)
	case Shape:
		fmt.Printf("     Shape area: %.2f\n", s.Area())
	default:
		fmt.Printf("     Not a shape: %T\n", s)
	}
}

func validateUser(user User) error {
	if user.Name == "" {
		return ValidationError{Field: "name", Message: "name is required"}
	}
	if user.Age < 0 {
		return ValidationError{Field: "age", Message: "age must be positive"}
	}
	return nil
}

// Interface definitions
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Drawable interface {
	Draw()
}

type Movable interface {
	Move(x, y float64)
}

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// Type definitions
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Point struct {
	X, Y float64
}

type File struct {
	name string
}

type Person struct {
	Name string
	Age  int
}

type User struct {
	Name string
	Age  int
}

type ValidationError struct {
	Field   string
	Message string
}

type IntSlice []int

type T struct{}

// Method implementations
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (p Point) Draw() {
	fmt.Printf("     Drawing point at (%.2f, %.2f)\n", p.X, p.Y)
}

func (p *Point) Move(x, y float64) {
	p.X += x
	p.Y += y
}

func (f *File) Read(data []byte) (int, error) {
	// Simulate reading
	return len(data), nil
}

func (f *File) Write(data []byte) (int, error) {
	// Simulate writing
	return len(data), nil
}

func (f *File) Close() error {
	// Simulate closing
	return nil
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (T) method1() {
	fmt.Println("     method1 called")
}

func (*T) method2() {
	fmt.Println("     method2 called")
}
