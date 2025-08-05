package ccwc_test

import (
	"testing"

	"github.com/pradist/coding-challenges/ccwc"
)

func TestCountLines(t *testing.T) {
	// Test cases for counting lines in a file
	tests := []struct {
		name     string
		filePath string
		expected int
	}{
		{"Empty file", "testdata/empty.txt", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ccwc.CountLines(tt.filePath)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %d lines, got %d", tt.expected, result)
			}
		})
	}
}
