package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareRows(t *testing.T) {
	tests := []struct {
		name     string
		row1     []string
		row2     []string
		expected bool
	}{
		{"Equal rows", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{"Different lengths", []string{"a", "b"}, []string{"a", "b", "c"}, false},
		{"Same content different order", []string{"1", "2", "3"}, []string{"3", "2", "1"}, false},
		{"Empty rows", []string{}, []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, CompareRows(tt.row1, tt.row2))
		})
	}
}

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name       string
		row        []string
		keyIndexes []int
		expected   string
	}{
		{"Basic key", []string{"a", "b", "c"}, []int{0, 2}, "a_c"},
		{"Single index", []string{"x", "y", "z"}, []int{1}, "y"},
		{"All indexes", []string{"1", "2"}, []int{0, 1}, "1_2"},
		{"Index out of range", []string{"only"}, []int{0, 1}, "only"},
		{"No indexes", []string{"a", "b"}, []int{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, GenerateKey(tt.row, tt.keyIndexes))
		})
	}
}
