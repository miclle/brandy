package compressor

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
)

// Merger func is the joining together of some files so that they become one.
func Merger(dist string, srcs ...string) error {

	if len(srcs) <= 0 {
		return errors.New("Source files must require")
	}

	_, err := os.Stat(dist)

	if os.IsNotExist(err) {
		if err = os.MkdirAll(dist, os.FileMode(0644)); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(dist, os.O_RDWR, os.FileMode(0644))

	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	for _, src := range srcs {
		f, err := os.Open(src)
		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(f)

		if err != nil {
			return err
		}

		w.Write(bytes)
	}

	return w.Flush()
}
