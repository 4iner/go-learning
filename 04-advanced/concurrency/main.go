package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// This example demonstrates Go's concurrency features
// Run this with: go run main.go

func main() {
	fmt.Println("=== Go Concurrency Examples ===")
	
	// Demonstrate goroutines
	demonstrateGoroutines()
	
	// Demonstrate channels
	demonstrateChannels()
	
	// Demonstrate select statement
	demonstrateSelect()
	
	// Demonstrate synchronization
	demonstrateSynchronization()
	
	// Demonstrate common patterns
	demonstrateCommonPatterns()
	
	// Demonstrate advanced concepts
	demonstrateAdvancedConcepts()
}

// demonstrateGoroutines shows basic goroutine usage
func demonstrateGoroutines() {
	fmt.Println("\n1. Goroutines:")
	
	// Basic goroutine
	fmt.Println("   Basic goroutine:")
	go sayHello("World")
	go sayHello("Go")
	time.Sleep(100 * time.Millisecond)
	
	// Goroutine with WaitGroup
	fmt.Println("\n   Goroutine with WaitGroup:")
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("     Goroutine %d\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("     All goroutines completed")
	
	// Goroutine with channel
	fmt.Println("\n   Goroutine with channel:")
	ch := make(chan string)
	
	go func() {
		ch <- "Hello from goroutine"
	}()
	
	message := <-ch
	fmt.Printf("     %s\n", message)
}

// demonstrateChannels shows channel operations
func demonstrateChannels() {
	fmt.Println("\n2. Channels:")
	
	// Unbuffered channel
	fmt.Println("   Unbuffered channel:")
	ch1 := make(chan int)
	
	go func() {
		ch1 <- 42
		fmt.Println("     Sent 42")
	}()
	
	value := <-ch1
	fmt.Printf("     Received %d\n", value)
	
	// Buffered channel
	fmt.Println("\n   Buffered channel:")
	ch2 := make(chan int, 3)
	
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	
	fmt.Printf("     Received %d\n", <-ch2)
	fmt.Printf("     Received %d\n", <-ch2)
	fmt.Printf("     Received %d\n", <-ch2)
	
	// Channel direction
	fmt.Println("\n   Channel direction:")
	ch3 := make(chan int)
	
	go sendData(ch3)
	go receiveData(ch3)
	
	time.Sleep(100 * time.Millisecond)
	
	// Closed channel
	fmt.Println("\n   Closed channel:")
	ch4 := make(chan int, 3)
	
	ch4 <- 1
	ch4 <- 2
	ch4 <- 3
	close(ch4)
	
	fmt.Println("     Reading from closed channel:")
	for value := range ch4 {
		fmt.Printf("       %d\n", value)
	}
	
	// Check if channel is closed
	value, ok := <-ch4
	if !ok {
		fmt.Println("     Channel is closed")
	}
}

// demonstrateSelect shows select statement usage
func demonstrateSelect() {
	fmt.Println("\n3. Select Statement:")
	
	// Basic select
	fmt.Println("   Basic select:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from ch2"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("     %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("     %s\n", msg2)
		}
	}
	
	// Select with default
	fmt.Println("\n   Select with default:")
	ch3 := make(chan string)
	
	select {
	case msg := <-ch3:
		fmt.Printf("     %s\n", msg)
	default:
		fmt.Println("     No message received")
	}
	
	// Select with timeout
	fmt.Println("\n   Select with timeout:")
	ch4 := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch4 <- "Hello"
	}()
	
	select {
	case msg := <-ch4:
		fmt.Printf("     %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("     Timeout!")
	}
}

// demonstrateSynchronization shows synchronization primitives
func demonstrateSynchronization() {
	fmt.Println("\n4. Synchronization:")
	
	// WaitGroup
	fmt.Println("   WaitGroup:")
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	
	wg.Wait()
	fmt.Println("     All workers completed")
	
	// Mutex
	fmt.Println("\n   Mutex:")
	counter := &Counter{}
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	
	wg.Wait()
	fmt.Printf("     Counter value: %d\n", counter.Value())
	
	// RWMutex
	fmt.Println("\n   RWMutex:")
	safeMap := &SafeMap{data: make(map[string]int)}
	
	// Multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeMap.Get("key")
		}()
	}
	
	// One writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.Set("key", 42)
	}()
	
	wg.Wait()
	fmt.Println("     RWMutex operations completed")
}

// demonstrateCommonPatterns shows common concurrency patterns
func demonstrateCommonPatterns() {
	fmt.Println("\n5. Common Patterns:")
	
	// Worker pool
	fmt.Println("   Worker pool:")
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	for w := 1; w <= 3; w++ {
		go workerPool(w, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
	
	// Pipeline
	fmt.Println("\n   Pipeline:")
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	squares := make(chan int)
	go func() {
		for n := range numbers {
			squares <- n * n
		}
		close(squares)
	}()
	
	fmt.Println("     Pipeline results:")
	for s := range squares {
		fmt.Printf("       %d\n", s)
	}
	
	// Fan-out/Fan-in
	fmt.Println("\n   Fan-out/Fan-in:")
	input := make(chan int)
	
	worker1 := process(input)
	worker2 := process(input)
	worker3 := process(input)
	
	output := merge(worker1, worker2, worker3)
	
	go func() {
		for i := 1; i <= 6; i++ {
			input <- i
		}
		close(input)
	}()
	
	fmt.Println("     Fan-out/Fan-in results:")
	for result := range output {
		fmt.Printf("       %d\n", result)
	}
	
	// Context for cancellation
	fmt.Println("\n   Context for cancellation:")
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("     Operation completed")
	case <-ctx.Done():
		fmt.Println("     Operation cancelled")
	}
}

// demonstrateAdvancedConcepts shows advanced concurrency concepts
func demonstrateAdvancedConcepts() {
	fmt.Println("\n6. Advanced Concepts:")
	
	// Channel of channels
	fmt.Println("   Channel of channels:")
	ch := make(chan chan int)
	
	go func() {
		innerCh := make(chan int)
		ch <- innerCh
		innerCh <- 42
		close(innerCh)
	}()
	
	innerCh := <-ch
	value := <-innerCh
	fmt.Printf("     Received %d from inner channel\n", value)
	
	// Non-blocking operations
	fmt.Println("\n   Non-blocking operations:")
	ch2 := make(chan int)
	
	if nonBlockingSend(ch2, 42) {
		fmt.Println("     Successfully sent 42")
	} else {
		fmt.Println("     Failed to send 42")
	}
	
	// Channel leak prevention
	fmt.Println("\n   Channel leak prevention:")
	ch3 := make(chan int)
	
	go func() {
		defer close(ch3)
		for i := 0; i < 3; i++ {
			ch3 <- i
		}
	}()
	
	fmt.Println("     Reading from channel:")
	for value := range ch3 {
		fmt.Printf("       %d\n", value)
	}
}

// Helper functions
func sayHello(name string) {
	fmt.Printf("     Hello, %s!\n", name)
}

func sendData(ch chan<- int) {
	ch <- 42
}

func receiveData(ch <-chan int) {
	value := <-ch
	fmt.Printf("     Received %d\n", value)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("     Worker %d starting\n", id)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("     Worker %d done\n", id)
}

func workerPool(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("     Worker %d processing job %d\n", id, j)
		time.Sleep(50 * time.Millisecond)
		results <- j * 2
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

func nonBlockingSend(ch chan<- int, value int) bool {
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

// Type definitions
type Counter struct {
	mu    sync.Mutex
	value int
}

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// Method implementations
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
