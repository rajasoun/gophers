package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopic(t *testing.T) {
	t.Run("Check GetTopic is equal to git status", func(t *testing.T) {
		mockUserInput := struct {
			title func() string
		}{
			title: func() string { return "git status" },
		}

		got := GetTopic(mockUserInput.title)
		want := "git status"
		assert.Equal(t, got, want)
	})
	t.Run("Check GetTopic is not equal to git dummy", func(t *testing.T) {
		mockUserInput := struct {
			title func() string
		}{
			title: func() string { return "git status" },
		}

		got := GetTopic(mockUserInput.title)
		want := "git dummy"
		assert.NotEqual(t, got, want)
	})
}
