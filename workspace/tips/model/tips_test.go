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
			got := GetTip(tt.input)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestLoadTipsFromJson(t *testing.T) {
	t.Run("Load Tips From Json File and check if there are 166 tips ", func(t *testing.T) {
		got, _ := loadTipsFromJson()
		expected := 166
		assert.Equal(t, len(got), expected)
	})
}

func TestGetTipJsonFilePath(t *testing.T) {
	t.Run("Check Getting Tips Json File Path Dynalically", func(t *testing.T) {
		got := getJsonFilePath()
		want := "/data/tips.json"
		assert.Contains(t, got, want)
	})
}

func TestGetCurrentWorkingDir(t *testing.T) {
	t.Run("checking current dir error", func(t *testing.T) {
		// Mocked function for os.Getwd
		myGetWd := func() (string, error) {
			myErr := errors.New("Simulated error")
			return "", myErr
		}
		// Update the var to this mocked function
		osGetWd = myGetWd
		// This will return error
		_, err := getCurrentWorkingDir()
		assert.Error(t, err)
	})
	t.Run("Checking Current Working directory path", func(t *testing.T) {
		// Mocked function for os.Getwd
		myGetWd := func() (string, error) {
			return "/gophers/workspace/tips", nil
		}
		osGetWd = myGetWd
		got, _ := getCurrentWorkingDir()
		want := "/gophers/workspace/tips"
		assert.Equal(t, got, want)
	})
}
func TestReadJsonFile(t *testing.T) {
	t.Run("Loading invalid Json File should fail", func(t *testing.T) {
		// Mocked function for os.ReadFile
		file_read := func(string) ([]byte, error) {
			myErr := errors.New("Simulated error")
			return nil, myErr
		}
		fileRead = file_read
		_, err := readJsonFile("/data")
		assert.Error(t, err)
	})
	t.Run("Unit Testing readjson file data", func(t *testing.T) {
		file_read := func(string) ([]byte, error) {
			var data = []byte(`[{
				"title":"Rebases 'feature' to 'master' and merges it in to master ",
				"tip":"git rebase master feature && git checkout master && git merge -"
			 }]`)
			return data, nil
		}
		fileRead = file_read
		got, _ := readJsonFile("/gophers/workspace//data/tips.json")
		want := "Rebases 'feature' to 'master'"
		assert.Contains(t, string(got), want)
	})
}
