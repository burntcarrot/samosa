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
	slice, err := walkModDir(want)
	t.Logf("dir list:%v\n", slice)
	t.Logf("err:%v\n", err)

}

func TestProfile(t *testing.T) {
	_, err := getProfiles("./coverage.out")
	if err != nil {
		t.Log(err)
	}

}

func TestFile_names(t *testing.T) {
	profiles, err := getProfiles("./coverage.out")
	if err != nil {
		t.Log(err)
	}
	_, _, _, err = getFunctionInfo(profiles)
	if err != nil {
		t.Fatal("no error expected")
	}
}

func TestWalk_file(t *testing.T) {
	_, err := walkDir()
	if err != nil {
		t.Fatal("no error expected")
	}
}
