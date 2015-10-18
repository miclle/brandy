package command

import (
	"fmt"
	"os"
)

func Help() {
	fmt.Println(help)
	os.Exit(0)
}

var help = `
Usage: brandy COMMAND [ARGS]

The most common brandy commands are:
 generate    Generate new code (short-cut alias: "g")
 server      Start the brandy server (short-cut alias: "s")
 new         Create a new brandy application. "brandy new my_app" creates a
             new application called MyApp in "./my_app"

In addition to those, there are:
 destroy      Undo code generated with "generate" (short-cut alias: "d")

All commands can be run with -h (or --help) for more information.

More info https://github.com/miclle/brandy
`
