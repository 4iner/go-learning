# Go Data Structures

This comprehensive guide covers Go's fundamental data structures including arrays, slices, maps, and structs. Understanding these data structures is essential for writing efficient and idiomatic Go programs.

## Table of Contents
1. [Arrays](#arrays)
2. [Slices](#slices)
3. [Maps](#maps)
4. [Structs](#structs)
5. [Pointers](#pointers)
6. [Memory Management](#memory-management)
7. [Best Practices](#best-practices)

## Arrays

### Basic Array Declaration

```go
// Array declaration with size and type
var arr [5]int                    // [0 0 0 0 0]
var arr2 [3]string                // ["" "" ""]

// Array initialization
var arr3 [5]int = [5]int{1, 2, 3, 4, 5}
var arr4 = [5]int{1, 2, 3, 4, 5}  // Type inference
arr5 := [5]int{1, 2, 3, 4, 5}     // Short declaration

// Array with ellipsis (compiler determines size)
arr6 := [...]int{1, 2, 3, 4, 5}   // Size is 5

// Partial initialization
arr7 := [5]int{1, 2}              // [1 2 0 0 0]
arr8 := [5]int{1: 10, 3: 30}      // [0 10 0 30 0]
```

### Array Operations

```go
// Accessing elements
arr := [5]int{1, 2, 3, 4, 5}
first := arr[0]    // 1
last := arr[4]      // 5

// Modifying elements
arr[0] = 10         // [10 2 3 4 5]

// Array length
length := len(arr)  // 5

// Iterating over array
for i := 0; i < len(arr); i++ {
    fmt.Printf("arr[%d] = %d\n", i, arr[i])
}

// Range over array
for index, value := range arr {
    fmt.Printf("arr[%d] = %d\n", index, value)
}
```

### Array Characteristics

```go
// Arrays are value types (copied by value)
arr1 := [3]int{1, 2, 3}
arr2 := arr1        // Copy of arr1
arr2[0] = 10        // arr1 unchanged
fmt.Println(arr1)   // [1 2 3]
fmt.Println(arr2)   // [10 2 3]

// Array size is part of the type
var arr3 [3]int
var arr4 [5]int
// arr3 = arr4  // Compile error: cannot assign [5]int to [3]int
```

## Slices

### Slice Declaration and Initialization

```go
// Slice declaration
var slice []int                    // nil slice
var slice2 []string               // nil slice

// Slice initialization
slice3 := []int{1, 2, 3, 4, 5}    // [1 2 3 4 5]
slice4 := []string{"a", "b", "c"}  // ["a" "b" "c"]

// Make slice with length
slice5 := make([]int, 5)           // [0 0 0 0 0]
slice6 := make([]int, 5, 10)       // Length 5, capacity 10

// Slice from array
arr := [5]int{1, 2, 3, 4, 5}
slice7 := arr[1:4]                 // [2 3 4]
slice8 := arr[:3]                  // [1 2 3]
slice9 := arr[2:]                  // [3 4 5]
slice10 := arr[:]                  // [1 2 3 4 5]
```

### Slice Operations

```go
// Basic operations
slice := []int{1, 2, 3, 4, 5}

// Length and capacity
fmt.Println(len(slice))  // 5
fmt.Println(cap(slice))  // 5

// Accessing elements
first := slice[0]        // 1
last := slice[len(slice)-1] // 5

// Modifying elements
slice[0] = 10            // [10 2 3 4 5]

// Appending elements
slice = append(slice, 6) // [10 2 3 4 5 6]
slice = append(slice, 7, 8, 9) // [10 2 3 4 5 6 7 8 9]

// Appending slice to slice
slice2 := []int{10, 11, 12}
slice = append(slice, slice2...) // [10 2 3 4 5 6 7 8 9 10 11 12]
```

### Slice Slicing

```go
slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Basic slicing
slice1 := slice[2:5]     // [3 4 5]
slice2 := slice[:5]      // [1 2 3 4 5]
slice3 := slice[5:]      // [6 7 8 9 10]
slice4 := slice[:]       // [1 2 3 4 5 6 7 8 9 10]

// Slicing with capacity
slice5 := slice[2:5:7]   // [3 4 5], capacity 5

// Copying slices
slice6 := make([]int, len(slice))
copy(slice6, slice)      // Copy slice to slice6
```

### Slice Internals

```go
// Slice header structure
type slice struct {
    data uintptr  // Pointer to underlying array
    len  int      // Length
    cap  int      // Capacity
}

// Understanding slice behavior
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1[1:4]    // [2 3 4]

// Both slices share the same underlying array
slice2[0] = 10           // Modifies slice1 too!
fmt.Println(slice1)      // [1 10 3 4 5]
fmt.Println(slice2)       // [10 3 4]

// Creating independent slice
slice3 := make([]int, len(slice1))
copy(slice3, slice1)      // Independent copy
slice3[0] = 100           // Doesn't affect slice1
```

## Maps

### Map Declaration and Initialization

```go
// Map declaration
var m map[string]int     // nil map
var m2 map[int]string    // nil map

// Map initialization
m3 := make(map[string]int)           // Empty map
m4 := map[string]int{                // Map literal
    "apple":  5,
    "banana": 3,
    "cherry": 8,
}

// Map with different types
m5 := map[int]string{
    1: "one",
    2: "two",
    3: "three",
}

// Nested maps
m6 := map[string]map[string]int{
    "fruits": {
        "apple":  5,
        "banana": 3,
    },
    "vegetables": {
        "carrot": 10,
        "lettuce": 2,
    },
}
```

### Map Operations

```go
// Basic operations
m := map[string]int{
    "apple":  5,
    "banana": 3,
    "cherry": 8,
}

// Adding/updating elements
m["orange"] = 4        // Add new element
m["apple"] = 6         // Update existing element

// Accessing elements
appleCount := m["apple"]    // 6
grapeCount := m["grape"]    // 0 (zero value)

// Checking if key exists
if count, exists := m["grape"]; exists {
    fmt.Printf("Grape count: %d\n", count)
} else {
    fmt.Println("Grape not found")
}

// Deleting elements
delete(m, "banana")    // Remove "banana" key

// Map length
fmt.Println(len(m))    // 2
```

### Map Iteration

```go
m := map[string]int{
    "apple":  5,
    "banana": 3,
    "cherry": 8,
}

// Iterate over map
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// Iterate over keys only
for key := range m {
    fmt.Printf("Key: %s\n", key)
}

// Iterate over values only
for _, value := range m {
    fmt.Printf("Value: %d\n", value)
}
```

### Map Characteristics

```go
// Maps are reference types
m1 := map[string]int{"a": 1, "b": 2}
m2 := m1              // m2 points to same map as m1
m2["c"] = 3           // Modifies m1 too!
fmt.Println(m1)       // map[a:1 b:2 c:3]
fmt.Println(m2)       // map[a:1 b:2 c:3]

// Maps are not safe for concurrent access
// Use sync.Map or mutex for concurrent access
```

## Structs

### Struct Declaration

```go
// Basic struct
type Person struct {
    Name string
    Age  int
}

// Struct with different field types
type User struct {
    ID       int
    Username string
    Email    string
    Active   bool
    Scores   []int
    Profile  map[string]string
}

// Anonymous struct
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  30,
}

// Embedded struct
type Address struct {
    Street string
    City   string
    State  string
    Zip    string
}

type Employee struct {
    Person  // Embedded struct
    Address // Embedded struct
    ID      int
    Salary  float64
}
```

### Struct Initialization

```go
// Zero value initialization
var p Person  // {"" 0}

// Field initialization
p1 := Person{
    Name: "Alice",
    Age:  30,
}

// Positional initialization
p2 := Person{"Bob", 25}

// Partial initialization
p3 := Person{Name: "Charlie"}  // {Charlie 0}

// Pointer to struct
p4 := &Person{
    Name: "David",
    Age:  35,
}

// Using new
p5 := new(Person)  // Returns pointer to zero value
```

### Struct Methods

```go
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

// Usage
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()        // 50
rect.Scale(2)              // Width: 20, Height: 10
perimeter := rect.Perimeter() // 60
```

### Struct Tags

```go
import "encoding/json"

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`  // Ignore in JSON
    Age      int    `json:"age,omitempty"`  // Omit if zero
}

// JSON marshaling
user := User{
    ID:       1,
    Username: "alice",
    Email:    "alice@example.com",
    Password: "secret",
    Age:      30,
}

jsonData, err := json.Marshal(user)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))
// {"id":1,"username":"alice","email":"alice@example.com","age":30}
```

## Pointers

### Basic Pointers

```go
// Pointer declaration
var p *int        // nil pointer
var q *string     // nil pointer

// Getting address of variable
x := 42
p = &x            // p points to x

// Dereferencing pointer
fmt.Println(*p)   // 42
*p = 100          // Changes x to 100
fmt.Println(x)    // 100

// Pointer to pointer
pp := &p          // Pointer to pointer
fmt.Println(**pp) // 100
```

### Pointers with Structs

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

// Usage
person := Person{Name: "Alice", Age: 30}
name := person.GetName()    // "Alice"
person.SetAge(31)           // Modifies person.Age
fmt.Println(person.Age)     // 31
```

### Pointers with Slices and Maps

```go
// Slices are already reference types
slice := []int{1, 2, 3}
slice2 := slice              // Both point to same underlying array
slice2[0] = 10               // Modifies slice too
fmt.Println(slice)           // [10 2 3]

// Maps are already reference types
m := map[string]int{"a": 1}
m2 := m                      // Both point to same map
m2["b"] = 2                  // Modifies m too
fmt.Println(m)               // map[a:1 b:2]

// Arrays are value types
arr := [3]int{1, 2, 3}
arr2 := arr                  // Copy of arr
arr2[0] = 10                 // Doesn't affect arr
fmt.Println(arr)             // [1 2 3]
fmt.Println(arr2)            // [10 2 3]
```

## Memory Management

### Garbage Collection

```go
// Go has automatic garbage collection
// No need to manually free memory

func createSlice() []int {
    slice := make([]int, 1000)
    // ... use slice
    return slice  // slice will be garbage collected when no longer referenced
}

// Memory is automatically reclaimed
```

### Memory Optimization

```go
// Pre-allocate slices when size is known
slice := make([]int, 0, 1000)  // Length 0, capacity 1000
for i := 0; i < 1000; i++ {
    slice = append(slice, i)    // No reallocation needed
}

// Use sync.Pool for frequently allocated objects
import "sync"

var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

// Get from pool
buf := pool.Get().([]byte)
// Use buf
pool.Put(buf)  // Return to pool
```

## Best Practices

### 1. Use Slices Instead of Arrays

```go
// Good: Use slices for dynamic data
func processItems(items []int) {
    for _, item := range items {
        // process item
    }
}

// Avoid: Arrays for dynamic data
func processItems(items [100]int) {
    // Limited to exactly 100 items
}
```

### 2. Pre-allocate Slices

```go
// Good: Pre-allocate when size is known
func createSlice(size int) []int {
    slice := make([]int, 0, size)
    for i := 0; i < size; i++ {
        slice = append(slice, i)
    }
    return slice
}

// Avoid: Growing slice multiple times
func createSliceBad(size int) []int {
    var slice []int
    for i := 0; i < size; i++ {
        slice = append(slice, i)  // May cause multiple reallocations
    }
    return slice
}
```

### 3. Use Map for Lookups

```go
// Good: Use map for O(1) lookups
func findUser(users map[string]User, username string) (User, bool) {
    user, exists := users[username]
    return user, exists
}

// Avoid: Linear search through slice
func findUserBad(users []User, username string) (User, bool) {
    for _, user := range users {
        if user.Username == username {
            return user, true
        }
    }
    return User{}, false
}
```

### 4. Use Struct Tags

```go
// Good: Use struct tags for serialization
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// Avoid: Manual field mapping
func (u User) ToJSON() string {
    return fmt.Sprintf(`{"id":%d,"username":"%s","email":"%s"}`, 
        u.ID, u.Username, u.Email)
}
```

### 5. Use Pointer Receivers for Methods

```go
// Good: Use pointer receivers for methods that modify the struct
func (p *Person) SetAge(age int) {
    p.Age = age
}

// Good: Use value receivers for methods that don't modify
func (p Person) GetAge() int {
    return p.Age
}
```

## References

- [Go Language Specification - Types](https://golang.org/ref/spec#Types)
- [Effective Go - Slices](https://golang.org/doc/effective_go.html#slices)
- [Go by Example - Arrays](https://gobyexample.com/arrays)
- [Go by Example - Slices](https://gobyexample.com/slices)
- [Go by Example - Maps](https://gobyexample.com/maps)
- [Go by Example - Structs](https://gobyexample.com/structs)

## Next Steps

After mastering data structures, continue with:
- [Functions](../functions/) - Learn function declarations and usage
- [Interfaces](../interfaces/) - Master Go's interface system
