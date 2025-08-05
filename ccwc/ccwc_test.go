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
		{"Single line file", "testdata/single_line.txt", 1},
		{"Multiple lines file", "testdata/multiline.txt", 3},
		{"File without trailing newline", "testdata/no_newline.txt", 1},
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

func TestCountLinesFileNotFound(t *testing.T) {
	// Test error case when file doesn't exist
	_, err := ccwc.CountLines("nonexistent_file.txt")
	if err == nil {
		t.Fatal("expected error for nonexistent file, got nil")
	}
}

func TestCountLinesLongLine(t *testing.T) {
	// Test with a very long line to trigger scanner error
	_, err := ccwc.CountLines("testdata/long_line.txt")
	if err == nil {
		t.Fatal("expected scanner error for very long line, got nil")
	}
}
