package cli

import (
	"strings"
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

func TestCase(t *testing.T) {
	t.Run("Check GetTopicc with reader interface", func(t *testing.T) {
		key := strings.NewReader("git commit")
		got, _ := GetTopicc(key)
		want := "git commit"
		assert.Equal(t, got, want)
	})
	t.Run("Check GetTopicc with reader interface", func(t *testing.T) {
		key := strings.NewReader("com")
		_, err := GetTopicc(key)
		assert.Error(t, err)
	})
}
