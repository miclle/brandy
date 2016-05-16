package compressor

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Merger func is the joining together of some files so that they become one.
func Merger(dist string, srcs ...string) error {

	if len(srcs) <= 0 {
		return errors.New("Source files must required")
	}

	dir := filepath.Dir(dist)

	log.Print(dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.FileMode(0755)); err != nil {
			log.Print(err.Error())
			return err
		}
	}

	_, err := os.Stat(dist)

	var f *os.File

	if err != nil {
		if os.IsNotExist(err) {
			if f, err = os.Create(dist); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if f, err = os.OpenFile(dist, os.O_RDWR, os.FileMode(0644)); err != nil {
			return err
		}
	}

	w := bufio.NewWriter(f)

	for _, src := range srcs {
		f, err := os.Open(src)

		defer f.Close()

		if err != nil {
			return err
		}

		r := bufio.NewReader(f)

		buffer := make([]byte, 1024)

		for {
			n, err := r.Read(buffer)
			if err == io.EOF {
				break
			} else {
				w.Write(buffer[:n])
			}
		}
	}

	return w.Flush()
}
