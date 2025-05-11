package analyser

import (
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/utils.go"
)

func readAndValidateCSVFiles(file1, file2 string) ([][]string, [][]string, error) {
	records1, err := utils.ReadCSV(file1)
	if err != nil {
		return nil, nil, err
	}
	records2, err := utils.ReadCSV(file2)
	if err != nil {
		return nil, nil, err
	}

	if !utils.CompareRows(records1[0], records2[0]) {
		return nil, nil, fmt.Errorf("headers do not match between the files")
	}

	return records1, records2, nil
}

func indexRowsByKey(data [][]string, keyIndexes []int) map[string][]string {
	indexed := make(map[string][]string)
	for _, row := range data {
		key := utils.GenerateKey(row, keyIndexes)
		indexed[key] = row
	}
	return indexed
}
