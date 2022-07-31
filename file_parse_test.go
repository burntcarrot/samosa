package samosa

import (
	"testing"
)

func Test_parse(t *testing.T) {
	testOut := make(map[string]int)
	err := readCoverReport("./coverage.out", testOut)
	if err != nil {
		t.Fatal(err)
	}
	if len(testOut)<1{
		t.Fatal("expected to get a dataset with coverage")
	}
}

func Test_coverage(t *testing.T) {
	inp := "25.31,27.2 4 1"
	out := getCoverage(inp)
	if out != 4 {
		t.Fatalf("expected to get `4`got:%v", out)
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
