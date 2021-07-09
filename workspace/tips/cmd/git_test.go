package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := gitCmd()

	t.Run("Check error if argument length less than 1", func(t *testing.T) {
		cmd.SetArgs([]string{})
		err := cmd.Execute()
		assert.Error(t, err)
	})
	t.Run("Check output if argument length more than 1", func(t *testing.T) {
		cmd.SetArgs([]string{"push"})
		err := cmd.Execute()
		assert.NoError(t, err)
	})
}
