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

//Integration Testing:
func TestReadFile(t *testing.T) {
	reader := Reader{}
	t.Run("Load Json File and check if it contains the tip starting with Everyday ", func(t *testing.T) {
		got, _ := reader.readFile("../data/tips.json")
		expected := "Everyday Git in twenty commands or so"
		assert.Contains(t, string(got), expected)
	})

}
func TestGet_wd(t *testing.T) {
	reader := Reader{}
	got, _ := reader.get_wd()
	want := "/gophers/workspace/tips"
	assert.Contains(t, got, want)
}
