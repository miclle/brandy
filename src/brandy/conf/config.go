package conf

import (
	"github.com/namsral/flag"
)

var Port int

func Init() {
	flag.IntVar(&Port, "port", 3000, "port of brandy")
	flag.IntVar(&Port, "p", 3000, "port of brandy")
	flag.Parse()
}
