package examples

import (
	"testing"
	"time"
)

// TestBasicEmbedding verifies basic embedding examples.
func TestBasicEmbedding(t *testing.T) {
	BasicEmbedding()
}

// TestMultipleEmbedding verifies multiple embedding examples.
func TestMultipleEmbedding(t *testing.T) {
	MultipleEmbedding()
}

// TestOverridingMethods verifies method overriding examples.
func TestOverridingMethods(t *testing.T) {
	OverridingMethods()
}

// TestEmbeddingInterfaces verifies embedding interfaces in structs.
func TestEmbeddingInterfaces(t *testing.T) {
	EmbeddingInterfaces()
}

// TestComposingInterfaces verifies interface composition examples.
func TestComposingInterfaces(t *testing.T) {
	ComposingInterfaces()
}

// TestReusableComponents verifies reusable component pattern.
func TestReusableComponents(t *testing.T) {
	ReusableComponents()
}

// TestCompositeInterfaces verifies composite interface examples.
func TestCompositeInterfaces(t *testing.T) {
	CompositeInterfaces()
}

// TestZeroValueEmbedding verifies zero value embedding behavior.
func TestZeroValueEmbedding(t *testing.T) {
	ZeroValueEmbedding()
}

// TestCarHasEngine verifies Car has promoted Engine fields.
func TestCarHasEngine(t *testing.T) {
	car := Car{
		Engine: Engine{
			Horsepower: 300,
			Type:       "V6",
		},
		Brand: "Toyota",
		Model: "Camry",
	}

	// Access promoted fields
	if car.Horsepower != 300 {
		t.Errorf("Expected horsepower 300, got %d", car.Horsepower)
	}

	if car.Type != "V6" {
		t.Errorf("Expected type V6, got %s", car.Type)
	}

	// Access through embedded struct
	if car.Engine.Horsepower != 300 {
		t.Errorf("Expected engine horsepower 300, got %d", car.Engine.Horsepower)
	}
}

// TestHybridCarMultipleEmbedding verifies multiple embedding.
func TestHybridCarMultipleEmbedding(t *testing.T) {
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

	// Can access both embedded structs' fields
	if hybrid.Horsepower != 200 {
		t.Errorf("Expected horsepower 200, got %d", hybrid.Horsepower)
	}

	if hybrid.Kilowatts != 100 {
		t.Errorf("Expected kilowatts 100, got %d", hybrid.Kilowatts)
	}
}

// TestFileHandlerImplementsReadWriteCloser verifies FileHandler satisfies interface.
func TestFileHandlerImplementsReadWriteCloser(t *testing.T) {
	handler := &FileHandler{filename: "test.txt"}

	// Verify it implements ReadWriteCloser
	var rwc ReadWriteCloser = handler

	// Test Write
	rwc.Write("test data")

	// Test Read
	data := rwc.Read()
	if data != "test data" {
		t.Errorf("Expected 'test data', got %q", data)
	}

	// Test Close
	if err := rwc.Close(); err != nil {
		t.Errorf("Close failed: %v", err)
	}
}

// TestUserEmbedsBase verifies AppUser embeds Base.
func TestUserEmbedsBase(t *testing.T) {
	user := AppUser{
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Username: "alice",
		Email:    "alice@example.com",
	}

	// Access promoted ID field
	if user.ID != 1 {
		t.Errorf("Expected ID 1, got %d", user.ID)
	}

	// Verify CreatedAt is accessible
	if user.CreatedAt.IsZero() {
		t.Error("CreatedAt should not be zero")
	}
}

// TestProductEmbedsBase verifies Product embeds Base.
func TestProductEmbedsBase(t *testing.T) {
	product := Product{
		Base: Base{
			ID:        101,
			CreatedAt: time.Now(),
		},
		Name:  "Laptop",
		Price: 999.99,
	}

	// Access promoted ID field
	if product.ID != 101 {
		t.Errorf("Expected ID 101, got %d", product.ID)
	}

	if product.Price != 999.99 {
		t.Errorf("Expected price 999.99, got %.2f", product.Price)
	}
}

// TestAccountImplementsModel verifies Account implements Model interface.
func TestAccountImplementsModel(t *testing.T) {
	account := Account{
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
		},
		Balance: 1000.50,
	}

	// Verify it implements Model interface
	var m Model = account

	// Test Validate
	if err := m.Validate(); err != nil {
		t.Errorf("Valid account failed validation: %v", err)
	}

	// Test Save
	if err := m.Save(); err != nil {
		t.Errorf("Save failed: %v", err)
	}
}

// TestAccountValidation verifies Account validation.
func TestAccountValidation(t *testing.T) {
	validAccount := Account{Balance: 100}
	if err := validAccount.Validate(); err != nil {
		t.Errorf("Valid account should pass validation: %v", err)
	}

	invalidAccount := Account{Balance: -100}
	if err := invalidAccount.Validate(); err == nil {
		t.Error("Invalid account should fail validation")
	}
}

// TestServiceEmbeddsLogger verifies Service embeds Logger interface.
func TestServiceEmbedsLogger(t *testing.T) {
	service := Service{
		Logger: TimestampLogger{},
		Name:   "TestService",
	}

	// Should be able to call Log method (promoted from Logger)
	service.Log("test message") // Should not panic
}

// TestEmbeddingInterfacesComposition verifies interface composition.
func TestEmbeddingInterfacesComposition(t *testing.T) {
	handler := &FileHandler{filename: "test.txt"}

	// Can assign to ReadWriteCloser
	var rwc ReadWriteCloser = handler
	if rwc == nil {
		t.Error("FileHandler should implement ReadWriteCloser")
	}

	// Can also assign to individual interfaces
	var r Reader = handler
	if r == nil {
		t.Error("FileHandler should implement Reader")
	}

	var w Writer = handler
	if w == nil {
		t.Error("FileHandler should implement Writer")
	}

	var c Closer = handler
	if c == nil {
		t.Error("FileHandler should implement Closer")
	}
}
