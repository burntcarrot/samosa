package internal

import (
	"fmt"
	"go/build"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/thoas/go-funk"
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

func filterByRegex(pattern string, fi []*funcInfo) ([]*funcInfo, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	filteredFuncInfo := funk.Filter(fi, func(f *funcInfo) bool {
		return r.Match([]byte(f.fileName))
	}).([]*funcInfo)

	return filteredFuncInfo, nil
}

func getFilename(filePath string) (string, error) {
	dir, file := filepath.Split(filePath)
	var pkg *build.Package
	var err error
	fext := strings.Split(file, ".")[1]
	if strings.EqualFold(fext, "go") {
		pkg, err = build.Import(strings.TrimSuffix(dir, "/"), ".", build.FindOnly)
		if err != nil {
			if strings.Contains(err.Error(), "go get") {
				if err := importModule(strings.TrimSuffix(dir, "/")); err != nil {
					return "", err
				}
			}
			return "", err
		}

	} else {
		pkg, err = build.ImportDir(dir, build.FindOnly)
		if err != nil {
			pkg, err = build.Import(dir, ".", build.FindOnly)
			if err != nil {
				if strings.Contains(err.Error(), "add") {
					if err := importModule(strings.TrimSuffix(dir, "/")); err != nil {
						return "", err
					}
					return "", err
				}

			}
			return "", err
		}
	}
	return filepath.Join(pkg.Dir, file), nil
}

func importModule(s string) error {
	cmd := exec.Command("go")
	cmd.Args = append(cmd.Args, "get", fmt.Sprintf("%s@latest", s))
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
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
