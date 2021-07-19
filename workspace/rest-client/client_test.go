package restclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadfromEnv(t *testing.T) {
	got := loadfromEnv()
	want := "token"
	assert.Contains(t, got, want)
}
