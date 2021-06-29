// +build !integration

package controller

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestGetTipForTopic(t *testing.T) {
	output_buffer := bytes.Buffer{}
	controller_mock := controller_mock{}
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
			GetTipForTopic(tt.input, &output_buffer, controller_mock)
			got := output_buffer.String()
			assert.Contains(t, got, tt.want)
		})
	}
}

type controller_mock struct{}

func (d controller_mock) getTip(input string) string {
	data := `{"title":"Delete remote branch","tip":"git push origin --delete <remote_branchname>",}`
	data = string(data)
	if strings.Contains(data, input) && input != "" {
		return data
	} else if input == "" {
		return "should not be Empty"
	}
	return "Tips Not Available for Topic"
}
