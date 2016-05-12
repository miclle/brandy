package action

import (
	"io/ioutil"
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
}
