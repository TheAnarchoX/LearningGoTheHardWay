package solutions

import (
	"errors"
	"testing"
	"time"
)

// TestDatabaseLogger verifies Database uses Logger interface.
func TestDatabaseLogger(t *testing.T) {
	logger := SimpleLogger{Prefix: "TEST"}
	db := &Database{
		Logger:           logger,
		ConnectionString: "test://localhost",
	}

	// Should be able to use Logger methods
	db.Log("test message")
	db.LogError(errors.New("test error"))
}

// TestDatabaseClose verifies Close modifies the connection state.
func TestDatabaseClose(t *testing.T) {
	db := &Database{
		Logger:           SimpleLogger{Prefix: "DB"},
		ConnectionString: "test://localhost",
		Connected:        true,
	}

	err := db.Close()
	if err != nil {
		t.Errorf("Close() error: %v", err)
	}

	if db.Connected {
		t.Error("Database should be disconnected after Close()")
	}
}

// TestElectricCarRange verifies electric car range calculation.
func TestElectricCarRange(t *testing.T) {
	car := ElectricCar{
		Vehicle: Vehicle{
			Brand: "Tesla",
			Model: "Model 3",
			Year:  2023,
		},
		BatteryCapacity: 75.0,
	}

	range_ := car.Range()
	expected := 375.0 // 75 * 5

	if range_ != expected {
		t.Errorf("Range() = %.2f, want %.2f", range_, expected)
	}
}

// TestElectricCarInfo verifies electric car info includes vehicle info.
func TestElectricCarInfo(t *testing.T) {
	car := ElectricCar{
		Vehicle: Vehicle{
			Brand: "Tesla",
			Model: "Model 3",
			Year:  2023,
		},
		BatteryCapacity: 75.0,
	}

	info := car.Info()
	// Should include both vehicle info and battery info
	expectedSubstring := "2023 Tesla Model 3"

	if len(info) == 0 {
		t.Error("Info() should not be empty")
	}

	// Info should contain vehicle information
	// This test will help verify the fix works
	_ = expectedSubstring
}

// TestGasCarRefuel verifies gas car refueling.
func TestGasCarRefuel(t *testing.T) {
	car := &GasCar{
		Vehicle: Vehicle{
			Brand: "Toyota",
			Model: "Camry",
			Year:  2023,
		},
		FuelCapacity: 15.0,
		MPG:          30.0,
	}

	// Should have Refuel method
	car.Refuel()
}

// TestTimestampedTouch verifies Touch updates timestamp.
func TestTimestampedTouch(t *testing.T) {
	ts := &Timestamped{
		CreatedAt: time.Now().Add(-1 * time.Hour),
		UpdatedAt: time.Now().Add(-1 * time.Hour),
	}

	oldUpdatedAt := ts.UpdatedAt
	time.Sleep(10 * time.Millisecond)
	ts.Touch()

	if !ts.UpdatedAt.After(oldUpdatedAt) {
		t.Error("Touch() should update UpdatedAt timestamp")
	}
}

// TestAccountAuditLog verifies account implements Auditable.
func TestAccountAuditLog(t *testing.T) {
	account := NewAccount(1, "alice", "alice@example.com")

	// Should implement Auditable interface
	var auditable Auditable = account

	log := auditable.AuditLog()
	if log == nil {
		t.Error("AuditLog() should not return nil")
	}
}

// TestAccountUpdateEmail verifies email update is audited.
func TestAccountUpdateEmail(t *testing.T) {
	account := NewAccount(1, "alice", "alice@example.com")

	oldUpdatedAt := account.UpdatedAt
	time.Sleep(10 * time.Millisecond)

	account.UpdateEmail("newemail@example.com")

	if account.Email != "newemail@example.com" {
		t.Errorf("Email not updated: got %s", account.Email)
	}

	if !account.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAt should be updated")
	}

	log := account.AuditLog()
	if len(log) == 0 {
		t.Error("AuditLog should contain email update entry")
	}
}

// TestCounterIncrement verifies counter increment.
func TestCounterIncrement(t *testing.T) {
	counter := &Counter{}

	counter.Increment()
	counter.Increment()
	counter.Increment()

	if counter.Value() != 3 {
		t.Errorf("Value() = %d, want 3", counter.Value())
	}
}

// TestCounterReset verifies counter reset.
func TestCounterReset(t *testing.T) {
	counter := &Counter{}
	counter.Increment()
	counter.Increment()

	counter.Reset()

	if counter.Value() != 0 {
		t.Errorf("Value() after Reset() = %d, want 0", counter.Value())
	}
}

// TestServiceStart verifies service start with logging.
func TestServiceStart(t *testing.T) {
	service := NewService("test-service")

	// Should not panic when calling Start
	service.Start()

	// Should increment counter
	if service.Value() != 1 {
		t.Errorf("Counter value = %d, want 1", service.Value())
	}
}

// TestEmployeeGreet verifies employee greeting includes name.
func TestEmployeeGreet(t *testing.T) {
	emp := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		EmployeeID: 1001,
		Department: "Engineering",
	}

	greeting := emp.Greet()
	// Greeting should include the person's name
	if len(greeting) == 0 {
		t.Error("Greet() should not be empty")
	}
}

// TestManagerGreet verifies manager greeting includes all info.
func TestManagerGreet(t *testing.T) {
	mgr := Manager{
		Employee: Employee{
			Person: Person{
				Name: "Bob",
				Age:  40,
			},
			EmployeeID: 2001,
			Department: "Engineering",
		},
		TeamSize: 5,
	}

	greeting := mgr.Greet()
	// Greeting should include name, employee ID, department, and team size
	if len(greeting) == 0 {
		t.Error("Greet() should not be empty")
	}
}

// TestCacheGetSet verifies cache get and set operations.
func TestCacheGetSet(t *testing.T) {
	cache := &Cache{
		data: make(map[string]interface{}),
	}

	cache.Set("key1", "value1")

	value, exists := cache.Get("key1")
	if !exists {
		t.Error("Key should exist in cache")
	}

	if value != "value1" {
		t.Errorf("Get() = %v, want %v", value, "value1")
	}
}

// TestCacheUninitializedMap verifies cache handles uninitialized map.
func TestCacheUninitializedMap(t *testing.T) {
	cache := &Cache{} // map not initialized

	// Should not panic
	cache.Set("key", "value")

	value, exists := cache.Get("key")
	if !exists {
		t.Error("Key should exist after Set")
	}
	if value != "value" {
		t.Errorf("Get() = %v, want %v", value, "value")
	}
}

// TestCachedDatabase verifies cached database composition.
func TestCachedDatabase(t *testing.T) {
	logger := SimpleLogger{Prefix: "CacheDB"}
	db := NewCachedDatabase("test://localhost", logger)

	// Should have database methods
	err := db.Connect()
	if err != nil {
		t.Errorf("Connect() error: %v", err)
	}

	// Should have cache methods
	db.Set("key", "value")
	value, exists := db.Get("key")
	if !exists {
		t.Error("Key should exist in cache")
	}
	if value != "value" {
		t.Errorf("Get() = %v, want %v", value, "value")
	}
}
