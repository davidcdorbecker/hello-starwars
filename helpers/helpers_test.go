package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedirectHttps(t *testing.T) {
	testUrl := "http://api.com/test"

	result := RedirectHttps(testUrl)

	assert.EqualValues(t, "https://api.com/test", result)
}