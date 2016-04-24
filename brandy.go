package main

import (
	"os"

	"github.com/miclle/brandy/action"

	"github.com/codegangsta/cli"
	// "gopkg.in/fsnotify.v1"

	"github.com/miclle/brandy/logger"
)

var version = "0.0.1-dev"

const usage = `Full stack build system.

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

func main() {

	app := cli.NewApp()
	app.Name = "brandy"
	app.Usage = usage
	app.Version = version
	app.Author = "Miclle Zheng"
	app.Email = "miclle.zheng@gmail.com"

	app.CommandNotFound = func(c *cli.Context, command string) {
		logger.ExitCode(99)
		logger.Die("Command %s does not exist.", command)
	}

	app.Before = startup

	app.Commands = commands()

	if err := app.Run(os.Args); err != nil {
		logger.Err(err.Error())
		os.Exit(1)
	}

	// If there was a Error message exit non-zero.
	if logger.HasErrored() {
		m := logger.Color(logger.Red, "An Error has occurred")
		logger.Msg(m)
		os.Exit(2)
	}

	// addr := fmt.Sprintf(":%d", conf.Config.Port)
	// handler := http.FileServer(http.Dir(conf.Config.Dir))
	// panic(http.ListenAndServe(addr, handler))
}

func startup(c *cli.Context) error {
	// TODO
	return nil
}

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "about",
			Usage: "Learn about brandy",
			Action: func(c *cli.Context) {
				action.About()
			},
		},
	}
}
