package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTip(t *testing.T) {
	input_ouputData := []struct {
		name  string
		input string
		want  string
	}{
		{name: "Get Tip for valid Topic - rebase", input: "rebase", want: "Rebases 'feature' to 'master' and merges it in to master  : git rebase master feature && git checkout master && git merge -"},
		{name: "Get Tip for valid Topic - help", input: "help", want: "Everyday Git in twenty commands or so : git help everyday"},
		{name: "Get Tip for invalid Topic - dummy", input: "dummy", want: "Tips Not Available for Topic"},
		{name: "Get Tip for invalid Topic - Empty", input: "", want: "should not be Empty"},
	}
	for _, tt := range input_ouputData {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTip(tt.input)
			assert.Equal(t, got, tt.want)
		})
	}

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
	t.Run("Checking valid data from path ", func(t *testing.T) {
		path = "data/tips.json"
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
