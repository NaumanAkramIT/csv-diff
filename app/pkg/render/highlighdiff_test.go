package render

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighlightDiff(t *testing.T) {
	oldStr := "Hello bold world"
	newStr := "Hello brave world"

	expectedOld := "Hello b" + Red("old") + " world"
	expectedNew := "Hello b" + Green("rave") + " world"

	actualOld, actualNew := HighlightDiff(oldStr, newStr)

	assert.Equal(t, expectedOld, actualOld, "Old string diff highlighting should match")
	assert.Equal(t, expectedNew, actualNew, "New string diff highlighting should match")
}
