package action

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbout(t *testing.T) {
	assert := assert.New(t)
	stdout = &bytes.Buffer{}
	About()
	assert.Equal(fmt.Sprint(stdout), aboutMessage, "The two words should be the same.")
}
