package main

import (
	"testing"
)

func TestCalculateOverlap(t *testing.T) {
	overlap := CalculateOverlap(3, 5, 5)
	if overlap != 60 {
		t.Errorf("Overlap value was incorrect, got %f expected %f", overlap, 60.00)
	}
}
