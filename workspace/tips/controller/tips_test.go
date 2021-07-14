// +build !integration

package controller

import (
	"bytes"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestGetTipForTopicIntegration(t *testing.T) {
	output_buffer := bytes.Buffer{}
	input_ouputData := []struct {
		name  string
		input string
		want  string
	}{
		{name: "Checking with Valid input", input: "delete", want: "Delete remote branch"},
		{name: "Checking with invalid input", input: "hello", want: "invalid command"},
	}
	for _, tt := range input_ouputData {
		t.Run(tt.name, func(t *testing.T) {
			GetTipForTopic(tt.input, &output_buffer)
			got := output_buffer.String()
			assert.Contains(t, got, tt.want)
		})
	}

}

func Test_GetTipForTopic_LoggerLevel(t *testing.T) {
	t.Run("checking logger level", func(t *testing.T) {
		logLevel, err := logrus.ParseLevel("debug")
		if err != nil {
			t.Fatal(err)
		}
		output_buffer := bytes.Buffer{}
		logrus.SetLevel(logLevel)
		GetTipForTopic("pull", &output_buffer)
		got := output_buffer.String()
		expected := ""
		assert.Equal(t, got, expected)
	})

}
