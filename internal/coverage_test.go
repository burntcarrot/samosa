package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/cover"
)

func TestGetProfiles(t *testing.T) {
	t.Run("must return profiles when valid path is specified", func(t *testing.T) {
		got, _ := getProfiles("./testdata/test_coverage.out")
		want, _ := cover.ParseProfiles("./testdata/test_coverage.out")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v\n", got, want)
		}
	})

	t.Run("must return error when invalid path is specified", func(t *testing.T) {
		_, err := getProfiles("./test/test_coverage.out")
		assert.NotNil(t, err)
	})
}

func TestGetFunctionInfo(t *testing.T) {
	// t.Run("must return function info for valid profiles", func(t *testing.T) {
	// profiles, _ := cover.ParseProfiles("./testdata/test_coverage.out")
	// funcInfo, covered, total, err := getFunctionInfo(profiles)
	// assert.Nil(t, err)
	// assert.NotNil(t, funcInfo)
	// assert.Zero(t, covered) // the test data has 0 covered lines
	// assert.NotZero(t, total)
	// })
}

func TestGetCoverageData(t *testing.T) {
	t.Run("must display results without returning error", func(t *testing.T) {
		_, _, _, err := GetCoverageData("./testdata/test_coverage.out")
		assert.Nil(t, err)
	})

	t.Run("must return error when invalid path is specified", func(t *testing.T) {
		_, _, _, err := GetCoverageData("./test/test_coverage.out")
		assert.NotNil(t, err)
	})
}

func TestGetFileFunctions(t *testing.T) {
	got, err := getProfiles("./testdata/test_coverage.out")
	if err != nil {
		t.Fatal(err)
	}
	fname, start, end, err := getFunctionInfo(got)
	t.Logf("%v\n", start)
	t.Logf("%v\n", end)
	t.Logf("%v\n", err)
	for _, f := range fname {
		t.Logf("%v\n", f.functionName)
	}

}
