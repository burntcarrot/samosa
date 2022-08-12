package samosa

import (
	"regexp"
	"sort"

	funk "github.com/thoas/go-funk"
)

func FilterFunctionInfo(fi []funcInfo, filterOpts FilterOptions) ([]funcInfo, error) {
	var filteredFuncInfo []funcInfo
	var err error

	if filterOpts.Include != "" {
		filteredFuncInfo, err = filterByRegex(filterOpts.Include, fi)
		if err != nil {
			return nil, err
		}
	} else if filterOpts.Exclude != "" {
		filteredFuncInfo, err = filterByRegex(filterOpts.Exclude, fi)
		if err != nil {
			return nil, err
		}

		filteredFuncInfo = funk.Subtract(fi, filteredFuncInfo).([]funcInfo)
	} else {
		filteredFuncInfo = fi
	}

	if !filterOpts.SortFile {
		fi = sortFuncInfo(filteredFuncInfo)
	} else {
		fi = filteredFuncInfo
	}

	return fi, nil
}

func filterByRegex(pattern string, fi []funcInfo) ([]funcInfo, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	filteredFuncInfo := funk.Filter(fi, func(f funcInfo) bool {
		return r.Match([]byte(f.FileName))
	}).([]funcInfo)

	return filteredFuncInfo, nil
}

// sortFuncInfo returns function information sorted by the number of uncovered lines.
func sortFuncInfo(fi []funcInfo) []funcInfo {
	var filteredFuncInfos []funcInfo

	for _, f := range fi {
		if f.UncoveredLines > 0 {
			filteredFuncInfos = append(filteredFuncInfos, f)
		}
	}

	sort.Slice(filteredFuncInfos, func(i, j int) bool {
		return filteredFuncInfos[i].UncoveredLines > filteredFuncInfos[j].UncoveredLines
	})

	return filteredFuncInfos
}
