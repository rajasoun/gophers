package model

import (
	"reflect"
	"testing"
)

func TestGetTip(t *testing.T) {
	//tips := Tips{}
	assertEquals := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Get Tip for valid Topic", func(t *testing.T) {
		got := GetTip("help")
		want := "Everyday Git in twenty commands or so : git help everyday"
		assertEquals(t, got, want)
	})
	t.Run("Get Tip for invalid Topic", func(t *testing.T) {
		got := GetTip("dummy")
		want := "Tips Not Available for Topic"
		assertEquals(t, got, want)
	})
	t.Run("Get Tip for valid Topic", func(t *testing.T) {
		got := GetTip("stash")
		want := "Saving current state of tracked files without commiting : git stash"
		assertEquals(t, got, want)
	})

}
func TestLoadTipsFromJson(t *testing.T) {
	_, got := LoadTipsFromJson()
	expected := MockReadJsonFile()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %q want %q", got, expected)
	}
}

func TestReadJsonFileNegative(t *testing.T) {
	path := "tips/data"
	_, got := readJsonFile(path)
	want := ErrInsufficient
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
