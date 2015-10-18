package conf

import (
	"github.com/namsral/flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Conf struct {
	Dir  string
	Port int
}

var Config Conf

func InitArgs() {

	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "", 0)

	fs.StringVar(&Config.Dir, "dir", "", "dir of brandy")
	fs.IntVar(&Config.Port, "port", 3000, "port of brandy")

	for i, arg := range os.Args {
		if strings.HasPrefix(arg, "-") {
			fs.Parse(os.Args[i:])
			break
		}
	}

	// Get the directory of the currently
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if Config.Dir == "" {
		Config.Dir = dir
	}

	log.Println("Config.Dir", Config.Dir)
}
