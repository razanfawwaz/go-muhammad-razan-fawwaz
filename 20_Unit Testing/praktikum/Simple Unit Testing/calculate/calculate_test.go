package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected 3")
	}
	if Addition(2, 2) != 4 {
		t.Error("Expected 4")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(1, 2) != -1 {
		t.Error("Expected -1")
	}
	if Substraction(2, 2) != 0 {
		t.Error("Expected 0")
	}
}

func TestDivision(t *testing.T) {
	if Division(4, 2) != 2 {
		t.Error("Expected 2")
	}
	if Division(2, 2) != 1 {
		t.Error("Expected 1")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(1, 2) != 2 {
		t.Error("Expected 2")
	}
	if Multiplication(2, 2) != 4 {
		t.Error("Expected 4")
	}
}
