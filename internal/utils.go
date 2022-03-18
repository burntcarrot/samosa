package internal

import (
	"errors"
	"go/build"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
)

// sortFuncInfo returns function information sorted by the number of uncovered lines.
func sortFuncInfo(fi []*funcInfo) []*funcInfo {
	var filteredFuncInfos []*funcInfo

	for _, f := range fi {
		if f.uncoveredLines > 0 {
			filteredFuncInfos = append(filteredFuncInfos, f)
		}
	}

	sort.Slice(filteredFuncInfos, func(i, j int) bool {
		return filteredFuncInfos[i].uncoveredLines > filteredFuncInfos[j].uncoveredLines
	})

	return filteredFuncInfos
}

func getFilename(filePath string) (string, error) {
	dir, file := filepath.Split(filePath)
	pkg, err := build.Import(dir, ".", build.FindOnly)
	if err != nil {
		return "", errors.New("can't find file")
	}

	return filepath.Join(pkg.Dir, file), nil
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
