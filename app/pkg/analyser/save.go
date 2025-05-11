package analyser

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/changespkg"
	"github.com/manishjalui11/csv-compare/app/pkg/render"
	"github.com/manishjalui11/csv-compare/app/pkg/utils.go"
	"os"
)

func WriteJSON(changes []*changespkg.Change, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	fmt.Println(render.Green("✅ Changes saved to"), filename)
	return enc.Encode(changes)
}

func WriteCSV(changes []*changespkg.Change, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	err = writer.Write([]string{"ChangeType", "Key", "OldValue", "New", "ChangeNote"})
	if err != nil {
		return err
	}

	for _, change := range changes {
		err := writer.Write([]string{
			change.ChangeType,
			change.Key,
			utils.RowToString(change.OldValue),
			utils.RowToString(change.NewValue),
			change.ChangeNote,
		})
		if err != nil {
			return err
		}
	}

	fmt.Println(render.Green("✅ Changes saved to"), filename)
	return nil
}
