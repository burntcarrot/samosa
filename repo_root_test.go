package samosa

import (
	"errors"
	"testing"
)

func TestRoot(t *testing.T) {
	got, err := getRoot()
	if err != nil {
		t.Fatalf("test failed: %v\n", err.Error())
	}

	if len(got) < 1 {
		t.Fatalf("test failed: %v\n", errors.New("non empty data expected"))
	}
}

func TestDecode(t *testing.T) {
	got := map[string]interface{}{}
	data, err := getRoot()
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}

	err = decodeJSON(data, got)
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}

	if len(got) < 1 {
		t.Fatal("expected to decode all values got empty")
	}
}

func TestGetMod(t *testing.T) {
	got, err := getMod()
	if err != nil {
		t.Fatalf("no err expected:%v\n", err)
	}

	if len(got) < 1 {
		t.Fatal("expected to get complete go.mod path for the repo")
	}
}
