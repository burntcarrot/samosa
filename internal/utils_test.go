package internal

import "testing"

func TestCalculateCoverage(t *testing.T) {
	type info struct {
		covered int
		total   int
	}

	tests := []struct {
		name string
		info info
		want float64
	}{
		{"complete coverage", info{100, 100}, 100.0},
		{"partial coverage", info{79, 100}, 79.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateCoverage(tt.info.covered, tt.info.total)
			if got != tt.want {
				t.Errorf("got: %v, want: %v\n", got, tt.want)
			}
		})
	}
}
func Test_get_file_name(t *testing.T) {
	t.Run("modulepath", func(t *testing.T) {
		_, err := getFilename("github.com/web-alytics/meditate/pkg/logging/logging.go")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("localfilepath", func(t *testing.T) {
		_, err := getFilename("./testdata/test_coverage.out")
		if err != nil {
			t.Fatal(err)
		}
	})
}
