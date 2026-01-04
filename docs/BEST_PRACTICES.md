# Go Best Practices for Senior Developers

This guide compiles best practices specifically for experienced developers transitioning to Go.

## Code Organization

### Package Structure
```
myproject/
├── cmd/              # Main applications
│   └── myapp/
│       └── main.go
├── internal/         # Private packages
│   ├── service/
│   └── repository/
├── pkg/             # Public packages (use sparingly)
│   └── api/
├── api/             # API definitions (protobuf, OpenAPI)
├── web/             # Web assets
├── scripts/         # Build and deployment scripts
├── test/            # Additional test files
└── go.mod
```

### Naming Conventions

**Packages:**
- Short, concise, lowercase, no underscores
- `httputil`, not `http_util`
- Name by what it provides, not what it contains
- `http`, not `httpserver`

**Interfaces:**
- One-method interfaces named with `-er` suffix
- `Reader`, `Writer`, `Formatter`
- Defined where they're used, not where they're implemented

**Variables:**
- Use short names in small scopes: `i`, `r`, `w`
- Use longer names for package-level vars: `defaultTimeout`
- Acronyms: `userID`, `httpServer` (not `userId`, `HTTPServer`)

## Idiomatic Go

### Error Handling
```go
// ✅ GOOD: Check errors immediately
result, err := doSomething()
if err != nil {
    return fmt.Errorf("do something: %w", err)
}

// ❌ BAD: Don't ignore errors
result, _ := doSomething()

// ❌ BAD: Don't use panic for normal errors
if err != nil {
    panic(err)  // Only for truly exceptional situations
}
```

### Return Early
```go
// ✅ GOOD: Return early, reduce nesting
func process(data []byte) error {
    if len(data) == 0 {
        return ErrEmptyData
    }
    
    if err := validate(data); err != nil {
        return err
    }
    
    return save(data)
}

// ❌ BAD: Deep nesting
func process(data []byte) error {
    if len(data) > 0 {
        if err := validate(data); err == nil {
            return save(data)
        } else {
            return err
        }
    }
    return ErrEmptyData
}
```

### Use defer for Cleanup
```go
// ✅ GOOD: defer ensures cleanup
func process(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    
    // Process file...
    return nil
}
```

### Prefer Composition
```go
// ✅ GOOD: Composition
type Logger struct {
    writer io.Writer
}

type Service struct {
    logger Logger
    db     Database
}

// ❌ Avoid: Deep inheritance hierarchies don't exist in Go
```

## Interface Design

### Accept Interfaces, Return Structs
```go
// ✅ GOOD: Accept interface
func Save(w io.Writer, data []byte) error {
    _, err := w.Write(data)
    return err
}

// ✅ GOOD: Return concrete type
func NewConfig() *Config {
    return &Config{}
}

// ❌ Avoid: Returning interfaces (usually)
func NewService() Service {  // Usually not needed
    return &service{}
}
```

### Small Interfaces
```go
// ✅ GOOD: Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

// ❌ Avoid: Large interfaces
type Everything interface {
    Read() error
    Write() error
    Close() error
    Open() error
    // ... 10 more methods
}
```

## Concurrency

### Use Channels to Communicate
```go
// ✅ GOOD: Channel for coordination
func worker(jobs <-chan Job, results chan<- Result) {
    for j := range jobs {
        results <- process(j)
    }
}

// ⚠️ OK but less idiomatic: Mutex
type Counter struct {
    mu    sync.Mutex
    value int
}
```

### Don't Start Goroutines in Libraries
```go
// ❌ BAD: Library controls goroutines
func ProcessAsync(data []byte) {
    go process(data)  // Caller loses control
}

// ✅ GOOD: Let caller control concurrency
func Process(data []byte) error {
    return process(data)
}

// Caller decides:
go Process(data)
```

### Always Handle Context
```go
// ✅ GOOD: Respect context cancellation
func doWork(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work...
        }
    }
}
```

## Performance

### Preallocate Slices
```go
// ✅ GOOD: Preallocate if size is known
items := make([]Item, 0, expectedSize)

// ⚠️ OK but slower: Let it grow
items := []Item{}
```

### Use Pointers for Large Structs
```go
type LargeStruct struct {
    data [10000]byte
}

// ✅ GOOD: Pass pointer
func process(s *LargeStruct) {}

// ❌ BAD: Pass by value (copies 10KB)
func process(s LargeStruct) {}
```

### Avoid String Concatenation in Loops
```go
// ❌ BAD: Creates many intermediate strings
var result string
for _, s := range strings {
    result += s
}

// ✅ GOOD: Use strings.Builder
var builder strings.Builder
for _, s := range strings {
    builder.WriteString(s)
}
result := builder.String()
```

## Testing

### Table-Driven Tests
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

### Use Interfaces for Mocking
```go
// ✅ GOOD: Interface enables testing
type Storage interface {
    Save(data []byte) error
}

type Service struct {
    storage Storage
}

// Easy to mock in tests
type MockStorage struct{}
func (m *MockStorage) Save(data []byte) error {
    return nil
}
```

## Common Mistakes from Other Languages

### 1. Trying to Use Exceptions
```go
// ❌ Python/Java habit: Using panic like exceptions
if err != nil {
    panic(err)  // Don't do this!
}

// ✅ Go way: Return errors
if err != nil {
    return err
}
```

### 2. Making Everything an Interface
```go
// ❌ Java habit: Interface for everything
type UserService interface {
    GetUser(id int) (*User, error)
}

type userService struct{}

// ✅ Go way: Start with concrete type, extract interface when needed
type UserService struct{}
```

### 3. Ignoring Error Returns
```go
// ❌ JavaScript/Python habit: Ignoring errors
json.Unmarshal(data, &result)

// ✅ Go way: Always check errors
if err := json.Unmarshal(data, &result); err != nil {
    return err
}
```

### 4. Not Using defer
```go
// ❌ C++/Java habit: Manual cleanup
f, err := os.Open(file)
if err != nil {
    return err
}
result := process(f)
f.Close()
return result

// ✅ Go way: defer
f, err := os.Open(file)
if err != nil {
    return err
}
defer f.Close()
return process(f)
```

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Go Proverbs](https://go-proverbs.github.io/)
- [Standard Library](https://pkg.go.dev/std) - Best practices by example

## Tools

Essential tools for idiomatic Go:

```bash
go fmt          # Format code
go vet          # Static analysis
golangci-lint   # Comprehensive linter
go test -race   # Race detector
go test -cover  # Coverage analysis
```
