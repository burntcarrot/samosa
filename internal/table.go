package internal

import (
	"fmt"
	"log"

	"github.com/pterm/pterm"
)

const TRIM_LIMIT = 50

func trimString(str string, limit int) string {
	if len(str) > limit {
		return "..." + str[len(str)-limit:]
	}

	return str
}

func PrintTable(fi []*funcInfo, covered, total int, pkg bool) {
	table := make([][]string, len(fi)+1)

	table[0] = []string{"File", "Function", "Impact", "Uncovered Lines", "Start Line", "End Line"}

	totalCoverage := calculateCoverage(covered, total)

	for i, f := range fi {
		var fileName string
		if pkg {
			fileName = f.pkgFileName
		} else {
			fileName = f.fileName
		}

		impact := calculateCoverage((covered+f.uncoveredLines), total) - totalCoverage

		impactStr := formatImpact(impact)

		table[i+1] = []string{trimString(fileName, TRIM_LIMIT), f.functionName, impactStr, fmt.Sprint(f.uncoveredLines), fmt.Sprint(f.startLine), fmt.Sprint(f.endLine)}
	}

	err := pterm.DefaultTable.WithSeparator("\t").WithData(table).WithHasHeader(true).Render()
	if err != nil {
		log.Fatalf("failed to display results: %v\n", err)
	}
}
