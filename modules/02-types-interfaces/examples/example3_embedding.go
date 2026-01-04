// Package examples demonstrates Go's embedding and composition for experienced developers.
//
// This file shows:
// - Struct embedding vs inheritance
// - Method promotion and overriding
// - Interface embedding
// - Composition patterns
// - Multiple embedding
package examples

import (
	"fmt"
	"time"
)

// Engine represents a basic engine.
//
// Coming from OOP: Instead of inheritance, Go uses composition.
// This is the "has-a" relationship instead of "is-a".
type Engine struct {
	Horsepower int
	Type       string
}

// Start starts the engine.
func (e Engine) Start() {
	fmt.Printf("Starting %s engine with %d HP\n", e.Type, e.Horsepower)
}

// Stop stops the engine.
func (e Engine) Stop() {
	fmt.Printf("Stopping %s engine\n", e.Type)
}

// Car embeds Engine - it "has-an" Engine.
type Car struct {
	Engine // Embedded field - no field name, just type
	Brand  string
	Model  string
}

// Drive is a method specific to Car.
func (c Car) Drive() {
	fmt.Printf("Driving %s %s\n", c.Brand, c.Model)
}

// BasicEmbedding demonstrates basic struct embedding.
func BasicEmbedding() {
	fmt.Println("=== Basic Embedding ===")

	car := Car{
		Engine: Engine{
			Horsepower: 300,
			Type:       "V6",
		},
		Brand: "Toyota",
		Model: "Camry",
	}

	// Access embedded fields directly (promoted)
	fmt.Printf("Car: %s %s with %d HP %s engine\n",
		car.Brand, car.Model, car.Horsepower, car.Type)

	// Call promoted methods from Engine
	car.Start()
	car.Drive()
	car.Stop()

	// Can still access embedded struct explicitly
	fmt.Printf("Engine type: %s\n", car.Engine.Type)
}

// ElectricMotor represents an electric motor.
type ElectricMotor struct {
	Kilowatts  int
	BatteryKWH float64
}

// Start starts the electric motor.
func (e ElectricMotor) Start() {
	fmt.Printf("Starting electric motor with %d kW\n", e.Kilowatts)
}

// Charge charges the electric motor.
func (e ElectricMotor) Charge() {
	fmt.Printf("Charging battery (%.1f kWh capacity)\n", e.BatteryKWH)
}

// HybridCar embeds both Engine and ElectricMotor.
type HybridCar struct {
	Engine
	ElectricMotor
	Brand string
	Model string
}

// MultipleEmbedding demonstrates embedding multiple structs.
func MultipleEmbedding() {
	fmt.Println("\n=== Multiple Embedding ===")

	hybrid := HybridCar{
		Engine: Engine{
			Horsepower: 200,
			Type:       "I4",
		},
		ElectricMotor: ElectricMotor{
			Kilowatts:  100,
			BatteryKWH: 15.5,
		},
		Brand: "Toyota",
		Model: "Prius",
	}

	fmt.Printf("Hybrid: %s %s\n", hybrid.Brand, hybrid.Model)
	fmt.Printf("Gas engine: %d HP, Electric motor: %d kW\n",
		hybrid.Horsepower, hybrid.Kilowatts)

	// Ambiguous: both Engine and ElectricMotor have Start() method
	// Must specify which one explicitly
	hybrid.Engine.Start()
	hybrid.ElectricMotor.Start()

	// Unique methods are still promoted
	hybrid.Stop()   // From Engine
	hybrid.Charge() // From ElectricMotor
}

// MethodOverriding demonstrates overriding promoted methods.
type SportsCar struct {
	Engine
	Brand    string
	Model    string
	TopSpeed int
}

// Start overrides the Engine's Start method.
func (s SportsCar) Start() {
	fmt.Printf("🏎️  Starting %s %s (Top speed: %d mph)\n",
		s.Brand, s.Model, s.TopSpeed)
	fmt.Printf("Engine: %d HP %s\n", s.Horsepower, s.Type)
}

// OverridingMethods demonstrates method overriding.
func OverridingMethods() {
	fmt.Println("\n=== Overriding Methods ===")

	sports := SportsCar{
		Engine: Engine{
			Horsepower: 500,
			Type:       "V8",
		},
		Brand:    "Ferrari",
		Model:    "F8 Tributo",
		TopSpeed: 211,
	}

	// Calls SportsCar's Start, not Engine's Start
	sports.Start()

	// Can still call Engine's Start explicitly
	fmt.Println("\nCalling embedded Engine's Start:")
	sports.Engine.Start()
}

// Logger provides logging functionality.
type Logger interface {
	Log(message string)
}

// TimestampLogger adds timestamps to logs.
type TimestampLogger struct{}

// Log logs a message with timestamp.
func (l TimestampLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), message)
}

// Service has logging capability through embedding.
type Service struct {
	Logger // Embedded interface
	Name   string
}

// Process demonstrates using the embedded interface.
func (s Service) Process() {
	s.Log(fmt.Sprintf("Processing in %s service", s.Name))
}

// EmbeddingInterfaces demonstrates embedding interfaces in structs.
func EmbeddingInterfaces() {
	fmt.Println("\n=== Embedding Interfaces ===")

	service := Service{
		Logger: TimestampLogger{},
		Name:   "UserAuth",
	}

	service.Process()
	service.Log("Additional log message")
}

// Reader defines a reading interface.
type Reader interface {
	Read() string
}

// Writer defines a writing interface.
type Writer interface {
	Write(data string)
}

// ReadWriter combines Reader and Writer through embedding.
//
// Coming from Java: Like extending multiple interfaces.
// Coming from Go: Cleaner than listing all methods.
type ReadWriter interface {
	Reader
	Writer
}

// Closer defines a closing interface.
type Closer interface {
	Close() error
}

// ReadWriteCloser combines multiple interfaces.
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// FileHandler implements ReadWriteCloser.
type FileHandler struct {
	filename string
	data     string
}

// Read implements Reader interface.
func (f *FileHandler) Read() string {
	return f.data
}

// Write implements Writer interface.
func (f *FileHandler) Write(data string) {
	f.data = data
}

// Close implements Closer interface.
func (f *FileHandler) Close() error {
	fmt.Printf("Closing file: %s\n", f.filename)
	return nil
}

// ComposingInterfaces demonstrates interface composition.
func ComposingInterfaces() {
	fmt.Println("\n=== Composing Interfaces ===")

	handler := &FileHandler{filename: "data.txt"}

	// Can use as ReadWriteCloser
	var rwc ReadWriteCloser = handler
	rwc.Write("Hello, Go!")
	data := rwc.Read()
	fmt.Printf("Data: %s\n", data)
	rwc.Close()

	// Can also use as just Reader or Writer
	var r Reader = handler
	fmt.Printf("Reading: %s\n", r.Read())

	var w Writer = handler
	w.Write("Updated data")
}

// Base is a base struct with common fields.
type Base struct {
	ID        int
	CreatedAt time.Time
}

// AppUser embeds Base to get common fields.
type AppUser struct {
	Base
	Username string
	Email    string
}

// Product embeds Base to get common fields.
type Product struct {
	Base
	Name  string
	Price float64
}

// ReusableComponents demonstrates reusing common fields.
func ReusableComponents() {
	fmt.Println("\n=== Reusable Components ===")

	user := AppUser{
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Username: "alice",
		Email:    "alice@example.com",
	}

	product := Product{
		Base: Base{
			ID:        101,
			CreatedAt: time.Now(),
		},
		Name:  "Laptop",
		Price: 999.99,
	}

	fmt.Printf("User: ID=%d, Username=%s, Created=%s\n",
		user.ID, user.Username, user.CreatedAt.Format("2006-01-02"))

	fmt.Printf("Product: ID=%d, Name=%s, Created=%s\n",
		product.ID, product.Name, product.CreatedAt.Format("2006-01-02"))
}

// Validator is an interface for validation.
type Validator interface {
	Validate() error
}

// Saveable is an interface for saving.
type Saveable interface {
	Save() error
}

// Model combines Validator and Saveable.
type Model interface {
	Validator
	Saveable
}

// Account implements Model interface through embedding.
type Account struct {
	Base
	Balance float64
}

// Validate implements Validator interface.
func (a Account) Validate() error {
	if a.Balance < 0 {
		return fmt.Errorf("balance cannot be negative")
	}
	return nil
}

// Save implements Saveable interface.
func (a Account) Save() error {
	fmt.Printf("Saving account %d with balance %.2f\n", a.ID, a.Balance)
	return nil
}

// ProcessModel demonstrates working with composite interfaces.
func ProcessModel(m Model) {
	if err := m.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return
	}

	if err := m.Save(); err != nil {
		fmt.Printf("Save error: %v\n", err)
		return
	}

	fmt.Println("Model processed successfully")
}

// CompositeInterfaces demonstrates using composite interfaces.
func CompositeInterfaces() {
	fmt.Println("\n=== Composite Interfaces ===")

	account := Account{
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Balance: 1000.50,
	}

	ProcessModel(account)

	invalidAccount := Account{
		Base: Base{
			ID:        2,
			CreatedAt: time.Now(),
		},
		Balance: -100, // Invalid
	}

	ProcessModel(invalidAccount)
}

// ZeroValueEmbedding demonstrates zero value behavior with embedding.
func ZeroValueEmbedding() {
	fmt.Println("\n=== Zero Value Embedding ===")

	// Zero value of Car has zero value Engine embedded
	var car Car
	fmt.Printf("Zero value car: %+v\n", car)
	fmt.Printf("Embedded engine: %+v\n", car.Engine)

	// Methods can still be called on zero values
	car.Start() // Prints "Starting  engine with 0 HP"
}
