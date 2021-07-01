// +build integration

package model

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

//integration testing:
func TestReadFile(t *testing.T) {
	t.Run("Load Json File and check if it contains the tip starting with Everyday ", func(t *testing.T) {
		got, _ := readJsonFile("../data/tips.json")
		expected := "Everyday Git in twenty commands or so"
		assert.Contains(t, string(got), expected)
	})

}
func TestGet_wd(t *testing.T) {
	t.Run("checking current working directory path", func(t *testing.T) {
		got, _ := getCurrentWorkingDir()
		want := "/gophers/workspace/tips"
		assert.Contains(t, got, want)
	})
}
