# Go Interfaces

This comprehensive guide covers Go's interface system, which is one of the most powerful and unique features of the language. Interfaces in Go provide a way to define behavior without specifying implementation, enabling polymorphism and clean abstractions.

## Table of Contents
1. [What are Interfaces?](#what-are-interfaces)
2. [Basic Interface Declaration](#basic-interface-declaration)
3. [Interface Implementation](#interface-implementation)
4. [Empty Interface](#empty-interface)
5. [Interface Composition](#interface-composition)
6. [Type Assertions](#type-assertions)
7. [Type Switches](#type-switches)
8. [Interface Best Practices](#interface-best-practices)
9. [Common Interface Patterns](#common-interface-patterns)
10. [Advanced Interface Concepts](#advanced-interface-concepts)

## What are Interfaces?

An interface in Go is a collection of method signatures. A type implements an interface by implementing all the methods defined in the interface. Unlike other languages, Go interfaces are **implicitly implemented** - you don't need to explicitly declare that a type implements an interface.

### Key Concepts

- **Implicit Implementation**: Types implement interfaces automatically
- **Duck Typing**: "If it walks like a duck and quacks like a duck, it's a duck"
- **Interface Satisfaction**: A type satisfies an interface if it implements all methods
- **Polymorphism**: Different types can be used interchangeably through interfaces

## Basic Interface Declaration

### Simple Interface

```go
// Basic interface declaration
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// Interface with multiple methods
type ReadWriter interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}
```

### Interface with Parameters and Return Values

```go
// Interface with complex method signatures
type Calculator interface {
    Add(a, b int) int
    Subtract(a, b int) int
    Multiply(a, b int) int
    Divide(a, b int) (int, error)
}

// Interface with different parameter types
type Storage interface {
    Save(key string, value interface{}) error
    Load(key string) (interface{}, error)
    Delete(key string) error
    List() ([]string, error)
}
```

## Interface Implementation

### Implementing Interfaces

```go
// Define interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Implement interface with Rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Implement interface with Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Usage
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    circle := Circle{Radius: 3}
    
    printShapeInfo(rect)   // Area: 50.00, Perimeter: 30.00
    printShapeInfo(circle) // Area: 28.27, Perimeter: 18.85
}
```

### Multiple Interface Implementation

```go
// Multiple interfaces
type Drawable interface {
    Draw()
}

type Movable interface {
    Move(x, y float64)
}

type DrawableMovable interface {
    Drawable
    Movable
}

// Type implementing multiple interfaces
type Point struct {
    X, Y float64
}

func (p Point) Draw() {
    fmt.Printf("Drawing point at (%.2f, %.2f)\n", p.X, p.Y)
}

func (p *Point) Move(x, y float64) {
    p.X += x
    p.Y += y
}

// Usage
func drawShape(d Drawable) {
    d.Draw()
}

func moveShape(m Movable, x, y float64) {
    m.Move(x, y)
}

func main() {
    point := Point{X: 0, Y: 0}
    
    drawShape(point)           // Implements Drawable
    moveShape(&point, 5, 10)   // Implements Movable
    drawShape(point)           // Now at (5.00, 10.00)
}
```

## Empty Interface

### The `interface{}` Type

```go
// Empty interface can hold any type
func processValue(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// Usage
processValue(42)           // Value: 42, Type: int
processValue("hello")      // Value: hello, Type: string
processValue(3.14)         // Value: 3.14, Type: float64
processValue([]int{1, 2})  // Value: [1 2], Type: []int
```

### Type Assertions with Empty Interface

```go
func processValue(value interface{}) {
    // Type assertion
    if str, ok := value.(string); ok {
        fmt.Printf("String: %s (length: %d)\n", str, len(str))
    } else if num, ok := value.(int); ok {
        fmt.Printf("Integer: %d\n", num)
    } else if slice, ok := value.([]int); ok {
        fmt.Printf("Slice: %v (length: %d)\n", slice, len(slice))
    } else {
        fmt.Printf("Unknown type: %T\n", value)
    }
}
```

## Interface Composition

### Embedding Interfaces

```go
// Base interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Composed interfaces
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// Implementation
type File struct {
    name string
}

func (f *File) Read(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

func (f *File) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

func (f *File) Close() error {
    // Implementation
    return nil
}

// File implements ReadWriteCloser automatically
```

### Interface Inheritance

```go
// Base interface
type Animal interface {
    Speak() string
    Move() string
}

// Extended interface
type Pet interface {
    Animal
    Play() string
}

// Implementation
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Running"
}

func (d Dog) Play() string {
    return "Fetching ball"
}

// Dog implements both Animal and Pet
```

## Type Assertions

### Basic Type Assertions

```go
func processValue(value interface{}) {
    // Type assertion
    str := value.(string)
    fmt.Printf("String: %s\n", str)
}

// Safe type assertion
func processValueSafe(value interface{}) {
    if str, ok := value.(string); ok {
        fmt.Printf("String: %s\n", str)
    } else {
        fmt.Printf("Not a string: %T\n", value)
    }
}
```

### Type Assertions with Interfaces

```go
func processShape(value interface{}) {
    if shape, ok := value.(Shape); ok {
        fmt.Printf("Area: %.2f\n", shape.Area())
    } else {
        fmt.Printf("Not a Shape: %T\n", value)
    }
}

// Multiple type assertions
func processValueMultiple(value interface{}) {
    switch v := value.(type) {
    case string:
        fmt.Printf("String: %s\n", v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    case Shape:
        fmt.Printf("Shape area: %.2f\n", v.Area())
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## Type Switches

### Basic Type Switch

```go
func processValue(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    case float64:
        fmt.Printf("Float: %.2f\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

### Type Switch with Interfaces

```go
func processShape(value interface{}) {
    switch s := value.(type) {
    case Rectangle:
        fmt.Printf("Rectangle: %.2f x %.2f\n", s.Width, s.Height)
    case Circle:
        fmt.Printf("Circle: radius %.2f\n", s.Radius)
    case Shape:
        fmt.Printf("Shape area: %.2f\n", s.Area())
    default:
        fmt.Printf("Not a shape: %T\n", s)
    }
}
```

## Interface Best Practices

### 1. Keep Interfaces Small

```go
// Good: Small, focused interface
type Writer interface {
    Write([]byte) (int, error)
}

// Good: Another small interface
type Closer interface {
    Close() error
}

// Good: Compose small interfaces
type WriteCloser interface {
    Writer
    Closer
}

// Avoid: Large interface with many methods
type BadInterface interface {
    Method1()
    Method2()
    Method3()
    Method4()
    Method5()
    Method6()
    Method7()
    Method8()
}
```

### 2. Use Interface Names Ending with 'er'

```go
// Good: Interface names ending with 'er'
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Good: Other common patterns
type Stringer interface {
    String() string
}

type Formatter interface {
    Format(f fmt.State, c rune)
}
```

### 3. Accept Interfaces, Return Concrete Types

```go
// Good: Accept interface, return concrete type
func ProcessData(reader Reader) ([]byte, error) {
    data := make([]byte, 1024)
    n, err := reader.Read(data)
    if err != nil {
        return nil, err
    }
    return data[:n], nil
}

// Avoid: Accepting concrete types
func ProcessDataBad(file *os.File) ([]byte, error) {
    // Limited to File type only
}
```

### 4. Use Interface Composition

```go
// Good: Compose interfaces
type ReadWriter interface {
    Reader
    Writer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// Avoid: Duplicating methods
type BadReadWriter interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}
```

## Common Interface Patterns

### 1. The Stringer Interface

```go
// Stringer interface (from fmt package)
type Stringer interface {
    String() string
}

// Implementation
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Usage
person := Person{Name: "Alice", Age: 30}
fmt.Println(person) // Person{Name: Alice, Age: 30}
```

### 2. The Error Interface

```go
// Error interface (from builtin package)
type error interface {
    Error() string
}

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
        return ValidationError{Field: "name", Message: "name is required"}
    }
    return nil
}
```

### 3. The fmt.Formatter Interface

```go
// Formatter interface (from fmt package)
type Formatter interface {
    Format(f fmt.State, c rune)
}

// Implementation
type CustomType struct {
    Value int
}

func (c CustomType) Format(f fmt.State, c rune) {
    switch c {
    case 'v':
        if f.Flag('#') {
            fmt.Fprintf(f, "CustomType{Value: %d}", c.Value)
        } else {
            fmt.Fprintf(f, "%d", c.Value)
        }
    case 's':
        fmt.Fprintf(f, "CustomType(%d)", c.Value)
    default:
        fmt.Fprintf(f, "%%!%c(CustomType=%d)", c, c.Value)
    }
}
```

### 4. The sort.Interface

```go
// sort.Interface
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// Implementation
type IntSlice []int

func (s IntSlice) Len() int {
    return len(s)
}

func (s IntSlice) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Usage
numbers := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
sort.Sort(numbers)
fmt.Println(numbers) // [1 1 2 3 4 5 6 9]
```

## Advanced Interface Concepts

### 1. Interface Values

```go
// Interface values contain (type, value) pairs
func demonstrateInterfaceValues() {
    var w Writer
    
    // w is nil
    fmt.Printf("w is nil: %t\n", w == nil)
    
    // Assign concrete value
    w = &File{name: "test.txt"}
    
    // w is not nil, but contains (File, &File{name: "test.txt"})
    fmt.Printf("w is nil: %t\n", w == nil)
    fmt.Printf("w type: %T\n", w)
}
```

### 2. Interface Comparison

```go
// Interfaces can be compared if their underlying types are comparable
func compareInterfaces() {
    var w1, w2 Writer
    
    w1 = &File{name: "test.txt"}
    w2 = &File{name: "test.txt"}
    
    // This will be false because pointers are different
    fmt.Printf("w1 == w2: %t\n", w1 == w2)
    
    // But we can compare the underlying values
    f1 := w1.(*File)
    f2 := w2.(*File)
    fmt.Printf("f1.name == f2.name: %t\n", f1.name == f2.name)
}
```

### 3. Interface Embedding

```go
// Interfaces can embed other interfaces
type ReadWriter interface {
    Reader
    Writer
}

// This is equivalent to:
type ReadWriter interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}
```

### 4. Interface Method Sets

```go
// Method sets determine which methods are available
type T struct{}

func (T) method1() {}    // Value receiver
func (*T) method2() {}   // Pointer receiver

func demonstrateMethodSets() {
    var t T
    var p *T = &t
    
    // T has method1
    t.method1()
    
    // *T has both method1 and method2
    p.method1()
    p.method2()
    
    // Interface with value receiver
    var i1 interface{ method1() } = t
    i1.method1()
    
    // Interface with pointer receiver
    var i2 interface{ method2() } = p
    i2.method2()
}
```

## References

- [Go Language Specification - Interface Types](https://golang.org/ref/spec#Interface_types)
- [Effective Go - Interfaces](https://golang.org/doc/effective_go.html#interfaces)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Go by Example - Interface Satisfaction](https://gobyexample.com/interfaces)

## Next Steps

After mastering interfaces, continue with:
- [Concurrency](../concurrency/) - Learn goroutines and channels
- [Error Handling](../error-handling/) - Master Go's error handling patterns
