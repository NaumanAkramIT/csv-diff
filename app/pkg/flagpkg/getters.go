package flagpkg

import (
	"github.com/manishjalui11/csv-compare/app/pkg/changespkg"
	"strconv"
	"strings"
)

func (inF *InputFlag) GetFilePaths() (string, string) {
	return inF.files[0], inF.files[1]
}

func (inF *InputFlag) GetParsedKeyIndexes() []int {
	var indexes []int
	for _, s := range strings.Split(inF.keyIndexesFlag, ",") {
		index, err := strconv.Atoi(s)
		if err == nil {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

func (inF *InputFlag) GetFilterMap() map[string]bool {
	filterMap := map[string]bool{
		changespkg.Added:     true,
		changespkg.Deleted:   true,
		changespkg.Updated:   true,
		changespkg.Reordered: false,
	}
	if inF.filterAdded || inF.filterDeleted || inF.filterUpdated || inF.filterReordered {
		filterMap[changespkg.Added] = inF.filterAdded
		filterMap[changespkg.Deleted] = inF.filterDeleted
		filterMap[changespkg.Updated] = inF.filterUpdated
		filterMap[changespkg.Reordered] = inF.filterReordered
	}
	return filterMap
}

func (inF *InputFlag) GetSaveFormat() string {
	return inF.saveFormat
}
func (inF *InputFlag) GetOutputFilename() string {
	return inF.outputFilename
}
func (inF *InputFlag) ShowChangesEnabled() bool {
	return inF.showChanges
}
func (inF *InputFlag) IncludeUnchangedEnabled() bool {
	return inF.includeUnchanged
}
