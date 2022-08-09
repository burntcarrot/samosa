package samosa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/cover"
)

func TestWalkModDir(t *testing.T) {
	type test struct {
		description string
		filename    string
		wantErr     bool
	}

	tests := []test{
		{
			description: "valid directory",
			filename:    "./testdata/test_walkmoddir/go.mod.txt",
		},
		{
			description: "invalid path",
			filename:    "",
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		_, err := walkModDir(tc.filename)
		if err != nil && !tc.wantErr {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}
	}
}

func TestWalkDir(t *testing.T) {
	type test struct {
		description string
		filename    string
		wantErr     bool
	}

	tests := []test{
		{
			description: "valid directory",
			filename:    "./testdata/test_walkmoddir/go.mod.txt",
		},
		{
			description: "invalid path",
			filename:    "",
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		_, err := walkDir()
		if err != nil && !tc.wantErr {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}
	}
}

func TestProfile(t *testing.T) {
	type test struct {
		description string
		filename    string
		want        []*cover.Profile
		wantErr     bool
	}

	tests := []test{
		{
			description: "normal coverage",
			filename:    "./testdata/test_profile/test_coverage.txt",
			want: []*cover.Profile{
				{
					FileName: "github.com/burntcarrot/samosa/cover.go",
					Mode:     "atomic",
					Blocks: []cover.ProfileBlock{
						{
							StartLine: 3,
							StartCol:  78,
							EndLine:   5,
							EndCol:    16,
							NumStmt:   2,
							Count:     0,
						},
						{
							StartLine: 9,
							StartCol:  2,
							EndLine:   10,
							EndCol:    16,
							NumStmt:   2,
							Count:     0,
						},
					},
				},
			},
		},
		{
			description: "empty coverage",
			filename:    "./testdata/test_profile/test_coverage_empty.txt",
			want:        []*cover.Profile{},
		},
		{
			description: "missing coverage",
			filename:    "./testdata/test_profile/test_coverage_missing.txt",
			want:        []*cover.Profile{},
		},
		{
			description: "invalid path",
			filename:    "",
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		got, err := getProfiles(tc.filename)
		if err != nil && !tc.wantErr {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}

		diff := cmp.Diff(got, tc.want)
		if len(diff) > 0 {
			t.Fatalf("got != want; got = %v, want = %v\n", got, tc.want)
		}
	}
}

func TestGetFunctionInfo(t *testing.T) {
	type test struct {
		description string
		profiles    []*cover.Profile
		wantErr     bool
	}

	tests := []test{
		{
			description: "valid profile",
			profiles: []*cover.Profile{
				{
					FileName: "github.com/burntcarrot/samosa/cover.go",
					Mode:     "atomic",
					Blocks: []cover.ProfileBlock{
						{
							StartLine: 3,
							StartCol:  78,
							EndLine:   5,
							EndCol:    16,
							NumStmt:   2,
							Count:     0,
						},
						{
							StartLine: 9,
							StartCol:  2,
							EndLine:   10,
							EndCol:    16,
							NumStmt:   2,
							Count:     0,
						},
					},
				},
			},
		},
		{
			description: "empty profile",
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		_, _, _, err := getFunctionInfo(tc.profiles)
		if err != nil && !tc.wantErr {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}
	}
}
