package test_1_0_0

// Test file lives in the same package as the generated code.
// No imports needed for the generated types — they are used directly.

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAddressInstantiation(t *testing.T) {
	addr := Address{
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
	person := Person{
		Email:     "alice@example.com",
		FirstName: "Alice",
		LastName:  "Smith",
		Dob:       dob,
		Address: Address{
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
	emp := Employee{
		Person: Person{
			Email:     "bob@corp.com",
			FirstName: "Bob",
			LastName:  "Jones",
			Dob:       dob,
			Address:   Address{Street: "789 Oak Rd", City: "Capital City", Country: "US"},
		},
		EmployeeId: "EMP-001",
		Status:     EmploymentStatus_FULL_TIME,
		Salary:     95000.00,
	}
	if emp.EmployeeId != "EMP-001" {
		t.Errorf("expected EmployeeId 'EMP-001', got '%s'", emp.EmployeeId)
	}
	if emp.Salary != 95000.00 {
		t.Errorf("expected Salary 95000.00, got %f", emp.Salary)
	}
}

func TestEmploymentStatusEnumValues(t *testing.T) {
	statuses := []EmploymentStatus{
		EmploymentStatus_FULL_TIME,
		EmploymentStatus_PART_TIME,
		EmploymentStatus_CONTRACT,
		EmploymentStatus_UNEMPLOYED,
	}
	seen := make(map[EmploymentStatus]bool)
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

func TestPersonJSONRoundTrip(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1992-11-05T00:00:00Z")
	original := Person{
		Email:     "carol@example.com",
		FirstName: "Carol",
		LastName:  "White",
		Dob:       dob,
		Address:   Address{Street: "1 Infinite Loop", City: "Cupertino", Country: "US"},
	}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	var restored Person
	if err := json.Unmarshal(data, &restored); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if restored.Email != original.Email {
		t.Errorf("Email mismatch: got '%s', want '%s'", restored.Email, original.Email)
	}
	if restored.Address.City != original.Address.City {
		t.Errorf("Address.City mismatch: got '%s', want '%s'", restored.Address.City, original.Address.City)
	}
}

func TestEmployeeJSONRoundTrip(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1980-01-01T00:00:00Z")
	original := Employee{
		Person: Person{
			Email:     "dave@corp.com",
			FirstName: "Dave",
			LastName:  "Brown",
			Dob:       dob,
			Address:   Address{Street: "10 Downing", City: "London", Country: "GB"},
		},
		EmployeeId: "EMP-999",
		Status:     EmploymentStatus_CONTRACT,
		Salary:     120000.50,
	}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	var restored Employee
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
