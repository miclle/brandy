package compiler

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSassToCSS(t *testing.T) {
	assert := assert.New(t)

	r, err := os.Open("../test/compiler.scss")

	assert.Nil(err)
	assert.NotNil(r)
	assert.NotEmpty(r)

	err = SassToCSS(r, os.Stdout)
	assert.Nil(err)
	assert.NotEmpty(os.Stdout)
}
