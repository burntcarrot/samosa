package samosa

import (
	"testing"
)

func Test_parse(t *testing.T) {
	testOut := make(map[string]int)
	err := readCoverReport("./testdata/coverage.out", testOut)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_coverage(t *testing.T) {
	inp := "25.31,27.2 4 1"
        want := 4
	got := getCoverage(inp)
	if got != want {
		t.Fatalf("got: %v, want:%v\n", got, want)
	}

}

func Test_increment(t *testing.T) {
	inpMap := make(map[string]int)
	mapKey := t.Name()
	value := 1
	inpMap[mapKey] = value
	increment(inpMap, mapKey, value)
	if inpMap[mapKey] != 2 {
		t.Fatalf("expected:`2`,got:%v\n", inpMap[mapKey])
	}

}
