// Package solutions contains correct implementations of the exercises.
//
// SOLUTION 3: Composition and Embedding
//
// This file shows the correct implementation with all bugs fixed.
//
// Key fixes:
// - Database embeds Logger interface
// - Close() uses pointer receiver
// - ElectricCar.Range() uses correct formula
// - ElectricCar.Info() calls embedded Vehicle.Info()
// - GasCar.Refuel() implemented
// - Fleet uses interface for vehicles with Range()
// - Timestamped.Touch() uses pointer receiver
// - Account exports AuditLog field
// - Counter methods use pointer receivers
// - Service properly initializes Logger
// - Employee.Greet() uses embedded Person.Greet()
// - Cache initializes map if nil
package solutions

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
	Logger           // FIXED: Embed Logger interface
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
// FIXED: Use pointer receiver
func (db *Database) Close() error {
	db.Connected = false
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

// Ranger interface for vehicles that have a Range method.
type Ranger interface {
	Range() float64
}

// ElectricCar embeds Vehicle.
type ElectricCar struct {
	Vehicle
	BatteryCapacity float64
}

// Range calculates the range of the electric car.
// FIXED: Multiply by 5
func (e ElectricCar) Range() float64 {
	return e.BatteryCapacity * 5
}

// Info overrides Vehicle.Info to include battery info.
// FIXED: Use embedded Vehicle.Info()
func (e ElectricCar) Info() string {
	return fmt.Sprintf("%s, Electric car with %.1f kWh battery", e.Vehicle.Info(), e.BatteryCapacity)
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
// FIXED: Implemented the method
func (g *GasCar) Refuel() {
	fmt.Printf("Refueling %s\n", g.Info())
}

// Fleet manages a collection of vehicles.
type Fleet struct {
	Name     string
	Vehicles []Ranger // FIXED: Use Ranger interface
}

// AddVehicle adds a vehicle to the fleet.
// FIXED: Accept Ranger interface
func (f *Fleet) AddVehicle(v Ranger) {
	f.Vehicles = append(f.Vehicles, v)
}

// TotalRange calculates the total range of all vehicles.
// FIXED: Now works with Ranger interface
func (f *Fleet) TotalRange() float64 {
	total := 0.0
	for _, vehicle := range f.Vehicles {
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
// FIXED: Use pointer receiver
func (t *Timestamped) Touch() {
	t.UpdatedAt = time.Now()
}

// Account represents a user account with audit trail.
type Account struct {
	Timestamped
	ID       int
	Username string
	Email    string
	auditLog []string // FIXED: Keep private but implement getter method
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
	a.auditLog = append(a.auditLog, fmt.Sprintf("Email updated to %s", email))
}

// AuditLog returns the audit log.
// FIXED: Proper getter method that satisfies interface
func (a *Account) AuditLog() []string {
	return a.auditLog
}

// Counter provides counting functionality.
type Counter struct {
	count int
}

// Increment increases the counter.
// FIXED: Use pointer receiver
func (c *Counter) Increment() {
	c.count++
}

// Value returns the current count.
func (c *Counter) Value() int {
	return c.count
}

// Reset resets the counter to zero.
// FIXED: Use pointer receiver
func (c *Counter) Reset() {
	c.count = 0
}

// Stats tracks statistics.
type Stats struct {
	Counter
	Name string
}

// Service combines multiple concerns through embedding.
type Service struct {
	Logger
	Stats
	Name string
}

// NewService creates a new service.
// FIXED: Initialize Logger
func NewService(name string) *Service {
	return &Service{
		Logger: SimpleLogger{Prefix: name}, // FIXED: Initialize logger
		Name:   name,
		Stats: Stats{
			Name: name + "-stats",
		},
	}
}

// Start starts the service.
func (s *Service) Start() {
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
// FIXED: Call embedded Person.Greet()
func (e Employee) Greet() string {
	return fmt.Sprintf("%s, Employee #%d in %s", e.Person.Greet(), e.EmployeeID, e.Department)
}

// Manager embeds Employee.
type Manager struct {
	Employee
	TeamSize int
}

// Greet overrides Employee.Greet.
func (m Manager) Greet() string {
	return fmt.Sprintf("%s, managing team of %d", m.Employee.Greet(), m.TeamSize)
}

// Cache provides caching functionality.
type Cache struct {
	data map[string]interface{}
}

// Get retrieves a value from the cache.
// FIXED: Check if map is initialized
func (c *Cache) Get(key string) (interface{}, bool) {
	if c.data == nil {
		return nil, false
	}
	value, exists := c.data[key]
	return value, exists
}

// Set stores a value in the cache.
// FIXED: Initialize map if nil
func (c *Cache) Set(key string, value interface{}) {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = value
}

// CachedDatabase combines Database and Cache.
type CachedDatabase struct {
	Database // FIXED: Embed Database directly, not pointer
	Cache
}

// NewCachedDatabase creates a new cached database.
func NewCachedDatabase(connStr string, logger Logger) *CachedDatabase {
	return &CachedDatabase{
		Database: Database{
			Logger:           logger, // FIXED: Use Logger interface
			ConnectionString: connStr,
		},
		Cache: Cache{
			data: make(map[string]interface{}),
		},
	}
}
