package compressor

import "errors"

// Merger func is the joining together of some files so that they become one.
func Merger(dist string, srcs ...string) error {
	// TODO
	if len(srcs) <= 0 {
		return errors.New("Source files must require")
	}
	return nil
}
