package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := addition(10, 5)
	expected := 15.0

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	result := subtraction(10, 5)
	expected := 5.0

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestMultiplication(t *testing.T) {
	result := multiplication(10, 5)
	expected := 50.0

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestDivision(t *testing.T) {
	result := division(10, 5)
	expected := 2.0

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
