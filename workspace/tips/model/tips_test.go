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
	t.Run("Get Tip for invalid Topic - dummy", func(t *testing.T) {
		got := GetTip("dummy")
		want := "Tips Not Available for Topic"
		assert.Equal(t, got, want)
	})
}

func TestLoadTipsFromJson(t *testing.T) {
	t.Run("Load Tips From Json File and check if there are 166 tips ", func(t *testing.T) {
		got, _ := LoadTipsFromJson()
		expected := 166
		assert.Equal(t, len(got), expected)
	})
}

func TestReadJsonFile(t *testing.T) {
	t.Run("Load Json File and check if it contains the tip starting with Everyday ", func(t *testing.T) {
		got, _ := ReadJsonFile("../data/tips.json")
		expected := "Everyday Git in twenty commands or so"
		assert.Contains(t, string(got), expected)
	})
	t.Run("Loading invalid Json File should fail ", func(t *testing.T) {
		_, got := ReadJsonFile("../data1/tips.json")
		assert.Error(t, got)
	})
	t.Run("Load Json File and check if it not contains for DUMMY ", func(t *testing.T) {
		got, _ := ReadJsonFile("../data/tips.json")
		expected := "DUMMY"
		assert.NotContains(t, string(got), expected)
	})
}
