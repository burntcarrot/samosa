package samosa

import "testing"

func Test_walk(t *testing.T) {
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

func Test_profile(t *testing.T) {
	_, err := getProfiles("./coverage.out")
	if err != nil {
		t.Log(err)
	}

}

func Test_file_names(t *testing.T) {
	profiles, err := getProfiles("./coverage.out")
	if err != nil {
		t.Log(err)
	}
	_, _, _, err = getFunctionInfo(profiles)
	if err != nil {
		t.Fatal("no error expected")
	}
}

func Test_walk_file(t *testing.T) {
	_, err := walkDir()
	if err != nil {
		t.Fatal("no error expected")
	}
}
