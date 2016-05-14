package compiler

import (
	"io"
	"log"

	"github.com/wellington/go-libsass"
)

// ScssToCSS is convert Sass to CSS func
func ScssToCSS(src io.Reader, dst io.Writer) error {

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
