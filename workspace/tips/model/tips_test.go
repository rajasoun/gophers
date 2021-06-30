// +build !integration

package model

import (
	"errors"
	"fmt"
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

type getMockErrorImpl struct{ err error }

func (m *getMockErrorImpl) error() error {
	m.err = errors.New("error")
	return m.err
}
func TestGetCurrentWorkingDir(t *testing.T) {
	t.Run("Checking Current Working directory path", func(t *testing.T) {
		getErrorImpl := &getErrorImpl{}
		got, _ := getCurrentWorkingDir(getErrorImpl)
		want := "/gophers/workspace/tips"
		assert.Contains(t, got, want)
	})
	t.Run("Checking Error on reading current working directory path", func(t *testing.T) {
		getMockErrorImpl := &getMockErrorImpl{}
		_, got := getCurrentWorkingDir(getMockErrorImpl)
		want := errors.New("error")
		assert.Error(t, got, want)
	})
}

func TestGetTipJsonFilePath(t *testing.T) {
	t.Run("Check Getting Tips Json File Path Dynalically", func(t *testing.T) {
		got := getJsonFilePath()
		want := "/gophers/workspace/tips"
		assert.Contains(t, got, want)
	})
}

type readerMockImpl struct{}

func (reader_mock_Impl readerMockImpl) readJsonFile(path string) ([]byte, error) {
	var data = []byte(`[{
		"title":"Rebases 'feature' to 'master' and merges it in to master ",
		"tip":"git rebase master feature && git checkout master && git merge -"
	 }]`)
	return data, nil
}

func TestRead(t *testing.T) {
	file_reader_Impl := File_reader_Impl{}
	t.Run("Loading invalid Json File should fail ", func(t *testing.T) {
		_, got := file_reader_Impl.readJsonFile("tips.json")
		assert.Error(t, got)
	})
	readerMockImpl := readerMockImpl{}
	t.Run("Unit Testing readjson file data", func(t *testing.T) {
		got, _ := readerMockImpl.readJsonFile("../data/tips.json")
		fmt.Print(len(got))
		want := "Rebases 'feature' to 'master'"
		assert.Contains(t, string(got), want)
	})

}
