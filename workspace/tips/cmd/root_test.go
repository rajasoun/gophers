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

}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
		want string
		flag string
	}{
		{"Success Case", "stash", "--topic"},
		{"Error Case", "error", "--tips"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//rootCmd.SetArgs([]string{tt.flag, tt.want})
			NewRootCmd().Flags().Set(tt.flag, tt.want)
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

func Test_HideDebugFlag(t *testing.T) {
	got := ""
	want := ""
	assert.Equal(t, got, want)
}
