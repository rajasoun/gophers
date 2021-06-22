package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//:ToDo: Table Driven Tests
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
	t.Run("Get Tip for invalid Topic - dummy", func(t *testing.T) {
		got := GetTip("")
		want := "should not be Empty"
		assert.Equal(t, got, want)
	})

}

func TestLoadTipsFromJson(t *testing.T) {
	t.Run("Load Tips From Json File and check if there are 166 tips ", func(t *testing.T) {
		got, _ := loadTipsFromJson()
		expected := 166
		//Equal asserts that two objects are equal.
		assert.Equal(t, len(got), expected)
	})
}

func TestGetTipWithReadJsonFile(t *testing.T) {
	t.Run("Load Tips From Json File and check file path ", func(t *testing.T) {
		path = "data/tips.json"
		readJsonFile(path)
		got := GetTip("push")
		want := "failed loading jSON file"
		//Equal asserts that two objects are equal.
		assert.Equal(t, got, want)
	})
}
func TestReadJsonFile(t *testing.T) {
	t.Run("Load Json File and check if it contains the tip starting with Everyday ", func(t *testing.T) {
		got, _ := readJsonFile("../data/tips.json")
		expected := "Everyday Git in twenty commands or so"
		//Contains asserts that the specified string, list(array, slice...) or map contains the specified substring or element.
		assert.Contains(t, string(got), expected)
	})
	t.Run("Loading invalid Json File should fail ", func(t *testing.T) {
		_, got := readJsonFile("tips.json")
		//Error asserts that a function returned an error (i.e. not `nil`).
		assert.Error(t, got)
	})
}
