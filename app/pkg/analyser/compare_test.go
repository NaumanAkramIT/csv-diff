package analyser

import (
	"encoding/json"
	"github.com/manishjalui11/csv-compare/app/pkg/changespkg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeTempCSV(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "*.csv")
	assert.NoError(t, err)
	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	tmpFile.Close()
	return tmpFile.Name()
}

func TestCompareCSVDetailed_WithOutputCheck(t *testing.T) {
	csv1 := `ID,Name,Age
1,Alice,30
2,Bob,25
3,Charlie,35`

	csv2 := `ID,Name,Age
1,Alice,30
2,Bob,26
4,David,40`

	file1 := writeTempCSV(t, csv1)
	file2 := writeTempCSV(t, csv2)
	defer os.Remove(file1)
	defer os.Remove(file2)

	outputFile := "test_output.json"
	defer os.Remove(outputFile)

	filterTypes := map[string]bool{
		changespkg.Added:      true,
		changespkg.Deleted:    true,
		changespkg.Updated:    true,
		changespkg.Unmodified: true,
	}
	keyIndexes := []int{0}

	err := CompareCSVDetailed(file1, file2, "json", outputFile, true, true, filterTypes, keyIndexes)
	assert.NoError(t, err)

	outputData, err := os.ReadFile(outputFile)
	assert.NoError(t, err)

	var changes []changespkg.Change
	err = json.Unmarshal(outputData, &changes)
	assert.NoError(t, err)

	assert.Equal(t, 4, len(changes))

	var added, deleted, updated, unmodified int
	for _, change := range changes {
		switch change.ChangeType {
		case changespkg.Added:
			added++
			assert.Equal(t, "4", change.Key)
		case changespkg.Deleted:
			deleted++
			assert.Equal(t, "3", change.Key)
		case changespkg.Updated:
			updated++
			assert.Equal(t, "2", change.Key)
			assert.Equal(t, "25", change.OldValue[2])
			assert.Equal(t, "26", change.NewValue[2])
		case changespkg.Unmodified:
			unmodified++
			assert.Equal(t, "1", change.Key)
		}
	}

	assert.Equal(t, 1, added)
	assert.Equal(t, 1, deleted)
	assert.Equal(t, 1, updated)
	assert.Equal(t, 1, unmodified)
}
