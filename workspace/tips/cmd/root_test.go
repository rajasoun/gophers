package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}
func Test_NewRootCmdwithInvalidCommand(t *testing.T) {
	t.Run("checking unknown commands expect git", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		rootCmd.SetArgs([]string{"gi"})
		err := rootCmd.Execute()
		assert.Error(t, err)

	})
}
func Test_NewRootCmd(t *testing.T) {
	t.Run("checking valid inputs", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		expected := "push"
		rootCmd.SetArgs([]string{"--topic", expected})
		err := rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}
		out, err := ioutil.ReadAll(outputBuffer)
		if err != nil {
			t.Fatal(err)
		}
		got := string(out)
		want := expected
		assert.Contains(t, got, want, "expected \"%s\" got \"%s\"", want, got)
	})
	t.Run("checking invalid user inputs", func(t *testing.T) {
		inputBuffer := ""
		rootCmd.SetArgs([]string{"--topic", inputBuffer})
		err := rootCmd.Execute()
		if err != nil {
			assert.Error(t, err)
		}
	})
	t.Run("checking help commands", func(t *testing.T) {
		rootCmd.SetArgs([]string{})
		err := rootCmd.Execute()
		if err != nil {
			assert.NoError(t, err)
		}
	})

	t.Run("checking invalid input data", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		var inValidInput string = "lo"
		rootCmd.SetArgs([]string{"--topic", inValidInput})
		err := rootCmd.Execute()
		if err != nil {
			assert.Error(t, err)
		}
		out, _ := ioutil.ReadAll(outputBuffer)
		if err != nil {
			assert.Error(t, err)
		}
		got := string(out)
		want := "help"
		assert.Contains(t, got, want, "expected \"%s\" got \"%s\"", want, got)
	})

}

func TestExecute(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		flag  string
	}{
		{"Success Case", "stash", "stash", "--topic"},
		{"Error Case", "help", "help", "--tips"},
		{"Invalid Data", "gf", "help", "--topic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//rootCmd.SetArgs([]string{tt.flag, tt.want})
			writer := &bytes.Buffer{}
			NewRootCmd(writer).Flags().Set(tt.flag, tt.input)
			err := Execute(writer)
			if err != nil {
				assert.Error(t, err)
			} else {
				gotWriter := writer.String()
				assert.Contains(t, gotWriter, tt.want)
			}
		})
	}
}

func Test_SetLogger(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		level string
	}{
		{"Checking set level logger ", "", "debug"},
		{"invalid level logger", "error", "dummy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := bytes.Buffer{}
			err := setUpLogs(&output, tt.level)
			if err != nil {
				assert.Error(t, err)
			} else {
				got := output.String()
				want := tt.want
				assert.Equal(t, got, want)
			}

		})
	}

}
func Test_GitCommandwithinvalidCommand(t *testing.T) {
	t.Run("Checking invalid passing command", func(t *testing.T) {
		expected := "dummy"
		rootCmd.SetArgs([]string{"git", expected})
		err := gitCmd.Execute()
		assert.Error(t, err)
	})
}
func Test_GitCommand(t *testing.T) {
	t.Run("checking help command", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		rootCmd.SetArgs([]string{"git"})
		err := gitCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}
		out, err := ioutil.ReadAll(outputBuffer)
		if err != nil {
			t.Fatal(err)
		}
		got := string(out)
		fmt.Print(got)
		want := "help"
		assert.Contains(t, got, want, "want \"%s\" got \"%s\"", want, got)
	})

	t.Run("Checking valid data", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		expected := "checkout"
		rootCmd.SetArgs([]string{"git", "--arg", expected})
		err := gitCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}
		out, err := ioutil.ReadAll(outputBuffer)
		if err != nil {
			t.Fatal(err)
		}
		got := string(out)
		assert.Contains(t, got, expected, "expected \"%s\" got \"%s\"", expected, got)
	})
}
