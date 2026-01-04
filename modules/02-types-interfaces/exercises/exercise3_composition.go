// Package exercises contains practice problems with intentional bugs.
//
// EXERCISE 3: Composition and Embedding
//
// Fix the bugs in this file to make all tests pass.
// The bugs are marked with // BUG: comments.
//
// Learning objectives:
// - Use struct embedding correctly
// - Understand method promotion
// - Handle method overriding
// - Work with embedded interfaces
// - Build complex types through composition
package exercises

import (
	"fmt"
	"time"
)

// Logger provides logging functionality.
type Logger interface {
	Log(message string)
	LogError(err error)
}

// SimpleLogger implements basic logging.
type SimpleLogger struct {
	Prefix string
}

// Log logs a message with prefix.
func (l SimpleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.Prefix, message)
}

// LogError logs an error with prefix.
func (l SimpleLogger) LogError(err error) {
	fmt.Printf("[%s] ERROR: %v\n", l.Prefix, err)
}

// Database represents a database connection.
type Database struct {
	// BUG: Should embed Logger interface, not SimpleLogger struct
	SimpleLogger
	ConnectionString string
	Connected        bool
}

// Connect connects to the database.
func (db *Database) Connect() error {
	db.Connected = true
	db.Log("Database connected")
	return nil
}

// Close closes the database connection.
// BUG: Should use pointer receiver like Connect
func (db Database) Close() error {
	db.Connected = false // BUG: This won't modify the original
	db.Log("Database closed")
	return nil
}

// Vehicle is a base type for vehicles.
type Vehicle struct {
	Brand string
	Model string
	Year  int
}

// Info returns vehicle information.
func (v Vehicle) Info() string {
	return fmt.Sprintf("%d %s %s", v.Year, v.Brand, v.Model)
}

// ElectricCar embeds Vehicle.
type ElectricCar struct {
	Vehicle
	BatteryCapacity float64
}

// Range calculates the range of the electric car.
// BUG: Should return battery capacity * 5, not just battery capacity
func (e ElectricCar) Range() float64 {
	return e.BatteryCapacity
}

// Info overrides Vehicle.Info to include battery info.
// BUG: Doesn't call the embedded Vehicle.Info method
func (e ElectricCar) Info() string {
	// BUG: Should use e.Vehicle.Info() to get base info
	return fmt.Sprintf("Electric car with %.1f kWh battery", e.BatteryCapacity)
}

// GasCar embeds Vehicle.
type GasCar struct {
	Vehicle
	FuelCapacity float64
	MPG          float64
}

// Range calculates the range of the gas car.
func (g GasCar) Range() float64 {
	return g.FuelCapacity * g.MPG
}

// Refuel refuels the car.
// BUG: Missing implementation
// TODO: Implement this method to log refueling

// Fleet manages a collection of vehicles.
type Fleet struct {
	Name     string
	Vehicles []Vehicle // BUG: Should use interface for vehicles that have Range()
}

// AddVehicle adds a vehicle to the fleet.
// BUG: Type signature is too restrictive
func (f *Fleet) AddVehicle(v Vehicle) {
	f.Vehicles = append(f.Vehicles, v)
}

// TotalRange calculates the total range of all vehicles.
// BUG: Can't call Range() on Vehicle
func (f *Fleet) TotalRange() float64 {
	total := 0.0
	for _, vehicle := range f.Vehicles {
		// BUG: Vehicle doesn't have Range() method
		total += vehicle.Range()
	}
	return total
}

// Auditable defines types that can be audited.
type Auditable interface {
	AuditLog() []string
}

// Timestamped adds timestamp tracking.
type Timestamped struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch updates the UpdatedAt timestamp.
// BUG: Should use pointer receiver to modify the struct
func (t Timestamped) Touch() {
	t.UpdatedAt = time.Now()
}

// Account represents a user account with audit trail.
type Account struct {
	Timestamped
	ID       int
	Username string
	Email    string
	auditLog []string // BUG: Should be exported to satisfy Auditable interface
}

// NewAccount creates a new account.
func NewAccount(id int, username, email string) *Account {
	now := time.Now()
	return &Account{
		Timestamped: Timestamped{
			CreatedAt: now,
			UpdatedAt: now,
		},
		ID:       id,
		Username: username,
		Email:    email,
		auditLog: []string{},
	}
}

// UpdateEmail updates the account email.
func (a *Account) UpdateEmail(email string) {
	a.Email = email
	a.Touch()
	// BUG: auditLog is unexported, can't be accessed by interface
	a.auditLog = append(a.auditLog, fmt.Sprintf("Email updated to %s", email))
}

// AuditLog returns the audit log.
// BUG: Returns unexported field
func (a *Account) AuditLog() []string {
	return a.auditLog
}

// Counter provides counting functionality.
type Counter struct {
	count int
}

// Increment increases the counter.
// BUG: Should use pointer receiver to modify count
func (c Counter) Increment() {
	c.count++
}

// Value returns the current count.
func (c Counter) Value() int {
	return c.count
}

// Reset resets the counter to zero.
// BUG: Should use pointer receiver
func (c Counter) Reset() {
	c.count = 0
}

// Stats tracks statistics.
type Stats struct {
	Counter // Embedded
	Name    string
}

// Service combines multiple concerns through embedding.
type Service struct {
	Logger // BUG: Embedded interface but not initialized
	Stats  // Embedded struct
	Name   string
}

// NewService creates a new service.
// BUG: Doesn't initialize the Logger
func NewService(name string) *Service {
	return &Service{
		Name: name,
		Stats: Stats{
			Name: name + "-stats",
		},
	}
}

// Start starts the service.
func (s *Service) Start() {
	// BUG: Logger might be nil, causing panic
	s.Log(fmt.Sprintf("Starting service: %s", s.Name))
	s.Increment()
}

// Person represents a person.
type Person struct {
	Name string
	Age  int
}

// Greet returns a greeting.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Employee embeds Person.
type Employee struct {
	Person
	EmployeeID int
	Department string
}

// Greet overrides Person.Greet.
// BUG: Doesn't call embedded Person.Greet
func (e Employee) Greet() string {
	// BUG: Should incorporate Person.Greet()
	return fmt.Sprintf("Employee #%d in %s", e.EmployeeID, e.Department)
}

// Manager embeds Employee.
type Manager struct {
	Employee
	TeamSize int
}

// Greet overrides Employee.Greet.
func (m Manager) Greet() string {
	// BUG: Calls Employee.Greet but Employee.Greet doesn't use Person.Greet
	return fmt.Sprintf("%s, managing team of %d", m.Employee.Greet(), m.TeamSize)
}

// Cache provides caching functionality.
type Cache struct {
	data map[string]interface{}
}

// Get retrieves a value from the cache.
// BUG: Doesn't check if map is initialized
func (c *Cache) Get(key string) (interface{}, bool) {
	// BUG: Will panic if data map is nil
	value, exists := c.data[key]
	return value, exists
}

// Set stores a value in the cache.
// BUG: Doesn't initialize map if nil
func (c *Cache) Set(key string, value interface{}) {
	// BUG: Will panic if data map is nil
	c.data[key] = value
}

// CachedDatabase combines Database and Cache.
type CachedDatabase struct {
	*Database // BUG: Should be Database, not *Database for consistency
	Cache
}

// NewCachedDatabase creates a new cached database.
func NewCachedDatabase(connStr string, logger Logger) *CachedDatabase {
	return &CachedDatabase{
		Database: &Database{
			// BUG: Embedding SimpleLogger instead of using Logger interface
			SimpleLogger:     SimpleLogger{Prefix: "DB"},
			ConnectionString: connStr,
		},
		Cache: Cache{
			data: make(map[string]interface{}),
		},
	}
}
