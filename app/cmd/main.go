package main

import (
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/analyser"
	"github.com/manishjalui11/csv-compare/app/pkg/flagpkg"
	"github.com/manishjalui11/csv-compare/app/pkg/spinner"
	"os"
	"os/exec"
)

var version = "1.0.0"

func main() {
	version = getGitTag()

	inputFlags := flagpkg.New()
	if ok := inputFlags.Parse(version); !ok {
		os.Exit(1)
	}

	spinnerObj := spinner.NewSpinner()
	defer spinnerObj.Stop()

	file1, file2 := inputFlags.GetFilePaths()
	err := analyser.CompareCSVDetailed(file1, file2, inputFlags.GetSaveFormat(), inputFlags.GetOutputFilename(), inputFlags.ShowChangesEnabled(), inputFlags.IncludeUnchangedEnabled(), inputFlags.GetFilterMap(), inputFlags.GetParsedKeyIndexes())
	if err != nil {
		fmt.Println("Error comparing CSVs:", err)
		return
	}
}

func getGitTag() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return version
	}
	return string(output)
}
