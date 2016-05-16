package compressor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerger(t *testing.T) {
	assert := assert.New(t)

	srcs := []string{
		"../test/bootstrap.css",
		"../test/bootstrap-theme.css",
	}

	dist := "../.tmp/css/bootstrap/bootstrap-all.css"

	err := Merger(dist, srcs...)

	assert.Nil(err)

	var srcSum int64

	for _, src := range srcs {
		srcFileInfo, err := os.Stat(src)
		assert.Nil(err)
		assert.NotNil(srcFileInfo)

		srcSum = srcSum + srcFileInfo.Size()
	}

	distFileInfo, err := os.Stat(dist)
	assert.Nil(err)
	assert.NotNil(distFileInfo)

	assert.True(srcSum >= distFileInfo.Size())
}
