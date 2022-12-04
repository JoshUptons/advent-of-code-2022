package main

import (
	"fmt"
	"testing"
)

func TestValue(t *testing.T) {
	tests := []struct {
		Name  string
		Want  int
		Input int
		Err   bool
	}{
		{"a", 1, int(rune('a')), false},
		{"b", 2, int(rune('b')), false},
		{"c", 3, int(rune('c')), false},
		{"d", 5, int(rune('d')), true},
		{"A", 27, int(rune('A')), false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := Value(tt.Input)
			if got == tt.Want == tt.Err {
				t.Fatalf("Failed, got %d, wanted: %d", got, tt.Want)
			}
			fmt.Println(got, tt.Want)
		})
	}
}
