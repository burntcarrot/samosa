package samosa

import "testing"

func Test_root(t *testing.T) {
	data, err := getRoot()
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}
	if len(data) < 1 {
		t.Fatalf("non empty data is expected")
	}

}

func Test_decode(t *testing.T) {
	got := map[string]interface{}{}
	data, err := getRoot()
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}
	decodeJSON(data, got)
	if len(got) < 1 {
		t.Fatal("expected to decode all values got empty")
	}
}

func Test_mod_dir(t *testing.T) {
	got := map[string]interface{}{}
	data, err := getRoot()
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}
	decodeJSON(data, got)
	want := getModDir(got)
	if len(want) < 1 {
		t.Fatal("expected to get complete go.mod path for the repo")
	}

}
