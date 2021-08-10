// +build !integration

package controller

import (
	"bytes"
	"os"
	"testing"

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
		{name: "Checking with Valid input", input: "git delete", want: "Delete remote branch"},
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
