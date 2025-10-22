# Go Modules System

Go modules are the standard way to manage dependencies in Go projects. This comprehensive guide covers everything you need to know about Go modules, from basic concepts to advanced usage patterns.

## Table of Contents
1. [What are Go Modules?](#what-are-go-modules)
2. [Module Structure](#module-structure)
3. [go.mod File](#gomod-file)
4. [go.sum File](#gosum-file)
5. [Module Commands](#module-commands)
6. [Dependency Management](#dependency-management)
7. [Version Selection](#version-selection)
8. [Replace Directives](#replace-directives)
9. [Best Practices](#best-practices)
10. [Common Patterns](#common-patterns)

## What are Go Modules?

Go modules are collections of Go packages that are versioned together. They solve the dependency management problem by:

- **Providing reproducible builds**: Exact versions of dependencies are recorded
- **Enabling semantic versioning**: Clear versioning scheme for dependencies
- **Supporting proxy servers**: Fast, reliable dependency downloads
- **Eliminating GOPATH**: No need for GOPATH environment variable

### Key Concepts

- **Module**: A collection of Go packages with a `go.mod` file at the root
- **Package**: A directory containing Go source files
- **Import Path**: The path used to import a package (e.g., `github.com/user/repo/pkg`)
- **Module Path**: The module's name, usually the repository URL

## Module Structure

```
my-project/
├── go.mod          # Module definition and dependencies
├── go.sum          # Checksums for dependencies
├── main.go         # Main package
├── pkg/            # Internal packages
│   ├── utils/
│   │   └── helper.go
│   └── models/
│       └── user.go
└── cmd/            # Command-line applications
    └── server/
        └── main.go
```

## go.mod File

The `go.mod` file defines the module and its dependencies:

```go
module github.com/username/project-name

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
)

require (
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
)
```

### go.mod Directives

- **module**: Defines the module path
- **go**: Minimum Go version required
- **require**: Direct dependencies
- **exclude**: Exclude specific versions
- **replace**: Replace dependencies with local or different versions
- **retract**: Retract published versions

## go.sum File

The `go.sum` file contains cryptographic checksums of dependencies:

```
github.com/bytedance/sonic v1.9.1 h1:6iJ6NqdoxCDr6mbY8x18oSO+c1FqMYriHtv70tM9v=
github.com/bytedance/sonic v1.9.1/go.mod h1:7204Tv+7Lg3gJeqg4AWj8v3ZIMf3N+Jf+6iWxHf0=
```

**Important**: Always commit `go.sum` to version control!

## Module Commands

### Initialize a Module
```bash
# Initialize a new module
go mod init github.com/username/project-name

# Initialize in existing directory
go mod init .
```

### Add Dependencies
```bash
# Add a dependency
go get github.com/gin-gonic/gin

# Add specific version
go get github.com/gin-gonic/gin@v1.9.1

# Add latest version
go get github.com/gin-gonic/gin@latest
```

### Remove Dependencies
```bash
# Remove unused dependencies
go mod tidy

# Remove specific dependency
go get github.com/gin-gonic/gin@none
```

### Update Dependencies
```bash
# Update all dependencies
go get -u

# Update specific dependency
go get -u github.com/gin-gonic/gin

# Update to latest minor/patch versions
go get -u=patch
```

## Dependency Management

### Semantic Versioning

Go modules use semantic versioning (semver):

- **Major version (v2+)**: Breaking changes
- **Minor version (v1.2)**: New features, backward compatible
- **Patch version (v1.2.3)**: Bug fixes, backward compatible

### Version Selection

Go automatically selects the highest compatible version:

```go
// go.mod
require github.com/example/package v1.2.0

// Go will use v1.2.5 if available (highest v1.2.x)
// But won't use v2.0.0 (different major version)
```

### Major Version Upgrades

For major version upgrades, import paths change:

```go
// v1
import "github.com/example/package"

// v2
import "github.com/example/package/v2"
```

## Replace Directives

Replace directives allow you to substitute dependencies:

```go
// Replace with local path
replace github.com/example/package => ./local/package

// Replace with different version
replace github.com/example/package => github.com/example/package v1.2.3

// Replace with fork
replace github.com/example/package => github.com/myfork/package v1.2.3
```

## Best Practices

### 1. Module Naming
```go
// Good: Use repository URL
module github.com/username/project-name

// Avoid: Generic names
module my-project
```

### 2. Version Management
```bash
# Always use specific versions in production
go get github.com/gin-gonic/gin@v1.9.1

# Use latest for development
go get github.com/gin-gonic/gin@latest
```

### 3. Dependency Organization
```go
// Group related dependencies
require (
    // Web framework
    github.com/gin-gonic/gin v1.9.1
    
    // Database
    github.com/lib/pq v1.10.9
    
    // Testing
    github.com/stretchr/testify v1.8.4
)
```

### 4. Private Modules
```bash
# Set GOPRIVATE for private repositories
go env -w GOPRIVATE=github.com/yourcompany/*

# Configure authentication
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

## Common Patterns

### 1. Multi-Module Workspace
```bash
# Create workspace
go work init

# Add modules to workspace
go work use ./module1
go work use ./module2
```

### 2. Module Proxy
```bash
# Use module proxy
go env -w GOPROXY=https://proxy.golang.org,direct

# Use private proxy
go env -w GOPROXY=https://proxy.company.com,https://proxy.golang.org,direct
```

### 3. Vendor Dependencies
```bash
# Create vendor directory
go mod vendor

# Use vendored dependencies
go build -mod=vendor
```

## References

- [Go Modules Documentation](https://golang.org/ref/mod)
- [Go Modules Wiki](https://github.com/golang/go/wiki/Modules)
- [Semantic Versioning](https://semver.org/)
- [Go Module Proxy](https://proxy.golang.org/)

## Next Steps

After understanding Go modules, proceed to:
- [Basic Syntax](../basic-syntax/) - Learn Go's fundamental syntax
- [Control Flow](../control-flow/) - Master Go's control structures
