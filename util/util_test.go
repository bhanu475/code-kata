package util_test

import (
	"testing"

	"github.com/bhanu475/code-kata/util"
	"github.com/stretchr/testify/assert"
)

func TestIsUrl_ValidUrl(t *testing.T) {
	// Valid URL
	validUrl := "https://www.example.com"
	isValid := util.IsUrl(validUrl)
	assert.True(t, isValid, "Expected true for a valid URL")
}

func TestIsUrl_InvalidUrl(t *testing.T) {
	// Invalid URL
	invalidUrl := "example.com"
	isValid := util.IsUrl(invalidUrl)
	assert.False(t, isValid, "Expected false for an invalid URL")
}

func TestIsUrl_IPAddress(t *testing.T) {
	// IP address
	ipAddress := "192.168.0.1"
	isValid := util.IsUrl(ipAddress)
	assert.True(t, isValid, "Expected true for an IP address")
}

func TestIsUrl_Localhost(t *testing.T) {
	// Localhost
	localhost := "localhost"
	isValid := util.IsUrl(localhost)
	assert.True(t, isValid, "Expected true for localhost")
}

func TestIsUrl_EmptyString(t *testing.T) {
	// Empty string
	emptyString := ""
	isValid := util.IsUrl(emptyString)
	assert.False(t, isValid, "Expected false for an empty string")
}
