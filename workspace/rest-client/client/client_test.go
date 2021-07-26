package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LoadfromEnv(t *testing.T) {
	t.Run("Checking error", func(t *testing.T) {
		_, err := LoadfromEnv("/data")
		if err != nil {
			assert.Error(t, err)
		}
	})
	t.Run("Checking loaded data", func(t *testing.T) {
		got, _ := LoadfromEnv("../configfile")
		//want := configuration{}
		assert.NotNil(t, got)

	})

}
func Test_GetToken(t *testing.T) {
	t.Run("Checking token type", func(t *testing.T) {
		got, _ := getToken()
		assert.NotNil(t, got)
	})
}
