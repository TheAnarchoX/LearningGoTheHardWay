# Common Go Gotchas for Developers from Other Languages

A guide to help you avoid common mistakes when coming to Go from Python, Java, C++, or JavaScript.

## From Python

### 1. The Infamous Loop Variable
```python
# Python: Works as expected
for i in range(5):
    threading.Thread(target=lambda: print(i)).start()
```

```go
// ❌ Go: All goroutines might print 5
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)  // Captures loop variable!
    }()
}

// ✅ Fix: Pass as parameter
for i := 0; i < 5; i++ {
    go func(id int) {
        fmt.Println(id)
    }(i)
}
```

### 2. No List Comprehensions
```python
# Python
squares = [x**2 for x in range(10)]
```

```go
// Go: Use explicit loop
squares := make([]int, 10)
for i := 0; i < 10; i++ {
    squares[i] = i * i
}
```

### 3. No Duck Typing
```python
# Python: If it quacks like a duck...
def process(obj):
    obj.do_something()  # Works if obj has method
```

```go
// Go: Types must satisfy interface explicitly
type Processor interface {
    DoSomething()
}

func process(p Processor) {
    p.DoSomething()
}
```

### 4. Zero Values, Not None
```python
# Python
x = None
if x is None:
    pass
```

```go
// Go: Each type has zero value
var x int       // 0, not nil
var s string    // "", not nil
var p *int      // nil
var slice []int // nil (but different from []int{})
```

## From Java

### 1. No Classes, Use Structs
```java
// Java
public class User {
    private String name;
    
    public String getName() {
        return name;
    }
}
```

```go
// Go: Exported fields are public
type User struct {
    Name string  // Public (exported)
    age  int     // Private (unexported)
}
```

### 2. No Constructors
```java
// Java
public class Config {
    public Config(String host) {
        this.host = host;
    }
}
```

```go
// Go: Use factory functions
type Config struct {
    host string
}

func NewConfig(host string) *Config {
    return &Config{host: host}
}
```

### 3. Interfaces Are Implicit
```java
// Java: Explicit implementation
public class Dog implements Animal {
    public void speak() {
        System.out.println("Woof");
    }
}
```

```go
// Go: Implicit satisfaction
type Animal interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof")
}
// Dog implements Animal automatically!
```

### 4. No Generics (until Go 1.18)
```java
// Java
List<String> names = new ArrayList<>();
```

```go
// Go < 1.18: Use interface{} or code generation
names := make([]string, 0)

// Go >= 1.18: Generics available
func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

## From C++

### 1. No Pointer Arithmetic
```cpp
// C++: Pointer arithmetic
int arr[] = {1, 2, 3};
int* p = arr;
*(p + 1) = 5;  // OK
```

```go
// Go: No pointer arithmetic
arr := []int{1, 2, 3}
// &arr + 1  // Compilation error!
arr[1] = 5  // Use indexing
```

### 2. No Manual Memory Management
```cpp
// C++: Manual allocation
User* user = new User();
delete user;
```

```go
// Go: Garbage collected
user := &User{}
// No delete needed
```

### 3. No Inheritance
```cpp
// C++: Inheritance
class Animal {
    virtual void speak() = 0;
};

class Dog : public Animal {
    void speak() override { }
};
```

```go
// Go: Composition via embedding
type Animal struct {
    name string
}

func (a Animal) Speak() {
    fmt.Println("Some sound")
}

type Dog struct {
    Animal  // Embedding
    breed string
}
// Dog has Speak() method via embedding
```

### 4. No Constructors/Destructors
```cpp
// C++: RAII
class File {
    File(string name) { /* open */ }
    ~File() { /* close */ }
};
```

```go
// Go: defer for cleanup
func processFile(name string) error {
    f, err := os.Open(name)
    if err != nil {
        return err
    }
    defer f.Close()  // Like destructor
    
    // Process...
    return nil
}
```

## From JavaScript

### 1. No Undefined
```javascript
// JavaScript
let x;
console.log(x);  // undefined
```

```go
// Go: Zero values
var x int
fmt.Println(x)  // 0 (not undefined)
```

### 2. Static Typing
```javascript
// JavaScript
let x = 5;
x = "hello";  // OK
```

```go
// Go: Type is fixed
var x int = 5
x = "hello"  // Compilation error!
```

### 3. No Truthy/Falsy
```javascript
// JavaScript
if ("") { }     // false
if (0) { }      // false  
if ([]) { }     // true (!)
```

```go
// Go: Only bool in conditions
if "" { }       // Error: non-bool used as condition
if 0 { }        // Error
if []int{} { }  // Error

// Must be explicit
if len(s) == 0 { }  // OK
```

### 4. No var Hoisting
```javascript
// JavaScript
console.log(x);  // undefined (hoisted)
var x = 5;
```

```go
// Go: Must declare before use
fmt.Println(x)  // Error: undefined: x
var x int = 5
```

### 5. goroutines ≠ async/await
```javascript
// JavaScript: Single-threaded event loop
async function fetchData() {
    const data = await fetch(url);
}
```

```go
// Go: True parallelism
func fetchData() {
    go func() {
        // Runs in parallel on different thread!
        data := fetch(url)
    }()
}
```

## Cross-Language Gotchas

### 1. Maps Are Not Ordered
```go
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    // Order is random and different each run!
}
```

### 2. Slices Are References
```go
a := []int{1, 2, 3}
b := a
b[0] = 100
fmt.Println(a[0])  // 100 - slices share backing array!

// To copy:
b := make([]int, len(a))
copy(b, a)
```

### 3. Channel Operations Can Block
```go
ch := make(chan int)
ch <- 1  // Blocks forever if no receiver!

// Need goroutine or buffered channel
ch := make(chan int, 1)
ch <- 1  // OK - buffered
```

### 4. Range Copies Values
```go
type Large struct {
    data [1000]int
}

items := []Large{{}, {}}
for _, item := range items {
    // item is a COPY of the Large struct!
    item.data[0] = 1  // Doesn't modify original
}

// Use index instead
for i := range items {
    items[i].data[0] = 1  // Modifies original
}
```

### 5. Interface{} Is Not Any Type
```go
var x interface{} = "hello"
// x.Length()  // Error: interface{} has no methods

// Need type assertion
s := x.(string)
fmt.Println(len(s))  // OK
```

### 6. Unused Variables Are Errors
```go
func example() {
    x := 5  // Declared but not used
    // Compilation error: x declared and not used
}

// Use _ for intentionally unused
_, err := doSomething()
```

## Testing Differences

### No Assert Library Needed
```go
// Don't import assert libraries
// Use standard testing
func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

### Table-Driven Tests Are Idiomatic
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {2, 3, 5},
        {0, 0, 0},
        {-1, 1, 0},
    }
    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d",
                tt.a, tt.b, got, tt.want)
        }
    }
}
```

## Tips for Success

1. **Read the standard library** - It's idiomatic Go
2. **Run `go vet` and `golangci-lint`** - Catch common mistakes
3. **Use `go fmt`** - Never argue about formatting
4. **Enable the race detector** - `go test -race`
5. **Read Effective Go** - It's short and essential
6. **Don't fight the language** - Embrace Go's simplicity

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Common Mistakes](https://github.com/golang/go/wiki/CommonMistakes)
- [Go FAQ](https://go.dev/doc/faq)
