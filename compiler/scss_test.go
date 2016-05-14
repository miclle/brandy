package compiler

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScssToCSS(t *testing.T) {
	assert := assert.New(t)

	r, err := os.Open("../test/compiler.scss")

	assert.Nil(err)
	assert.NotNil(r)

	err = ScssToCSS(r, os.Stdout)
	assert.Nil(err)

}
