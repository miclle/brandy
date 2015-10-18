package conf

import (
	"github.com/namsral/flag"
	"log"
	"os"
	"path/filepath"
)

var Dir string
var Port int

func Init() {
	flag.StringVar(&Dir, "dir", "", "dir of brandy")
	flag.IntVar(&Port, "port", 3000, "port of brandy")
	flag.Parse()

	// Get the directory of the currently
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if Dir == "" {
		Dir = dir
	}
}
