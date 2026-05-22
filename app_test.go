package main

import (
	"app/math"
	"testing"
)

// Test function names must start with "Test" followed by a capitalized letter
func TestSum(t *testing.T) {
	got := float64(math.Sum(2, 3))
	want := float64(5)

	if got != want {
		// Use t.Errorf to report a failure and continue, or t.Fatalf to stop immediately
		t.Errorf("Sum(2, 3) = %f; want %f", got, want)
	}
}
