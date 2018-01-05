package main

import "testing"

func TestOddEven(t *testing.T) {
	if isEven(1) != false {
		t.Error("Expected: 1 is false, but got", isEven(1))
	}

	if isEven(2) != true {
		t.Error("Expected: 1 is true, but got", isEven(1))
	}
}
