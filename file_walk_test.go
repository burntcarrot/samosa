package samosa

import "testing"

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
	_, err := getProfiles("./testdata/test_coverage.txt")
	if err != nil {
		t.Log(err)
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
