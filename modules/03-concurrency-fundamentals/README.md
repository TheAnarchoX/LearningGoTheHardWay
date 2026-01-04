# Module 03: Concurrency Fundamentals

## 🎯 Learning Objectives

By completing this module, you will:
- Master goroutines - Go's lightweight threads
- Understand channels for communication between goroutines
- Learn the "share memory by communicating" philosophy
- Avoid common concurrency pitfalls and race conditions
- Build concurrent programs that are safe and efficient

## 📚 Prerequisites

- Completed Modules 01 and 02
- Basic understanding of concurrent programming concepts
- Familiarity with threads/async programming in other languages

## 🗺️ Module Overview

### Coming From Other Languages

**Python Developers:** Goroutines are like async/await but truly parallel (no GIL!). Channels are better than queues.  
**Java Developers:** Goroutines are lighter than threads. Channels are safer than synchronized collections.  
**C++ Developers:** Goroutines are managed for you - no manual thread management. Channels eliminate most mutex needs.  
**JavaScript Developers:** Goroutines run in parallel, not just concurrently. No event loop - true OS threads.

## 📖 Key Concepts

### 1. Goroutines - Cheap Concurrency

```go
// Sequential execution
func doWork() {
    fmt.Println("Working...")
}
doWork()  // Waits for completion

// Concurrent execution
go doWork()  // Doesn't wait, runs in parallel
```

**Key Facts:**
- Goroutines start with ~2KB stack (vs 1MB+ for OS threads)
- Can have millions of goroutines
- Scheduled by Go runtime, not OS
- Always use with synchronization!

### 2. Channels - Communication Between Goroutines

```go
// Create a channel
ch := make(chan int)

// Send to channel (blocks until received)
ch <- 42

// Receive from channel (blocks until sent)
value := <-ch

// Buffered channel (doesn't block until full)
buffered := make(chan int, 10)
```

**Go Proverb:** "Don't communicate by sharing memory; share memory by communicating."

### 3. Channel Patterns

**Fan-Out (one producer, many consumers):**
```go
jobs := make(chan int, 100)
for i := 0; i < 3; i++ {
    go worker(jobs)  // 3 workers
}
```

**Fan-In (many producers, one consumer):**
```go
results := make(chan Result)
for _, task := range tasks {
    go func(t Task) {
        results <- process(t)
    }(task)
}
```

### 4. Select Statement - Multiplexing Channels

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
default:
    fmt.Println("No channels ready")
}
```

### 5. WaitGroup - Waiting for Goroutines

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        doWork(id)
    }(i)
}

wg.Wait()  // Wait for all goroutines
```

### 6. Mutex - Traditional Locking (Use Sparingly!)

```go
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

// Better with defer
func safeIncrement() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

**Prefer channels over mutexes when possible!**

### 7. Context - Cancellation and Timeouts

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case <-ctx.Done():
    fmt.Println("Operation cancelled or timed out")
case result := <-doWork(ctx):
    fmt.Println("Result:", result)
}
```

## 🏋️ Exercises

Work through the exercises in the `exercises/` directory:

1. **exercise1_goroutines.go** - Basic goroutine usage and synchronization
2. **exercise2_channels.go** - Channel patterns and communication
3. **exercise3_select.go** - Select statement and timeouts
4. **exercise4_race_conditions.go** - Find and fix race conditions

Run with race detector: `go test -race`

## 🎓 Common Pitfalls

### 1. Goroutine Leaks
```go
// BAD: Goroutine never exits
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // Blocks forever if nothing sends
        fmt.Println(val)
    }()
    // Channel is never used, goroutine leaks!
}
```

### 2. Loop Variable Capture
```go
// BAD: All goroutines see the same variable
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // Might print 5 five times!
    }()
}

// GOOD: Pass as parameter
for i := 0; i < 5; i++ {
    go func(id int) {
        fmt.Println(id)
    }(i)
}
```

### 3. Closing Channels
```go
// Only sender should close
ch := make(chan int)
go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // Signal no more values
}()

// Receive until closed
for val := range ch {
    fmt.Println(val)
}
```

### 4. Buffered vs Unbuffered
```go
// Unbuffered - synchronous
ch1 := make(chan int)
ch1 <- 1  // Blocks until received

// Buffered - asynchronous until full
ch2 := make(chan int, 1)
ch2 <- 1  // Doesn't block
ch2 <- 2  // Blocks (buffer full)
```

### 5. WaitGroup Gotcha
```go
// BAD: Copying WaitGroup
var wg sync.WaitGroup
func bad(wg sync.WaitGroup) {  // Copied!
    defer wg.Done()
}

// GOOD: Pass pointer
func good(wg *sync.WaitGroup) {
    defer wg.Done()
}
```

## 🔧 Debugging Concurrent Code

### Race Detector
```bash
go test -race
go run -race main.go
```

### Deadlock Detection
Go runtime detects deadlocks:
```
fatal error: all goroutines are asleep - deadlock!
```

## 🏗️ Real-World Patterns

### Worker Pool
```go
func workerPool(jobs <-chan Job, results chan<- Result) {
    for j := range jobs {
        results <- process(j)
    }
}

jobs := make(chan Job, 100)
results := make(chan Result, 100)

// Start workers
for w := 0; w < 3; w++ {
    go workerPool(jobs, results)
}
```

### Pipeline
```go
// Pipeline example with complete usage
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Use: Compose and consume the pipeline
func usePipeline() {
    // Chain stages: generator -> square
    for result := range square(generator(1, 2, 3, 4)) {
        fmt.Println(result)  // Prints: 1, 4, 9, 16
    }
}
```

## 📚 Additional Resources

- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency)
- [Share Memory By Communicating](https://go.dev/blog/codelab-share)

## ✅ Module Checklist

- [ ] Create and manage goroutines
- [ ] Use channels for communication
- [ ] Master the select statement
- [ ] Avoid race conditions
- [ ] Use WaitGroup for synchronization
- [ ] Understand context for cancellation
- [ ] Complete all exercises with -race flag
- [ ] Build a worker pool
- [ ] Build a pipeline

## ⏭️ Next Module

**[Module 04: Error Handling](../04-error-handling/)** - Master Go's explicit error handling patterns!
