// +build !integration

package model

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestGetTip(t *testing.T) {
	readerMockImpl := readerMockImpl{}
	input_ouputData := []struct {
		name  string
		input string
		want  string
	}{
		{name: "Get Tip for valid Topic - rebase", input: "rebase", want: "Rebases 'feature' to 'master' and merges it in to master  : git rebase master feature && git checkout master && git merge -"},
		{name: "Get Tip for invalid Topic - dummy", input: "dummy", want: "Tips Not Available for Topic"},
		{name: "Get Tip for invalid Topic - Empty", input: "", want: "should not be Empty"},
	}
	for _, tt := range input_ouputData {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTip(tt.input, readerMockImpl)
			assert.Equal(t, got, tt.want)
		})
	}

}
func TestLoadTipsFromJson(t *testing.T) {
	t.Run("Load Tips From Json File and check if there are 166 tips ", func(t *testing.T) {
		readerMockImpl := readerMockImpl{}
		got, _ := loadTipsFromJson(readerMockImpl)
		expected := 1
		assert.Equal(t, len(got), expected)
	})
}

func TestGetCurrentWorkingDir(t *testing.T) {
	handlerMockImpl := &readerMockImpl{}
	t.Run("Checking Current Working directory path", func(t *testing.T) {
		got, _ := getCurrentWorkingDir(handlerMockImpl)
		want := "/gophers/workspace/tips"
		assert.Equal(t, got, want)
	})

}

func TestGetTipJsonFilePath(t *testing.T) {
	t.Run("Check Getting Tips Json File Path Dynalically", func(t *testing.T) {
		readerMockImpl := readerMockImpl{}
		got := getJsonFilePath(readerMockImpl)
		want := "/data/tips.json"
		assert.Contains(t, got, want)
	})
}

func TestReadJsonFile(t *testing.T) {
	readerMockImpl := readerMockImpl{}
	t.Run("Unit Testing readjson file data", func(t *testing.T) {
		got, _ := readJsonFile("../data/tips.json", readerMockImpl)
		want := "Rebases 'feature' to 'master'"
		assert.Contains(t, string(got), want)
	})
	t.Run("Loading invalid Json File should fail ", func(t *testing.T) {
		_, got := readJsonFile("tips.json", readerMockImpl)
		assert.Error(t, got)
	})

}

// getWd mock impl & readFile mock impl
type readerMockImpl struct{}

func (reader_mock_Impl readerMockImpl) get_wd() (string, error) {
	return "/gophers/workspace/tips", nil
}

func (reader_mock_Impl readerMockImpl) readFile(path string) ([]byte, error) {
	var data = []byte(`[{
		"title":"Rebases 'feature' to 'master' and merges it in to master ",
		"tip":"git rebase master feature && git checkout master && git merge -"
	 }]`)
	if path == "tips.json" {
		return nil, errors.New("error")
	}
	return data, nil
}
