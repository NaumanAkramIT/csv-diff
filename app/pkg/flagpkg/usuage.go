package flagpkg

import (
	"fmt"
	"github.com/manishjalui11/csv-compare/app/pkg/render"
	flag "github.com/spf13/pflag"
)

func printUsage() {
	fmt.Println(render.Bold(render.Underline("\nCsv-Diff Command Usage:")))
	fmt.Println("  csv-diff [options] <file1.csv> <file2.csv>")
	fmt.Println(render.Underline(render.Magenta("\nOptions:")))
	flag.PrintDefaults()
	fmt.Println(render.Underline(render.Cyan("\nExample:")))
	fmt.Println("  csv-diff -c -s json -o changes.json file1.csv file2.csv")
}
