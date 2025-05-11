package utils

import (
	"fmt"
	"strings"
)

func RowToString(row []string) string {
	if row == nil {
		return ""
	}
	return fmt.Sprintf("%v", row)
}

func CompareRows(row1, row2 []string) bool {
	return RowToString(row1) == RowToString(row2)
}

func GenerateKey(row []string, keyIndexes []int) string {
	var keyParts []string
	for _, index := range keyIndexes {
		if index < len(row) {
			keyParts = append(keyParts, row[index])
		}
	}
	return strings.Join(keyParts, "_")
}
