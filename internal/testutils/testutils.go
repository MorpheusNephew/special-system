package testutils

import "testing"

// IfStringsNotEqual prints an error if the integers do not match
func IfStringsNotEqual(t *testing.T, received string, expected string) {
	if received != expected {
		t.Errorf("Expected %v, but received %v\n", expected, received)
	}
}

// IfIntsNotEqual prints an error if the integers do not match
func IfIntsNotEqual(t *testing.T, received int, expected int) {
	if received != expected {
		t.Errorf("Expected %v, but received %v\n", expected, received)
	}
}
