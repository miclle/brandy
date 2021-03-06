package action

import (
	"fmt"
	"io"
	"os"
)

const aboutMessage = `
Full stack build system.

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

var stdout io.Writer = os.Stdout

// About prints information about brandy.
func About() {
	fmt.Fprint(stdout, aboutMessage)
}
