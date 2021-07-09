package cmd

import (
	"testing"
)

func Test_ExecuteCommand_Empty(t *testing.T) {
	cmd := gitCmd()
	cmd.Execute()
}

func Test_ExecuteCommand(t *testing.T) {
	cmd := gitCmd()
	cmd.SetArgs([]string{"push"})
	cmd.Execute()

}
