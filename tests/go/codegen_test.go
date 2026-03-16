package codegen_test

// This file tests the correctness of Go code generated from Concerto models.
//
// Strategy:
//   1. Import the generated package (compilation itself is the first test).
//   2. Instantiate each generated struct and verify field assignment works.
//   3. Verify JSON marshal / unmarshal round-trips preserve values.
//   4. Verify enum values exist as expected constants.
//
// The generated package is expected at: ../../generated/go/
// Its package name follows concerto convention: <namespace>_<version_underscored>
// e.g. namespace test@1.0.0 → package test_1_0_0

import (
	"encoding/json"
	"testing"
	"time"

	// Import the generated package.
	// Adjust this import path if your namespace / version differs.
	generated "github.com/Concerto-Go/generated/go"
)

// ---------------------------------------------------------------------------
// Struct instantiation tests
// ---------------------------------------------------------------------------

func TestAddressInstantiation(t *testing.T) {
	addr := generated.Address{
		Street:  "123 Main St",
		City:    "Springfield",
		Country: "US",
	}

	if addr.Street != "123 Main St" {
		t.Errorf("expected Street '123 Main St', got '%s'", addr.Street)
	}
	if addr.City != "Springfield" {
		t.Errorf("expected City 'Springfield', got '%s'", addr.City)
	}
	if addr.Country != "US" {
		t.Errorf("expected Country 'US', got '%s'", addr.Country)
	}
}

func TestPersonInstantiation(t *testing.T) {
	dob, err := time.Parse(time.RFC3339, "1990-06-15T00:00:00Z")
	if err != nil {
		t.Fatalf("failed to parse date: %v", err)
	}

	person := generated.Person{
		Email:     "alice@example.com",
		FirstName: "Alice",
		LastName:  "Smith",
		Dob:       dob,
		Address: generated.Address{
			Street:  "456 Elm Ave",
			City:    "Shelbyville",
			Country: "US",
		},
	}

	if person.Email != "alice@example.com" {
		t.Errorf("expected Email 'alice@example.com', got '%s'", person.Email)
	}
	if person.FirstName != "Alice" {
		t.Errorf("expected FirstName 'Alice', got '%s'", person.FirstName)
	}
}

func TestEmployeeInstantiation(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1985-03-22T00:00:00Z")

	emp := generated.Employee{
		Person: generated.Person{
			Email:     "bob@corp.com",
			FirstName: "Bob",
			LastName:  "Jones",
			Dob:       dob,
			Address: generated.Address{
				Street:  "789 Oak Rd",
				City:    "Capital City",
				Country: "US",
			},
		},
		EmployeeId: "EMP-001",
		Status:     generated.EmploymentStatus_FULL_TIME,
		Salary:     95000.00,
	}

	if emp.EmployeeId != "EMP-001" {
		t.Errorf("expected EmployeeId 'EMP-001', got '%s'", emp.EmployeeId)
	}
	if emp.Salary != 95000.00 {
		t.Errorf("expected Salary 95000.00, got %f", emp.Salary)
	}
}

// ---------------------------------------------------------------------------
// Enum value tests
// ---------------------------------------------------------------------------

func TestEmploymentStatusEnumValues(t *testing.T) {
	// Verify all expected enum constants are defined and distinct
	statuses := []generated.EmploymentStatus{
		generated.EmploymentStatus_FULL_TIME,
		generated.EmploymentStatus_PART_TIME,
		generated.EmploymentStatus_CONTRACT,
		generated.EmploymentStatus_UNEMPLOYED,
	}

	seen := make(map[generated.EmploymentStatus]bool)
	for _, s := range statuses {
		if seen[s] {
			t.Errorf("duplicate enum value detected: %v", s)
		}
		seen[s] = true
	}

	if len(seen) != 4 {
		t.Errorf("expected 4 distinct EmploymentStatus values, got %d", len(seen))
	}
}

// ---------------------------------------------------------------------------
// JSON round-trip tests
// ---------------------------------------------------------------------------

func TestPersonJSONRoundTrip(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1992-11-05T00:00:00Z")

	original := generated.Person{
		Email:     "carol@example.com",
		FirstName: "Carol",
		LastName:  "White",
		Dob:       dob,
		Address: generated.Address{
			Street:  "1 Infinite Loop",
			City:    "Cupertino",
			Country: "US",
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	// Unmarshal back
	var restored generated.Person
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	// Compare key fields
	if restored.Email != original.Email {
		t.Errorf("Email mismatch after round-trip: got '%s', want '%s'", restored.Email, original.Email)
	}
	if restored.FirstName != original.FirstName {
		t.Errorf("FirstName mismatch after round-trip: got '%s', want '%s'", restored.FirstName, original.FirstName)
	}
	if restored.Address.City != original.Address.City {
		t.Errorf("Address.City mismatch after round-trip: got '%s', want '%s'", restored.Address.City, original.Address.City)
	}
}

func TestEmployeeJSONRoundTrip(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1980-01-01T00:00:00Z")

	original := generated.Employee{
		Person: generated.Person{
			Email:     "dave@corp.com",
			FirstName: "Dave",
			LastName:  "Brown",
			Dob:       dob,
			Address:   generated.Address{Street: "10 Downing", City: "London", Country: "GB"},
		},
		EmployeeId: "EMP-999",
		Status:     generated.EmploymentStatus_CONTRACT,
		Salary:     120000.50,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	var restored generated.Employee
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	if restored.EmployeeId != original.EmployeeId {
		t.Errorf("EmployeeId mismatch: got '%s', want '%s'", restored.EmployeeId, original.EmployeeId)
	}
	if restored.Status != original.Status {
		t.Errorf("Status mismatch: got %v, want %v", restored.Status, original.Status)
	}
	if restored.Salary != original.Salary {
		t.Errorf("Salary mismatch: got %f, want %f", restored.Salary, original.Salary)
	}
}
