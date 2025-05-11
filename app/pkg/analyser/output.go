package analyser

import (
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/changespkg"
	"github.com/manishjalui11/csv-compare/app/pkg/render"
	"github.com/manishjalui11/csv-compare/app/pkg/summarypkg"
)

func handleOutput(changes []*changespkg.Change, saveFormat, outputFilename string, showChanges bool, summary *summarypkg.Summary) error {
	render.PrintSummaryTable(summary)
	if showChanges {
		render.PrintDiffSummaryTable(changes)
	}

	if saveFormat != "" && outputFilename != "" {
		switch saveFormat {
		case JSON:
			return WriteJSON(changes, outputFilename)
		case CSV:
			return WriteCSV(changes, outputFilename)
		default:
			fmt.Println("Unsupported save format:", saveFormat)
		}
	}

	return nil
}
