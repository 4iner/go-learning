package main

import (
	"fmt"
	"strings"
)

// This example demonstrates Go's data structures
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Data Structures Examples ===")
	
	// Demonstrate arrays
	demonstrateArrays()
	
	// Demonstrate slices
	demonstrateSlices()
	
	// Demonstrate maps
	demonstrateMaps()
	
	// Demonstrate structs
	demonstrateStructs()
	
	// Demonstrate pointers
	demonstratePointers()
	
	// Demonstrate memory management
	demonstrateMemoryManagement()
}

// demonstrateArrays shows array operations
func demonstrateArrays() {
	fmt.Println("\n1. Arrays:")
	
	// Array declaration and initialization
	var arr1 [5]int
	fmt.Printf("   Zero value array: %v\n", arr1)
	
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("   Initialized array: %v\n", arr2)
	
	arr3 := [...]int{1, 2, 3, 4, 5}  // Compiler determines size
	fmt.Printf("   Ellipsis array: %v\n", arr3)
	
	arr4 := [5]int{1: 10, 3: 30}  // Partial initialization
	fmt.Printf("   Partial init array: %v\n", arr4)
	
	// Array operations
	fmt.Printf("   Array length: %d\n", len(arr2))
	fmt.Printf("   First element: %d\n", arr2[0])
	fmt.Printf("   Last element: %d\n", arr2[len(arr2)-1])
	
	// Modifying array
	arr2[0] = 100
	fmt.Printf("   Modified array: %v\n", arr2)
	
	// Iterating over array
	fmt.Print("   Array elements: ")
	for i, v := range arr2 {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()
	
	// Arrays are value types
	arr5 := arr2  // Copy
	arr5[0] = 200
	fmt.Printf("   Original array: %v\n", arr2)
	fmt.Printf("   Copied array: %v\n", arr5)
}

// demonstrateSlices shows slice operations
func demonstrateSlices() {
	fmt.Println("\n2. Slices:")
	
	// Slice declaration and initialization
	var slice1 []int  // nil slice
	fmt.Printf("   Nil slice: %v (len: %d, cap: %d)\n", slice1, len(slice1), cap(slice1))
	
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Initialized slice: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))
	
	slice3 := make([]int, 5)  // Length 5, capacity 5
	fmt.Printf("   Make slice: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	
	slice4 := make([]int, 5, 10)  // Length 5, capacity 10
	fmt.Printf("   Make slice with cap: %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))
	
	// Slice operations
	slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("   Original slice: %v\n", slice5)
	
	// Slicing
	slice6 := slice5[2:5]  // [3 4 5]
	fmt.Printf("   Slice [2:5]: %v\n", slice6)
	
	slice7 := slice5[:5]   // [1 2 3 4 5]
	fmt.Printf("   Slice [:5]: %v\n", slice7)
	
	slice8 := slice5[5:]   // [6 7 8 9 10]
	fmt.Printf("   Slice [5:]: %v\n", slice8)
	
	// Appending
	slice9 := append(slice5, 11, 12, 13)
	fmt.Printf("   Appended slice: %v (len: %d, cap: %d)\n", slice9, len(slice9), cap(slice9))
	
	// Appending slice to slice
	slice10 := []int{14, 15, 16}
	slice11 := append(slice9, slice10...)
	fmt.Printf("   Appended slice to slice: %v\n", slice11)
	
	// Copying slices
	slice12 := make([]int, len(slice11))
	copy(slice12, slice11)
	fmt.Printf("   Copied slice: %v\n", slice12)
	
	// Slice internals demonstration
	fmt.Println("\n   Slice internals:")
	slice13 := []int{1, 2, 3, 4, 5}
	slice14 := slice13[1:4]  // [2 3 4]
	fmt.Printf("   Original: %v\n", slice13)
	fmt.Printf("   Sliced: %v\n", slice14)
	
	slice14[0] = 100  // Modifies original slice
	fmt.Printf("   After modification: %v\n", slice13)
	fmt.Printf("   Sliced after modification: %v\n", slice14)
}

// demonstrateMaps shows map operations
func demonstrateMaps() {
	fmt.Println("\n3. Maps:")
	
	// Map declaration and initialization
	var m1 map[string]int  // nil map
	fmt.Printf("   Nil map: %v\n", m1)
	
	m2 := make(map[string]int)  // Empty map
	fmt.Printf("   Empty map: %v\n", m2)
	
	m3 := map[string]int{
		"apple":  5,
		"banana": 3,
		"cherry": 8,
	}
	fmt.Printf("   Initialized map: %v\n", m3)
	
	// Map operations
	fmt.Printf("   Map length: %d\n", len(m3))
	
	// Adding/updating elements
	m3["orange"] = 4
	m3["apple"] = 6
	fmt.Printf("   After adding/updating: %v\n", m3)
	
	// Accessing elements
	appleCount := m3["apple"]
	fmt.Printf("   Apple count: %d\n", appleCount)
	
	grapeCount := m3["grape"]
	fmt.Printf("   Grape count: %d (zero value)\n", grapeCount)
	
	// Checking if key exists
	if count, exists := m3["grape"]; exists {
		fmt.Printf("   Grape exists: %d\n", count)
	} else {
		fmt.Println("   Grape does not exist")
	}
	
	// Deleting elements
	delete(m3, "banana")
	fmt.Printf("   After deleting banana: %v\n", m3)
	
	// Iterating over map
	fmt.Println("   Map iteration:")
	for key, value := range m3 {
		fmt.Printf("     %s: %d\n", key, value)
	}
	
	// Nested maps
	m4 := map[string]map[string]int{
		"fruits": {
			"apple":  5,
			"banana": 3,
		},
		"vegetables": {
			"carrot": 10,
			"lettuce": 2,
		},
	}
	fmt.Printf("   Nested map: %v\n", m4)
}

// demonstrateStructs shows struct operations
func demonstrateStructs() {
	fmt.Println("\n4. Structs:")
	
	// Basic struct
	type Person struct {
		Name string
		Age  int
	}
	
	// Struct initialization
	var p1 Person  // Zero value
	fmt.Printf("   Zero value struct: %+v\n", p1)
	
	p2 := Person{Name: "Alice", Age: 30}
	fmt.Printf("   Initialized struct: %+v\n", p2)
	
	p3 := Person{"Bob", 25}  // Positional
	fmt.Printf("   Positional struct: %+v\n", p3)
	
	p4 := Person{Name: "Charlie"}  // Partial
	fmt.Printf("   Partial struct: %+v\n", p4)
	
	// Struct with different field types
	type User struct {
		ID       int
		Username string
		Email    string
		Active   bool
		Scores   []int
		Profile  map[string]string
	}
	
	user := User{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		Active:   true,
		Scores:   []int{95, 87, 92},
		Profile:  map[string]string{"city": "New York", "country": "USA"},
	}
	fmt.Printf("   Complex struct: %+v\n", user)
	
	// Anonymous struct
	anon := struct {
		Name string
		Age  int
	}{
		Name: "David",
		Age:  35,
	}
	fmt.Printf("   Anonymous struct: %+v\n", anon)
	
	// Embedded struct
	type Address struct {
		Street string
		City   string
		State  string
		Zip    string
	}
	
	type Employee struct {
		Person  // Embedded
		Address // Embedded
		ID      int
		Salary  float64
	}
	
	emp := Employee{
		Person: Person{Name: "Eve", Age: 28},
		Address: Address{
			Street: "123 Main St",
			City:   "Boston",
			State:  "MA",
			Zip:    "02101",
		},
		ID:     1001,
		Salary: 75000.0,
	}
	fmt.Printf("   Employee struct: %+v\n", emp)
	
	// Struct methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %.2f\n", rect.Area())
	fmt.Printf("   Perimeter: %.2f\n", rect.Perimeter())
	
	rect.Scale(2)
	fmt.Printf("   Scaled rectangle: %+v\n", rect)
}

// demonstratePointers shows pointer operations
func demonstratePointers() {
	fmt.Println("\n5. Pointers:")
	
	// Basic pointers
	x := 42
	p := &x
	fmt.Printf("   x = %d, p = %p, *p = %d\n", x, p, *p)
	
	*p = 100
	fmt.Printf("   After *p = 100: x = %d\n", x)
	
	// Pointer to pointer
	pp := &p
	fmt.Printf("   pp = %p, **pp = %d\n", pp, **pp)
	
	// Pointers with structs
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("   Person: %+v\n", person)
	
	// Value receiver
	name := person.GetName()
	fmt.Printf("   Name: %s\n", name)
	
	// Pointer receiver
	person.SetAge(31)
	fmt.Printf("   After SetAge(31): %+v\n", person)
	
	// Pointers with slices and maps
	slice := []int{1, 2, 3}
	slice2 := slice  // Both point to same underlying array
	slice2[0] = 100
	fmt.Printf("   Original slice: %v\n", slice)
	fmt.Printf("   Modified slice: %v\n", slice2)
	
	m := map[string]int{"a": 1, "b": 2}
	m2 := m  // Both point to same map
	m2["c"] = 3
	fmt.Printf("   Original map: %v\n", m)
	fmt.Printf("   Modified map: %v\n", m2)
	
	// Arrays are value types
	arr := [3]int{1, 2, 3}
	arr2 := arr  // Copy
	arr2[0] = 100
	fmt.Printf("   Original array: %v\n", arr)
	fmt.Printf("   Modified array: %v\n", arr2)
}

// demonstrateMemoryManagement shows memory management concepts
func demonstrateMemoryManagement() {
	fmt.Println("\n6. Memory Management:")
	
	// Pre-allocating slices
	fmt.Println("   Pre-allocating slices:")
	slice := make([]int, 0, 1000)  // Length 0, capacity 1000
	fmt.Printf("   Initial: len=%d, cap=%d\n", len(slice), cap(slice))
	
	for i := 0; i < 100; i++ {
		slice = append(slice, i)
	}
	fmt.Printf("   After 100 appends: len=%d, cap=%d\n", len(slice), cap(slice))
	
	// Slice growth demonstration
	fmt.Println("   Slice growth:")
	slice2 := make([]int, 0, 2)
	fmt.Printf("   Initial: len=%d, cap=%d\n", len(slice2), cap(slice2))
	
	for i := 0; i < 10; i++ {
		slice2 = append(slice2, i)
		fmt.Printf("   After append %d: len=%d, cap=%d\n", i, len(slice2), cap(slice2))
	}
	
	// Memory-efficient string building
	fmt.Println("   String building:")
	var builder strings.Builder
	builder.Grow(100)  // Pre-allocate capacity
	for i := 0; i < 10; i++ {
		builder.WriteString(fmt.Sprintf("item%d ", i))
	}
	result := builder.String()
	fmt.Printf("   Built string: %s\n", result)
}

// Rectangle struct for demonstration
type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Pointer receiver method
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Method with multiple receivers
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Person struct for demonstration
type Person struct {
	Name string
	Age  int
}

// Value receiver method
func (p Person) GetName() string {
	return p.Name
}

// Pointer receiver method
func (p *Person) SetAge(age int) {
	p.Age = age
}
