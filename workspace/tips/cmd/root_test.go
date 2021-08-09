package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func Test_NewRootCmd(t *testing.T) {
	t.Run("checking valid inputs", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		rootCmd.SetArgs([]string{})
		err := rootCmd.Execute()
		if err != nil {
			t.Fatal(err)
		}
		out, err := ioutil.ReadAll(outputBuffer)
		if err != nil {
			t.Fatal(err)
		}
		got := string(out)
		want := "help"
		assert.Contains(t, got, want, "expected \"%s\" got \"%s\"", want, got)
	})
	t.Run("checking invalid user inputs", func(t *testing.T) {
		inputBuffer := "dummy"
		rootCmd.SetArgs([]string{inputBuffer})
		err := rootCmd.Execute()
		if err != nil {
			assert.Error(t, err)
		}
		assert.Error(t, err)
	})
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
		flag  string
	}{
		{"Success Case", "stash", "stash", "git"},
		//{"Error Case", "help", "help", "--tips"},
		//{"Invalid Data", "gf", "help", "git"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs([]string{tt.flag, tt.want})
			//NewRootCmd().Flags().Set(tt.flag, tt.input)
			writer := &bytes.Buffer{}
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
		want := "help"
		assert.Contains(t, got, want, "want \"%s\" got \"%s\"", want, got)
	})

	t.Run("Checking valid data", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		expected := "checkout"
		rootCmd.SetArgs([]string{"git", expected})
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
	t.Run("Checking valid data", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		expected := "du"
		rootCmd.SetArgs([]string{"git", expected})
		err := gitCmd.Execute()
		if err != nil {
			//t.Fatal(err)
			assert.Error(t, err)
		}
		out, err := ioutil.ReadAll(outputBuffer)
		if err != nil {
			//t.Fatal(err)
			assert.Error(t, err)
		}
		got := string(out)
		assert.Contains(t, got, "help", "expected \"%s\" got \"%s\"", "help", got)
	})
	t.Run("checking valid command", func(t *testing.T) {
		outputBuffer := bytes.NewBufferString("")
		rootCmd.SetOut(outputBuffer)
		rootCmd.SetArgs([]string{"gi"})
		err := gitCmd.Execute()
		assert.Error(t, err)
	})

}
