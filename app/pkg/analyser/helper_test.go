package analyser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAndValidateCSVFiles(t *testing.T) {
	content1 := "id,name\n1,Alice\n2,Bob\n"
	content2 := "id,name\n3,Charlie\n4,Dave\n"

	file1 := writeTempCSV(t, content1)
	defer os.Remove(file1)
	file2 := writeTempCSV(t, content2)
	defer os.Remove(file2)

	records1, records2, err := readAndValidateCSVFiles(file1, file2)

	assert.NoError(t, err)
	assert.NotNil(t, records1)
	assert.NotNil(t, records2)
	assert.Len(t, records1, 3)
	assert.Len(t, records2, 3)
}

func TestReadAndValidateCSVFiles_HeaderMismatch(t *testing.T) {
	content1 := "id,name\n1,Alice\n"
	content2 := "uid,username\n3,Charlie\n"

	file1 := writeTempCSV(t, content1)
	defer os.Remove(file1)
	file2 := writeTempCSV(t, content2)
	defer os.Remove(file2)

	_, _, err := readAndValidateCSVFiles(file1, file2)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "headers do not match between the files")
}

func TestIndexRowsByKey(t *testing.T) {
	data := [][]string{
		{"id", "name"},
		{"1", "Alice"},
		{"2", "Bob"},
	}
	keyIndexes := []int{0} // Use "id" as key

	indexed := indexRowsByKey(data[1:], keyIndexes) // skip header

	assert.Len(t, indexed, 2)
	assert.Equal(t, []string{"1", "Alice"}, indexed["1"])
	assert.Equal(t, []string{"2", "Bob"}, indexed["2"])
}
