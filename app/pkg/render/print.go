package render

import (
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/changespkg"
	"github.com/manishjalui11/csv-compare/app/pkg/summarypkg"
	"github.com/manishjalui11/csv-compare/app/pkg/utils.go"
)

func PrintSummaryTable(summary *summarypkg.Summary) {
	fmt.Println("\n" + Bold("====== Summary of Changes ======"))
	fmt.Printf("| %-15s | %-10s |\n", "Change Type", "Count")
	fmt.Println("|-----------------|------------|")
	fmt.Printf("| %-15s | %-10d |\n", Green("Added rows     "), summary.Added)
	fmt.Printf("| %-15s | %-10d |\n", Red("Deleted rows   "), summary.Deleted)
	fmt.Printf("| %-15s | %-10d |\n", Yellow("Updated rows   "), summary.Updated)
	fmt.Printf("| %-15s | %-10d |\n", "Unmodified rows", summary.Unmodified)
	fmt.Printf("| %-15s | %-10d |\n", Blue("Reordered rows "), summary.Reordered)
	fmt.Println("================================")
}

func PrintDiffSummaryTable(changes []*changespkg.Change) {
	if len(changes) == 0 {
		fmt.Println("No changes detected.")
		return
	}

	fmt.Println("\n" + Bold("=== Detailed Changes ==="))
	fmt.Printf("| %-15s | %-20s | %-10s | %-10s |\n", "Change ChangeType", "Key", "OldValue Value", "New Value")
	fmt.Println("|------------------|----------------------|------------|------------|")

	for _, change := range changes {
		var oldValue, newValue string

		if len(change.OldValue) > 0 {
			oldValue = utils.RowToString(change.OldValue)
		} else {
			oldValue = "-"
		}

		if len(change.NewValue) > 0 {
			newValue = utils.RowToString(change.NewValue)
		} else {
			newValue = "-"
		}

		switch change.ChangeType {
		case changespkg.Added:
			fmt.Printf("| %-15s | %-20s | %-10s | %-10s |\n", Green("✅ Added"), change.Key, oldValue, newValue)
		case changespkg.Deleted:
			fmt.Printf("| %-15s | %-20s | %-10s | %-10s |\n", Red("❌ Deleted"), change.Key, oldValue, newValue)
		case changespkg.Updated:
			coloredOld, coloredNew := HighlightDiff(oldValue, newValue)
			fmt.Printf("| %-15s | %-20s | %-10s | %-10s |\n", Yellow("🔁 Updated"), change.Key, coloredOld, coloredNew)
		case changespkg.Unmodified:
			fmt.Printf("| %-15s | %-20s | %-10s | %-10s |\n", Magenta("✔️ Unmodified"), change.Key, oldValue, newValue)
		case changespkg.Reordered:
			fmt.Printf("| %-15s | %-20s | %-10s | %-10s |\n", Blue("🔀 Reordered"), change.Key, oldValue, newValue)
		}
	}

	fmt.Println("=============================")
}
