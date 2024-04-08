package util_test

import (
	"testing"

	"github.com/bhanu475/code-kata/util"
	"github.com/matryer/is"
)

func TestIsUrl_ValidUrl(t *testing.T) {
	is := is.New(t)

	// Valid URL
	validUrl := "https://www.example.com"
	is.True(util.IsUrl(validUrl))
}

func TestIsUrl_InvalidUrl(t *testing.T) {
	is := is.New(t)

	// Invalid URL
	invalidUrl := "example.com"
	is.True(!util.IsUrl(invalidUrl))
}

func TestIsUrl_EmptyString(t *testing.T) {
	is := is.New(t)

	emptyString := ""
	is.True(!util.IsUrl(emptyString))
}
