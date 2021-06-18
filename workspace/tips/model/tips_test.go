package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTip(t *testing.T) {
	t.Run("Get Tip for valid Topic - rebase", func(t *testing.T) {
		got := GetTip("rebase")
		want := "Rebases 'feature' to 'master' and merges it in to master"
		assert.Contains(t, got, want)
	})
	t.Run("Get Tip for valid Topic - help", func(t *testing.T) {
		got := GetTip("help")
		want := "Everyday Git in twenty commands or so : git help everyday"
		assert.Equal(t, got, want)
	})
	t.Run("Get Tip for invalid Topic - dummy", func(t *testing.T) {
		got := GetTip("dummy")
		want := "Tips Not Available for Topic"
		assert.Equal(t, got, want)
	})
}

// func TestLoadTipsFromJson(t *testing.T) {
// 	_, got := LoadTipsFromJson()
// 	expected := MockReadJsonFile()

// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("got %q want %q", got, expected)
// 	}
// }

// func TestReadJsonFileNegative(t *testing.T) {
// 	path := "tips/data"
// 	_, got := readJsonFile(path)
// 	want := ErrInsufficient
// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}

// }
