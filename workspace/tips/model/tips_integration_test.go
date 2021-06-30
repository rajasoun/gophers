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
func TestReadJsonFile(t *testing.T) {
	file_reader_Impl := File_reader_Impl{}
	t.Run("Load Json File and check if it contains the tip starting with Everyday ", func(t *testing.T) {
		got, _ := file_reader_Impl.readJsonFile("../data/tips.json")
		expected := "Everyday Git in twenty commands or so"
		assert.Contains(t, string(got), expected)
	})

}
