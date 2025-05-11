package flagpkg

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

func (inF *InputFlag) registerFlags() {
	flag.StringVarP(&inF.saveFormat, "save", "s", "", "Format to save output (json or csv)")
	flag.StringVarP(&inF.outputFilename, "output", "o", "", "Filename to save changes")
	flag.StringVarP(&inF.keyIndexesFlag, "key", "k", "0", "Comma-separated list of column indexes for composite key")
	flag.BoolVarP(&inF.showChanges, "show-changes", "c", false, "Show detailed changes")
	flag.BoolVarP(&inF.includeUnchanged, "include-unchanged", "z", false, "Include unchanged rows in the output")
	flag.BoolVarP(&inF.filterAdded, "filter-added", "a", false, "Only show added rows")
	flag.BoolVarP(&inF.filterDeleted, "filter-deleted", "d", false, "Only show deleted rows")
	flag.BoolVarP(&inF.filterUpdated, "filter-updated", "u", false, "Only show updated rows")
	flag.BoolVarP(&inF.filterReordered, "filter-reordered", "r", false, "Only show reordered rows")
	flag.BoolVarP(&inF.help, "help", "h", false, "Show usage information")
	flag.BoolVarP(&inF.version, "version", "v", false, "Show app version")
}

func (inF *InputFlag) Parse(version string) bool {
	flag.Parse()

	if inF.version {
		fmt.Println("App Version:", version)
		os.Exit(0)
	}

	if inF.help {
		printUsage()
		return false
	}

	inF.files = flag.Args()
	if len(inF.files) != 2 {
		fmt.Println("Error: Two CSV file paths are required.")
		printUsage()
		return false
	}

	return true
}
