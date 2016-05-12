package action

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)
	Init()
	dat, err := ioutil.ReadFile("brandy.yml")
	assert.Nil(err)
	assert.NotNil(dat)
	assert.Equal(string(dat), template, "The two words should be the same.")

	fi, err := os.Stat("brandy.yml")
	assert.Nil(err)
	assert.NotNil(fi)
	assert.Equal(fi.Mode(), os.FileMode(0644), "The brandy.yml file model should be 0644 `-rw-r--r--`")
}

func TestReadConfig(t *testing.T) {
	assert := assert.New(t)
	config, err := ReadConfig()
	assert.Nil(err)
	assert.NotNil(config)

	for _, c := range []DefaultConfig{
		config.Development,
		config.Production,
		config.Test,
	} {
		assert.NotNil(c)
		assert.Equal(c.App, "src")
		assert.Equal(c.Dist, "dist")
	}
}
