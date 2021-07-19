package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadfromEnv(t *testing.T) {
	got := loadfromEnv()
	want := 8
	assert.Equal(t, len(got), want)
}
