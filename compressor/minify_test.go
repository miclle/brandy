package compressor

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinify(t *testing.T) {
	assert := assert.New(t)
	src := "../test/compiler.css"
	dist := "../test/compiler.min.css"

	r, err := os.Open(src)
	assert.Nil(err)
	assert.NotNil(r)
	assert.NotEmpty(r)

	f, err := os.Create(dist)
	assert.Nil(err)

	w := bufio.NewWriter(f)

	err = Minify(MediaTypeTextCSS, w, r)
	assert.Nil(err)
	assert.Nil(w.Flush())

	srcFileInfo, err := os.Stat(src)
	assert.Nil(err)
	assert.NotNil(srcFileInfo)

	distFileInfo, err := os.Stat(dist)
	assert.Nil(err)
	assert.NotNil(distFileInfo)

	assert.True(srcFileInfo.Size() >= distFileInfo.Size())
}
