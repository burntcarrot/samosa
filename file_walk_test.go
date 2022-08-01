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
	profiles, err := getProfiles("./coverage.out")
	if err != nil {
		t.Log(err)
	}
	if len(profiles) < 1 {
		t.Fatal("expected to get valid profiles")
	}

}

func Test_file_names(t *testing.T) {
	profiles, err := getProfiles("./coverage.out")
	if err != nil {
		t.Log(err)
	}
	for _, profile := range profiles {
		t.Logf("file name:%v\n", profile.FileName)
	}
	finfo, st, end, err := getFunctionInfo(profiles)
	t.Log("finfo:", finfo)
	t.Log("start:", st)
	t.Log("end:", end)
	t.Log(err)

}

func Test_walk_file(t *testing.T) {
	got := walkDir()
	t.Logf("%v\n", got)

}
