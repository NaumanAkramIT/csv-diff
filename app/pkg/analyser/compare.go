package analyser

import (
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/changespkg"
	"github.com/manishjalui11/csv-compare/app/pkg/summarypkg"
	"github.com/manishjalui11/csv-compare/app/pkg/utils.go"
)

func computeDifferences(data1, data2 [][]string, map1, map2 map[string][]string, filterTypes map[string]bool, includeUnchanged bool, keyIndexes []int) ([]*changespkg.Change, *summarypkg.Summary) {
	var changes []*changespkg.Change
	summary := summarypkg.New()

	for key, row2 := range map2 {
		if row1, exists := map1[key]; exists {
			if utils.CompareRows(row1, row2) {
				summary.Unmodified++
				if includeUnchanged {
					changes = append(changes, changespkg.New(changespkg.Unmodified, key, row1, row2, ""))
				}
			} else {
				summary.Updated++
				if filterTypes[changespkg.Updated] {
					changes = append(changes, changespkg.New(changespkg.Updated, key, row1, row2, ""))
				}
			}
		} else {
			summary.Added++
			if filterTypes[changespkg.Added] {
				changes = append(changes, changespkg.New(changespkg.Added, key, nil, row2, ""))
			}
		}
	}

	for key, row1 := range map1 {
		if _, exists := map2[key]; !exists {
			summary.Deleted++
			if filterTypes[changespkg.Deleted] {
				changes = append(changes, changespkg.New(changespkg.Deleted, key, row1, nil, ""))
			}
		}
	}

	if len(data1) == len(data2) {
		for i := range data1 {
			k1 := utils.GenerateKey(data1[i], keyIndexes)
			k2 := utils.GenerateKey(data2[i], keyIndexes)
			if k1 != k2 && utils.CompareRows(data1[i], map2[k1]) {
				summary.Reordered++
				if filterTypes[changespkg.Reordered] {
					changes = append(changes, changespkg.New(changespkg.Reordered, k1, data1[i], map2[k1], fmt.Sprintf("Originally at line %d", i+2)))
				}
			}
		}
	}

	return changes, summary
}

func CompareCSVDetailed(file1, file2, saveFormat, outputFilename string, showChanges, includeUnchanged bool, filterTypes map[string]bool, keyIndexes []int) error {
	records1, records2, err := readAndValidateCSVFiles(file1, file2)
	if err != nil {
		return err
	}

	data1, data2 := records1[1:], records2[1:]
	map1 := indexRowsByKey(data1, keyIndexes)
	map2 := indexRowsByKey(data2, keyIndexes)

	changes, summary := computeDifferences(data1, data2, map1, map2, filterTypes, includeUnchanged, keyIndexes)

	if err := handleOutput(changes, saveFormat, outputFilename, showChanges, summary); err != nil {
		return err
	}

	return nil

}
