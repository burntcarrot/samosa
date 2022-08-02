package samosa

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/pterm/pterm"
)

const TRIM_LIMIT = 50

func trimString(str string, limit int) string {
	if len(str) > limit {
		return "..." + str[len(str)-limit:]
	}

	return str
}

// calculateCoverage returns coverage in float64.
func calculateCoverage(covered, total int) float64 {
	return float64(covered) / float64(total) * 100
}

// formatImpact returns a pretty-printed string for the impact value.
func formatImpact(impact float64) string {
	var impactStr string
	if impact > 5 {
		impactStr = color.RedString("%.2f", impact)
	} else if impact > 2 && impact < 5 {
		impactStr = color.YellowString("%.2f", impact)
	} else {
		impactStr = color.GreenString("%.2f", impact)
	}

	return impactStr
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
