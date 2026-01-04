# Module 02: Types and Interfaces

## 🎯 Learning Objectives

By completing this module, you will:
- Understand Go's approach to object-oriented programming
- Master Go's interface system and implicit satisfaction
- Learn when to use structs, interfaces, and embedding
- Understand the empty interface and type assertions
- Apply Go's composition over inheritance philosophy

## 📚 Prerequisites

- Completed Module 01: Basics
- Understanding of OOP from other languages (Java, Python, C++, etc.)

## 🗺️ Module Overview

### Coming From Other Languages

**Python Developers:** Interfaces are like Protocol or ABC, but implicit - no need to declare inheritance!  
**Java Developers:** No `implements` keyword - if a type has the methods, it satisfies the interface.  
**C++ Developers:** Think of interfaces like abstract base classes, but lighter weight.  
**JavaScript Developers:** More structured than duck typing - compile-time checked.

## 📖 Key Concepts

### 1. Structs - Go's "Objects"

```go
// Define a struct type
type Person struct {
    Name string
    Age  int
}

// Method on a struct (value receiver)
func (p Person) Greet() string {
    return "Hello, I'm " + p.Name
}

// Method with pointer receiver (can modify)
func (p *Person) HaveBirthday() {
    p.Age++
}
```

**Key Difference:** Methods are defined outside the struct, not inside like classes.

### 2. Interfaces - The Go Way

```go
// Define an interface
type Speaker interface {
    Speak() string
}

// Any type with Speak() method implements Speaker
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Robot struct {
    Model string
}

func (r Robot) Speak() string {
    return "Beep boop"
}

// Both Dog and Robot satisfy Speaker interface!
func MakeNoise(s Speaker) {
    fmt.Println(s.Speak())
}
```

**Implicit Satisfaction:** No need to declare `implements Speaker` - if it has the methods, it satisfies the interface!

### 3. Empty Interface - interface{}

```go
// interface{} can hold any type
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

PrintAnything(42)
PrintAnything("hello")
PrintAnything(Person{Name: "Alice", Age: 30})
```

**Note:** In Go 1.18+, prefer `any` which is an alias for `interface{}`.

### 4. Type Assertions and Type Switches

```go
var value interface{} = "hello"

// Type assertion
str, ok := value.(string)
if ok {
    fmt.Println("It's a string:", str)
}

// Type switch
switch v := value.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Integer:", v)
default:
    fmt.Println("Unknown type")
}
```

### 5. Embedding - Composition Over Inheritance

```go
type Animal struct {
    Name string
}

func (a Animal) Eat() {
    fmt.Println(a.Name, "is eating")
}

// Dog "has-a" Animal, not "is-a" Animal
type Dog struct {
    Animal  // Embedding
    Breed string
}

// Promoted methods
dog := Dog{
    Animal: Animal{Name: "Buddy"},
    Breed: "Golden Retriever",
}
dog.Eat()  // Calls Animal.Eat()
```

**vs Inheritance:** This is composition, not inheritance. Dog has access to Animal's methods but can override them.

### 6. Interface Composition

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Compose interfaces
type ReadWriter interface {
    Reader
    Writer
}
```

### 7. Pointer vs Value Receivers

```go
// Value receiver - operates on a copy
func (p Person) CannotModify() {
    p.Age = 100  // Modifies the copy, not original
}

// Pointer receiver - can modify original
func (p *Person) CanModify() {
    p.Age = 100  // Modifies original
}
```

**Rule of Thumb:**
- Use pointer receivers when you need to modify the receiver
- Use pointer receivers for large structs (avoid copying)
- Be consistent - if one method uses pointer receiver, all should

## 🏋️ Exercises

Work through the exercises in the `exercises/` directory:

1. **exercise1_interfaces.go** - Implement interfaces for different shapes
2. **exercise2_type_assertions.go** - Work with type assertions and switches
3. **exercise3_embedding.go** - Use embedding to build complex types

Each exercise has bugs or TODOs. Fix them to make tests pass!

## 🎓 Common Pitfalls

### 1. Interface Nil Confusion
```go
var p *Person = nil
var s Speaker = p  // s is not nil! It has type *Person but value nil

if s != nil {  // This is true!
    s.Speak()  // But this will panic
}
```

### 2. Pointer vs Value in Interfaces
```go
type Speaker interface {
    Speak() string
}

func (p *Person) Speak() string {  // Pointer receiver
    return "Hi"
}

// This works
var p *Person = &Person{Name: "Alice"}
var s Speaker = p

// This doesn't work!
var p2 Person = Person{Name: "Bob"}
var s2 Speaker = p2  // Compile error: Person doesn't implement Speaker
```

### 3. Empty Interface Doesn't Mean Any Methods
```go
var anything interface{} = "hello"
// anything.Length()  // ERROR: interface{} has no methods
str := anything.(string)  // Need type assertion first
fmt.Println(len(str))
```

## 📚 Additional Resources

- [Effective Go - Interfaces](https://go.dev/doc/effective_go#interfaces)
- [Go Blog: Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)
- [Interface Values](https://research.swtch.com/interfaces)

## ✅ Module Checklist

- [ ] Understand struct methods
- [ ] Master interface definitions and satisfaction
- [ ] Practice with type assertions
- [ ] Use embedding effectively
- [ ] Complete all exercises
- [ ] Compare solutions with your implementation

## ⏭️ Next Module

**[Module 03: Concurrency Fundamentals](../03-concurrency-fundamentals/)** - Master goroutines and channels!
