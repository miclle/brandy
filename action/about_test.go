package action

import (
	"bytes"
	"testing"

	"github.com/miclle/brandy/logger"
)

func TestAbout(t *testing.T) {
	var buf bytes.Buffer
	old := logger.Default.Stdout
	logger.Default.Stdout = &buf
	About()

	if buf.Len() < len(aboutMessage) {
		t.Errorf("expected this to match aboutMessage: %q", buf.String())
	}

	logger.Default.Stdout = old
}
