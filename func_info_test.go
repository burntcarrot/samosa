package samosa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilterFunctionInfo(t *testing.T) {
	type test struct {
		description   string
		fi            []funcInfo
		filterOptions FilterOptions
		want          []funcInfo
		wantErr       bool
	}

	tests := []test{
		{
			description: "no filter options",
			fi: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
			},
			filterOptions: FilterOptions{},
			want: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
			},
		},
		{
			description: "sort file",
			fi: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
				{
					FileName:       "repo_root.go",
					PkgFileName:    "github.com/burntcarrot/samosa/repo_root.go",
					FunctionName:   "getRoot",
					StartLine:      10,
					EndLine:        20,
					UncoveredLines: 2,
				},
			},
			filterOptions: FilterOptions{
				SortFile: true,
			},
			want: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
				{
					FileName:       "repo_root.go",
					PkgFileName:    "github.com/burntcarrot/samosa/repo_root.go",
					FunctionName:   "getRoot",
					StartLine:      10,
					EndLine:        20,
					UncoveredLines: 2,
				},
			},
		},
		{
			description: "include files",
			fi: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
				{
					FileName:       "repo_root.go",
					PkgFileName:    "github.com/burntcarrot/samosa/repo_root.go",
					FunctionName:   "getRoot",
					StartLine:      10,
					EndLine:        20,
					UncoveredLines: 2,
				},
			},
			filterOptions: FilterOptions{
				Include: "cover*",
			},
			want: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
			},
		},
		{
			description: "exclude files",
			fi: []funcInfo{
				{
					FileName:       "cover.go",
					PkgFileName:    "github.com/burntcarrot/samosa/cover.go",
					FunctionName:   "GetCoverageData",
					StartLine:      4,
					EndLine:        16,
					UncoveredLines: 1,
				},
				{
					FileName:       "repo_root.go",
					PkgFileName:    "github.com/burntcarrot/samosa/repo_root.go",
					FunctionName:   "getRoot",
					StartLine:      10,
					EndLine:        20,
					UncoveredLines: 2,
				},
			},
			filterOptions: FilterOptions{
				Exclude: "cover*",
			},
			want: []funcInfo{
				{
					FileName:       "repo_root.go",
					PkgFileName:    "github.com/burntcarrot/samosa/repo_root.go",
					FunctionName:   "getRoot",
					StartLine:      10,
					EndLine:        20,
					UncoveredLines: 2,
				},
			},
		},
	}

	for _, tc := range tests {
		got, err := FilterFunctionInfo(tc.fi, tc.filterOptions)
		if err != nil && !tc.wantErr {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}

		diff := cmp.Diff(got, tc.want)
		if len(diff) > 0 {
			t.Fatalf("test failed (%s): got != want, diff: %v\n", tc.description, diff)
		}
	}
}
