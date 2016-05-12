package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	Init()
	assert := assert.New(t)
	err := Build()
	assert.Nil(err)
}
