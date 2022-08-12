package samosa

import (
	"testing"
)

func TestGetCoverageData(t *testing.T) {
	type test struct {
		description string
		filename    string
		wantErr     bool
	}

	tests := []test{
		{
			description: "valid coverage report",
			filename:    "./testdata/test_getcoveragedata/test_coverage.txt",
		},
		{
			description: "invalid path",
			filename:    "",
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		_, _, _, err := GetCoverageData(tc.filename)
		if err != nil && !tc.wantErr {
			t.Fatalf("test failed (%s): %v\n", tc.description, err.Error())
		}
	}
}
