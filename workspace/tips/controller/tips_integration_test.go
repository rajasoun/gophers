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

func TestGetTipForTopicIntegrationTesting(t *testing.T) {
	output_buffer := bytes.Buffer{}
	controller := Controller_Impl{}
	input_ouputData := []struct {
		name  string
		input string
		want  string
	}{
		{name: "Delete", input: "delete", want: "Delete remote branch"},
		{name: "Empty string", input: "", want: "should not be Empty"},
	}
	for _, tt := range input_ouputData {
		t.Run(tt.name, func(t *testing.T) {
			GetTipForTopic(tt.input, &output_buffer, controller)
			got := output_buffer.String()
			assert.Contains(t, got, tt.want)
		})
	}
}
