# Go Concurrency

This comprehensive guide covers Go's concurrency model, which is one of the language's most powerful features. Go provides goroutines for concurrent execution and channels for communication between goroutines, following the principle "Don't communicate by sharing memory; share memory by communicating."

## Table of Contents
1. [Goroutines](#goroutines)
2. [Channels](#channels)
3. [Channel Types](#channel-types)
4. [Select Statement](#select-statement)
5. [Synchronization](#synchronization)
6. [Common Patterns](#common-patterns)
7. [Best Practices](#best-practices)
8. [Advanced Concepts](#advanced-concepts)

## Goroutines

### Basic Goroutines

```go
// Basic goroutine
func main() {
    go sayHello("World")
    go sayHello("Go")
    
    // Wait for goroutines to complete
    time.Sleep(1 * time.Second)
}

func sayHello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

### Goroutine with WaitGroup

```go
import "sync"

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All goroutines completed")
}
```

### Goroutine with Channels

```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello from goroutine"
    }()
    
    message := <-ch
    fmt.Println(message)
}
```

## Channels

### Basic Channel Operations

```go
// Channel declaration and creation
var ch chan int           // nil channel
ch = make(chan int)       // unbuffered channel
ch = make(chan int, 5)    // buffered channel with capacity 5

// Channel operations
ch <- 42        // Send value to channel
value := <-ch   // Receive value from channel
close(ch)       // Close channel
```

### Unbuffered Channels

```go
func main() {
    ch := make(chan int)  // Unbuffered channel
    
    go func() {
        ch <- 42
        fmt.Println("Sent 42")
    }()
    
    value := <-ch
    fmt.Printf("Received %d\n", value)
}
```

### Buffered Channels

```go
func main() {
    ch := make(chan int, 3)  // Buffered channel
    
    // Send values without blocking
    ch <- 1
    ch <- 2
    ch <- 3
    
    // Receive values
    fmt.Println(<-ch)  // 1
    fmt.Println(<-ch)  // 2
    fmt.Println(<-ch)  // 3
}
```

### Channel Direction

```go
// Send-only channel
func sendData(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiveData(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
}

// Bidirectional channel
func processData(ch chan int) {
    ch <- 42
    value := <-ch
    fmt.Println(value)
}
```

## Channel Types

### Unbuffered Channels

```go
func main() {
    ch := make(chan string)
    
    go func() {
        ch <- "Hello"
        ch <- "World"
    }()
    
    fmt.Println(<-ch)  // Hello
    fmt.Println(<-ch)  // World
}
```

### Buffered Channels

```go
func main() {
    ch := make(chan int, 2)
    
    ch <- 1
    ch <- 2
    
    fmt.Println(<-ch)  // 1
    fmt.Println(<-ch)  // 2
}
```

### Closed Channels

```go
func main() {
    ch := make(chan int, 3)
    
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    
    // Reading from closed channel
    for value := range ch {
        fmt.Println(value)
    }
    
    // Check if channel is closed
    value, ok := <-ch
    if !ok {
        fmt.Println("Channel is closed")
    }
}
```

## Select Statement

### Basic Select

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}
```

### Select with Default

```go
func main() {
    ch := make(chan string)
    
    select {
    case msg := <-ch:
        fmt.Println(msg)
    default:
        fmt.Println("No message received")
    }
}
```

### Select with Timeout

```go
func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Hello"
    }()
    
    select {
    case msg := <-ch:
        fmt.Println(msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

## Synchronization

### WaitGroup

```go
import "sync"

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    wg.Wait()
    fmt.Println("All workers completed")
}

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}
```

### Mutex

```go
import "sync"

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &Counter{}
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter value: %d\n", counter.Value())
}
```

### RWMutex

```go
import "sync"

type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    value, ok := sm.data[key]
    return value, ok
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}

func main() {
    sm := &SafeMap{data: make(map[string]int)}
    
    // Multiple readers
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            sm.Get("key")
        }()
    }
    
    wg.Wait()
}
```

## Common Patterns

### Worker Pool

```go
func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for a := 1; a <= 9; a++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}
```

### Pipeline

```go
func main() {
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        for i := 1; i <= 5; i++ {
            numbers <- i
        }
        close(numbers)
    }()
    
    // Stage 2: Square numbers
    squares := make(chan int)
    go func() {
        for n := range numbers {
            squares <- n * n
        }
        close(squares)
    }()
    
    // Stage 3: Print results
    for s := range squares {
        fmt.Println(s)
    }
}
```

### Fan-out/Fan-in

```go
func main() {
    input := make(chan int)
    
    // Fan-out: Multiple workers
    worker1 := process(input)
    worker2 := process(input)
    worker3 := process(input)
    
    // Fan-in: Merge results
    output := merge(worker1, worker2, worker3)
    
    // Send input
    go func() {
        for i := 1; i <= 10; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Collect output
    for result := range output {
        fmt.Println(result)
    }
}

func process(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for n := range input {
            output <- n * n
        }
    }()
    return output
}

func merge(channels ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch {
                output <- n
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}
```

### Context for Cancellation

```go
import "context"

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    go func() {
        time.Sleep(2 * time.Second)
        cancel()
    }()
    
    select {
    case <-time.After(5 * time.Second):
        fmt.Println("Operation completed")
    case <-ctx.Done():
        fmt.Println("Operation cancelled")
    }
}
```

## Best Practices

### 1. Use Channels for Communication

```go
// Good: Use channels for communication
func processData(input <-chan int, output chan<- int) {
    for data := range input {
        result := data * 2
        output <- result
    }
}

// Avoid: Shared memory with mutexes
type SharedData struct {
    mu   sync.Mutex
    data int
}

func (sd *SharedData) Process() {
    sd.mu.Lock()
    defer sd.mu.Unlock()
    sd.data *= 2
}
```

### 2. Close Channels Properly

```go
// Good: Close channels when done
func producer(ch chan<- int) {
    defer close(ch)
    for i := 0; i < 5; i++ {
        ch <- i
    }
}

// Good: Check if channel is closed
func consumer(ch <-chan int) {
    for {
        value, ok := <-ch
        if !ok {
            break
        }
        fmt.Println(value)
    }
}
```

### 3. Use Select for Non-blocking Operations

```go
// Good: Use select for non-blocking operations
func nonBlockingSend(ch chan<- int, value int) bool {
    select {
    case ch <- value:
        return true
    default:
        return false
    }
}

// Good: Use select with timeout
func operationWithTimeout(ch <-chan int) {
    select {
    case value := <-ch:
        fmt.Println(value)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout")
    }
}
```

### 4. Use WaitGroup for Synchronization

```go
// Good: Use WaitGroup to wait for goroutines
func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Do work
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All goroutines completed")
}

// Avoid: Using time.Sleep
func main() {
    for i := 0; i < 5; i++ {
        go func(id int) {
            // Do work
        }(i)
    }
    
    time.Sleep(1 * time.Second)  // Unreliable
    fmt.Println("All goroutines completed")
}
```

### 5. Use Context for Cancellation

```go
// Good: Use context for cancellation
func longRunningOperation(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
            time.Sleep(100 * time.Millisecond)
        }
    }
}

// Avoid: Manual cancellation
func longRunningOperationBad(cancel chan bool) error {
    for {
        select {
        case <-cancel:
            return errors.New("cancelled")
        default:
            // Do work
            time.Sleep(100 * time.Millisecond)
        }
    }
}
```

## Advanced Concepts

### 1. Channel of Channels

```go
func main() {
    ch := make(chan chan int)
    
    go func() {
        innerCh := make(chan int)
        ch <- innerCh
        innerCh <- 42
        close(innerCh)
    }()
    
    innerCh := <-ch
    value := <-innerCh
    fmt.Println(value)  // 42
}
```

### 2. Channel Direction Constraints

```go
// Send-only channel
func sendOnly(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiveOnly(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
}

// Bidirectional channel
func bidirectional(ch chan int) {
    ch <- 42
    value := <-ch
    fmt.Println(value)
}
```

### 3. Channel Buffering Strategies

```go
// Unbuffered: Synchronous communication
ch1 := make(chan int)

// Buffered: Asynchronous communication
ch2 := make(chan int, 10)

// Zero-capacity: Synchronous
ch3 := make(chan int, 0)
```

### 4. Channel Leak Prevention

```go
func main() {
    ch := make(chan int)
    
    go func() {
        defer close(ch)  // Always close channels
        for i := 0; i < 5; i++ {
            ch <- i
        }
    }()
    
    for value := range ch {
        fmt.Println(value)
    }
}
```

## References

- [Go Language Specification - Goroutines](https://golang.org/ref/spec#Go_statements)
- [Go Language Specification - Channels](https://golang.org/ref/spec#Channel_types)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go.html#concurrency)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Go by Example - Select](https://gobyexample.com/select)

## Next Steps

After mastering concurrency, continue with:
- [Error Handling](../error-handling/) - Learn Go's error handling patterns
- [Testing](../testing/) - Master Go's testing framework
