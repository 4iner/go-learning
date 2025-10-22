# Go Packages

This comprehensive guide covers Go's package system, which is fundamental to organizing and structuring Go code. Packages provide a way to group related code, control visibility, and create reusable modules.

## Table of Contents
1. [Package Basics](#package-basics)
2. [Package Declaration](#package-declaration)
3. [Import System](#import-system)
4. [Package Visibility](#package-visibility)
5. [Package Initialization](#package-initialization)
6. [Internal Packages](#internal-packages)
7. [Package Design](#package-design)
8. [Best Practices](#best-practices)

## Package Basics

### What is a Package?

A package in Go is a collection of Go source files in the same directory that are compiled together. Packages provide:

- **Code organization**: Group related functionality
- **Visibility control**: Control what is exported/imported
- **Namespace management**: Avoid naming conflicts
- **Reusability**: Share code across projects

### Package Structure

```
myproject/
├── main.go
├── go.mod
├── pkg/
│   ├── math/
│   │   ├── math.go
│   │   └── math_test.go
│   └── utils/
│       ├── string.go
│       └── file.go
└── cmd/
    └── server/
        └── main.go
```

## Package Declaration

### Basic Package Declaration

```go
// math.go
package math

// Exported function (starts with uppercase)
func Add(a, b int) int {
    return a + b
}

// Unexported function (starts with lowercase)
func add(a, b int) int {
    return a + b
}
```

### Package Main

```go
// main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

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

## Import System

### Basic Imports

```go
package main

import "fmt"
import "os"

func main() {
    fmt.Println("Hello")
    os.Exit(0)
}
```

### Multiple Imports

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.ToUpper("hello"))
    os.Exit(0)
}
```

### Import Aliases

```go
package main

import (
    "fmt"
    "math/rand"
    r "math/rand"  // Alias
    _ "database/sql/driver"  // Blank import
)

func main() {
    fmt.Println(rand.Intn(100))
    fmt.Println(r.Intn(100))
}
```

### Local Imports

```go
package main

import (
    "fmt"
    "./math"  // Local package
    "github.com/user/project/pkg/utils"  // External package
)

func main() {
    result := math.Add(2, 3)
    fmt.Println(result)
}
```

## Package Visibility

### Exported vs Unexported

```go
// math.go
package math

// Exported (public) - starts with uppercase
var Pi = 3.14159

func Add(a, b int) int {
    return a + b
}

// Unexported (private) - starts with lowercase
var pi = 3.14159

func add(a, b int) int {
    return a + b
}
```

### Using Exported Symbols

```go
// main.go
package main

import (
    "fmt"
    "yourproject/math"
)

func main() {
    // Can access exported symbols
    result := math.Add(2, 3)
    fmt.Println(result)
    fmt.Println(math.Pi)
    
    // Cannot access unexported symbols
    // result := math.add(2, 3)  // Compile error
    // fmt.Println(math.pi)      // Compile error
}
```

### Struct Field Visibility

```go
// user.go
package user

type User struct {
    Name string    // Exported field
    age  int       // Unexported field
}

func (u *User) GetAge() int {
    return u.age  // Can access unexported field
}

func (u *User) SetAge(age int) {
    u.age = age  // Can modify unexported field
}
```

## Package Initialization

### Package-level Variables

```go
// config.go
package config

import "os"

// Package-level variables
var (
    DatabaseURL string
    Port        int
    Debug       bool
)

// init function runs before main
func init() {
    DatabaseURL = os.Getenv("DATABASE_URL")
    if DatabaseURL == "" {
        DatabaseURL = "localhost:5432"
    }
    
    Port = 8080
    Debug = os.Getenv("DEBUG") == "true"
}
```

### Multiple Init Functions

```go
// math.go
package math

import "fmt"

func init() {
    fmt.Println("math package initialized")
}

// math_advanced.go
package math

import "fmt"

func init() {
    fmt.Println("math advanced package initialized")
}
```

### Init Order

```go
// Package initialization order:
// 1. Package-level variables
// 2. init() functions in order of file appearance
// 3. Imported packages are initialized first
```

## Internal Packages

### Internal Package Concept

```
myproject/
├── pkg/
│   ├── math/
│   │   └── math.go
│   └── internal/
│       ├── cache/
│       │   └── cache.go
│       └── db/
│           └── db.go
└── cmd/
    └── server/
        └── main.go
```

### Internal Package Usage

```go
// pkg/internal/cache/cache.go
package cache

// Internal package - only accessible within pkg/
type Cache struct {
    data map[string]interface{}
}

func New() *Cache {
    return &Cache{
        data: make(map[string]interface{}),
    }
}

// pkg/math/math.go
package math

import "yourproject/pkg/internal/cache"

func ProcessData() {
    c := cache.New()  // Can access internal package
    // ...
}
```

## Package Design

### Single Responsibility

```go
// Good: Single responsibility
// math/arithmetic.go
package math

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

// math/geometry.go
package math

func Area(width, height float64) float64 {
    return width * height
}

func Perimeter(width, height float64) float64 {
    return 2 * (width + height)
}
```

### Package Interfaces

```go
// user/user.go
package user

type User struct {
    ID   int
    Name string
}

// user/service.go
package user

type Service interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

type service struct {
    // implementation details
}

func NewService() Service {
    return &service{}
}
```

### Package Factories

```go
// database/connection.go
package database

import "database/sql"

type Connection struct {
    db *sql.DB
}

func NewConnection(dsn string) (*Connection, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    return &Connection{db: db}, nil
}

func (c *Connection) Close() error {
    return c.db.Close()
}
```

## Best Practices

### 1. Package Naming

```go
// Good: Short, descriptive names
package math
package user
package http

// Avoid: Long, verbose names
package mathematicaloperations
package userdatamanagement
package httpclient
```

### 2. Package Documentation

```go
// Good: Package-level documentation
// Package math provides mathematical operations.
//
// This package includes basic arithmetic operations
// and more advanced mathematical functions.
package math

// Good: Function documentation
// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}
```

### 3. Package Structure

```go
// Good: Logical package structure
// pkg/
//   ├── math/
//   │   ├── arithmetic.go
//   │   ├── geometry.go
//   │   └── math_test.go
//   ├── user/
//   │   ├── user.go
//   │   ├── service.go
//   │   └── user_test.go
//   └── internal/
//       ├── cache/
//       └── db/
```

### 4. Export Only What's Needed

```go
// Good: Export only necessary functions
package math

func Add(a, b int) int {
    return add(a, b)
}

func Subtract(a, b int) int {
    return subtract(a, b)
}

// Unexported helper functions
func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}
```

### 5. Use Package Initialization Sparingly

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
    // Complex business logic
    // Database connections
    // HTTP server startup
}
```

### 6. Package Dependencies

```go
// Good: Minimal dependencies
package math

import "fmt"

func Add(a, b int) int {
    return a + b
}

// Avoid: Heavy dependencies
package bad

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/redis/go-redis/v9"
)
```

## Common Patterns

### 1. Package Factory Pattern

```go
// database/connection.go
package database

type Connection struct {
    // fields
}

func NewConnection(config Config) (*Connection, error) {
    // Create connection
    return &Connection{}, nil
}

func (c *Connection) Close() error {
    // Close connection
    return nil
}
```

### 2. Package Interface Pattern

```go
// storage/storage.go
package storage

type Storage interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}) error
    Delete(key string) error
}

type memoryStorage struct {
    data map[string]interface{}
}

func NewMemoryStorage() Storage {
    return &memoryStorage{
        data: make(map[string]interface{}),
    }
}
```

### 3. Package Configuration Pattern

```go
// config/config.go
package config

type Config struct {
    DatabaseURL string
    Port        int
    Debug       bool
}

func Load() (*Config, error) {
    return &Config{
        DatabaseURL: os.Getenv("DATABASE_URL"),
        Port:        8080,
        Debug:       os.Getenv("DEBUG") == "true",
    }, nil
}
```

## References

- [Go Language Specification - Packages](https://golang.org/ref/spec#Packages)
- [Effective Go - Packages](https://golang.org/doc/effective_go.html#packages)
- [Go by Example - Packages](https://gobyexample.com/packages)
- [Go by Example - Imports](https://gobyexample.com/imports)

## Next Steps

After mastering packages, continue with:
- [Best Practices](../best-practices/) - Learn Go's best practices
- [Advanced Topics](../advanced/) - Explore advanced Go concepts
