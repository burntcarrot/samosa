package samosa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/cover"
)

func TestWalk(t *testing.T) {
	got := map[string]interface{}{}
	data, err := getRoot()
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}
	decodeJSON(data, got)
	want := getModDir(got)
	_, err = walkModDir(want)
}

func TestProfile(t *testing.T) {
	type test struct {
		description string
		filename    string
		want        []*cover.Profile
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
	}

	for _, tc := range tests {
		got, err := getProfiles(tc.filename)
		if err != nil {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}

		diff := cmp.Diff(got, tc.want)
		if len(diff) > 0 {
			t.Fatalf("got != want; got = %v, want = %v\n", got, tc.want)
		}
	}
}

func TestFilenames(t *testing.T) {
	profiles, err := getProfiles("./testdata/test_coverage.txt")
	if err != nil {
		t.Log(err)
	}
	_, _, _, err = getFunctionInfo(profiles)
	if err != nil {
		t.Fatal("no error expected")
	}
}

func TestWalkFile(t *testing.T) {
	_, err := walkDir()
	if err != nil {
		t.Fatal("no error expected")
	}
}
