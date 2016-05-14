package compiler

import (
	"io"
	"log"

	"github.com/wellington/go-libsass"
)

// SassToCSS is convert Sass to CSS func
func SassToCSS(src io.Reader, dst io.Writer) error {

	comp, err := libsass.New(dst, src)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := comp.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
