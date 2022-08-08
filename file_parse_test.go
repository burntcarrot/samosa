package samosa

import (
	"testing"
)

func TestParse(t *testing.T) {
	testOut := make(map[string]int)
	err := readCoverReport("./testdata/test_coverage.txt", testOut)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCoverage(t *testing.T) {
	inp := "25.31,27.2 4 1"
	want := 4
	got := getCoverage(inp)
	if got != want {
		t.Fatalf("got: %v, want:%v\n", got, want)
	}

}

func TestIncrement(t *testing.T) {
	inpMap := make(map[string]int)
	mapKey := t.Name()
	value := 1
	inpMap[mapKey] = value
	increment(inpMap, mapKey, value)
	if inpMap[mapKey] != 2 {
		t.Fatalf("expected:`2`,got:%v\n", inpMap[mapKey])
	}

}
