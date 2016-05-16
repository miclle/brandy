package compressor

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinify(t *testing.T) {
	assert := assert.New(t)
	src := "../test/test.css"
	dist := "../.tmp/test.min.css"
	tempPath := "../.tmp"

	r, err := os.Open(src)
	assert.Nil(err)
	assert.NotNil(r)
	assert.NotEmpty(r)

	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		err = os.Mkdir(tempPath, os.FileMode(0755))
		assert.Nil(err)
	}

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
