package main

import "testing"

func TestPermutationFirst(t *testing.T) {
	code, error := GetPermutationByIndex([]byte("ABCDE"), 3, 1)
	if error != nil {
		t.Errorf("Unexpected error: %s", error)
	}

	if code != "ABC" {
		t.Errorf("Expected 'ABC', got '%s'", code)
	}
}

func TestPermutationMultipleCharacters(t *testing.T) {
	code, error := GetPermutationByIndex([]byte("ABCDE"), 3, 20)
	if error != nil {
		t.Errorf("Unexpected error: %s", error)
	}

	if code != "BDC" {
		t.Errorf("Expected 'BDC', got '%s'", code)
	}
}

func TestPermutationLast(t *testing.T) {
	code, error := GetPermutationByIndex([]byte("ABCDE"), 3, 60)
	if error != nil {
		t.Errorf("Unexpected error: %s", error)
	}

	if code != "EDC" {
		t.Errorf("Expected 'EDC', got '%s'", code)
	}
}

func TestPermutationLong(t *testing.T) {
	code, error := GetPermutationByIndex([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), 10, 1)
	if error != nil {
		t.Errorf("Unexpected error: %s", error)
	}

	if code != "ABCDEFGHIJ" {
		t.Errorf("Expected 'ABCDEFGHIJ', got '%s'", code)
	}
}

func TestPermutationTooLow(t *testing.T) {
	_, error := GetPermutationByIndex([]byte("ABCDE"), 3, 0)
	if error == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPermutationTooHigh(t *testing.T) {
	_, error := GetPermutationByIndex([]byte("ABCDE"), 3, 61)
	if error == nil {
		t.Errorf("Expected error, got nil")
	}
}
