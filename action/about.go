package action

import "github.com/miclle/lisa/logger"

const aboutMessage = `
Full stack build system.

Usage: lisa COMMAND [ARGS]

The most common lisa commands are:
 generate    Generate new code (short-cut alias: "g")
 server      Start the lisa server (short-cut alias: "s")
 new         Create a new lisa application. "lisa new my_app" creates a
             new application called MyApp in "./my_app"

In addition to those, there are:
 destroy      Undo code generated with "generate" (short-cut alias: "d")

All commands can be run with -h (or --help) for more information.

More info https://github.com/miclle/lisa
`

// About prints information about lisa.
func About() {
	logger.Puts(aboutMessage)
}
