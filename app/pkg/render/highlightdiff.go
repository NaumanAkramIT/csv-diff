package render

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"strings"
)

func HighlightDiff(oldStr, newStr string) (string, string) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(oldStr, newStr, false)

	var oldResult, newResult strings.Builder

	for _, d := range diffs {
		switch d.Type {
		case diffmatchpatch.DiffEqual:
			oldResult.WriteString(d.Text)
			newResult.WriteString(d.Text)
		case diffmatchpatch.DiffDelete:
			oldResult.WriteString(Red(d.Text))
		case diffmatchpatch.DiffInsert:
			newResult.WriteString(Green(d.Text))
		}
	}
	return oldResult.String(), newResult.String()
}
