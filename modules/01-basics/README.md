# Module 01: Go Basics for Experienced Developers

## 🎯 Learning Objectives

By completing this module, you will:
- Understand Go's syntax and how it differs from your background language
- Master Go's type system and zero values
- Learn Go's approach to control flow
- Understand package organization and visibility
- Write your first idiomatic Go programs

## 📚 Prerequisites

- Experience with at least one programming language
- Basic understanding of programming concepts
- Go 1.21+ installed

## 🗺️ Module Overview

### 1. Hello, Go! - The Basics
**Coming from Python?** No `if __name__ == "__main__"`, just a `main()` function in package `main`.  
**Coming from Java?** No classes required - just functions and packages.  
**Coming from C++?** No header files - package imports handle everything.  
**Coming from JavaScript?** No `var` hoisting surprises - explicit declarations with `var`, `const`, or `:=`.

### 2. Go's Type System
- **Static typing** with type inference
- **Zero values** - no null pointer exceptions at initialization
- **Structs** instead of classes
- **Pointers** without pointer arithmetic

### 3. Control Flow
- `if`, `for`, `switch` - but with Go twists
- No `while` or `do-while` - just `for`
- `defer` for cleanup
- No exceptions - errors are values

### 4. Functions
- Multiple return values
- Named return values
- Variadic functions
- First-class functions

## 📖 Key Concepts

### Package and Imports

```go
package main  // Executable packages use "main"

import (
    "fmt"      // Standard library
    "strings"  // Group imports in parentheses
)

// Exported (public): Uppercase first letter
func PublicFunction() {}

// Unexported (private): Lowercase first letter  
func privateFunction() {}
```

**Key Difference:** Go uses capitalization for visibility, not keywords like `public`/`private`.

### Variables and Zero Values

```go
var x int           // Zero value: 0
var s string        // Zero value: ""
var b bool          // Zero value: false
var p *int          // Zero value: nil

// Short declaration (type inference, only inside functions)
y := 42             // int
name := "Go"        // string

// Constants
const Pi = 3.14159
```

**Why Zero Values Matter:** No uninitialized variable bugs! Every type has a sensible default.

### Types and Structs

```go
// Define a struct (like a class without methods... yet)
type Person struct {
    Name string
    Age  int
}

// Create instances
p1 := Person{Name: "Alice", Age: 30}
p2 := Person{"Bob", 25}  // Positional, but discouraged

// Zero value of struct
var p3 Person  // Person{Name: "", Age: 0}
```

### Control Flow

```go
// if statement - no parentheses!
if age >= 18 {
    fmt.Println("Adult")
}

// if with initialization
if err := doSomething(); err != nil {
    // err is scoped to this if block
    return err
}

// for loop - the only loop in Go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while-style for
for condition {
    // ...
}

// infinite loop
for {
    // break or return to exit
}

// range over collections
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// switch - no fallthrough by default!
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("Almost weekend!")
default:
    fmt.Println("Another day")
}
```

### Functions

```go
// Basic function
func add(a int, b int) int {
    return a + b
}

// Multiple return values (idiomatic for errors!)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Named return values
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // naked return
}

// Variadic functions
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```

### Defer

```go
func readFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()  // Executed when function returns
    
    // Do work with file
    return nil
}
```

**Python equivalent:** `with` statement  
**Java equivalent:** `try-with-resources`  
**C++ equivalent:** RAII destructors

## 🏋️ Exercises

Work through the exercises in the `exercises/` directory:

1. **exercise1_variables.go** - Practice with variables, types, and zero values
2. **exercise2_functions.go** - Write functions with multiple return values
3. **exercise3_control_flow.go** - Master if, for, and switch
4. **exercise4_structs.go** - Work with structs and methods (preview!)

Each exercise has intentional bugs marked with `// BUG:` comments. Fix them to make the tests pass!

### How to Work Through Exercises

```bash
# Run tests for an exercise (they will fail)
go test -v exercises/exercise1_variables_test.go exercises/exercise1_variables.go

# Fix the bugs in the exercise file
# Then run tests again

# Compare with solution
diff exercises/exercise1_variables.go solutions/exercise1_variables.go
```

## 🎓 Common Pitfalls

### 1. Unused Variables and Imports
```go
// ERROR: Unused variables cause compilation errors
func example() {
    x := 5  // declared but not used - won't compile!
}
```
**Solution:** Use `_` for intentionally unused values or remove them.

### 2. Short Declaration Scope
```go
if x := getValue(); x > 0 {
    // x is only available here
}
// x is not available here
```

### 3. Loop Variable Capture
```go
for i := 0; i < 10; i++ {
    go func() {
        fmt.Println(i)  // BUG: All goroutines will see the final value
    }()
}
```
**Solution:** Pass `i` as a parameter or create a new variable in the loop.

### 4. Capitalization Matters
```go
type person struct {  // unexported
    name string       // unexported field
}

type Person struct {  // exported
    Name string       // exported field
}
```

## 📚 Additional Resources

- [A Tour of Go](https://go.dev/tour/) - Interactive tutorial
- [Effective Go](https://go.dev/doc/effective_go) - Official style guide
- [Go by Example](https://gobyexample.com/) - Hands-on examples
- [Go Playground](https://go.dev/play/) - Run Go code in your browser

## ✅ Module Checklist

- [ ] Read through all concepts
- [ ] Run and understand all examples
- [ ] Complete exercise1_variables.go
- [ ] Complete exercise2_functions.go
- [ ] Complete exercise3_control_flow.go
- [ ] Complete exercise4_structs.go
- [ ] Compare your solutions with provided solutions
- [ ] Understand why the solution is idiomatic

## ⏭️ Next Module

Once you're comfortable with the basics, move on to:  
**[Module 02: Types and Interfaces](../02-types-interfaces/)** - Understand Go's powerful interface system

---

**Questions?** Open an issue or check the [CONTRIBUTING.md](../../CONTRIBUTING.md) guide!
