package utils

import (
	"testing"
)

func TestGetCommonStocks(t *testing.T) {
	stocks := GetCommonStocks([]string{"A", "B"}, []string{"A"})
	if len(stocks) != 1 {
		t.Errorf("Overlap value was incorrect, got %d expected %d", len(stocks), 1)
	}
}

func TestRemoveDups(t *testing.T) {
	dups := RemoveDups([]string{"A", "A"})
	if len(dups) != 1 {
		t.Errorf("Overlap value was incorrect, got %d expected %d", len(dups), 1)
	}
}
