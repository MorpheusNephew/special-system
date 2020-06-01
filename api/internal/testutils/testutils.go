package testutils

import "testing"

// IfStringsNotEqual prints an error if the integers do not match
func IfStringsNotEqual(t *testing.T, received string, expected string) {
	if received != expected {
		t.Errorf("Expected %v, but received %v\n", expected, received)
	}
}

// IfStringIsEmpty prints an error if the string is nil
func IfStringIsEmpty(t *testing.T, received string) {
	if len(received) == 0 {
		t.Error("Expected string to not be empty")
	}
}

// IfIntsNotEqual prints an error if the integers do not match
func IfIntsNotEqual(t *testing.T, received int, expected int) {
	if received != expected {
		t.Errorf("Expected %v, but received %v\n", expected, received)
	}
}
