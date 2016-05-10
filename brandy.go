package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/miclle/brandy/action"
	"github.com/op/go-logging"
	// "gopkg.in/fsnotify.v1"
	// "gopkg.in/yaml.v2"
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

var log = logging.MustGetLogger("brandy")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
}

func main() {

	app := cli.NewApp()
	app.Name = "brandy"
	app.Usage = usage
	app.Version = version
	app.Author = "Miclle Zheng"
	app.Email = "miclle.zheng@gmail.com"

	app.CommandNotFound = func(c *cli.Context, command string) {
		log.Errorf("Command %s does not exist.", command)
	}

	app.Before = startup

	app.Commands = commands()

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
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
		{
			Name:        "server",
			Usage:       "Start a static file server",
			Description: "Start the brandy server",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:  "build",
			Usage: "Build the project",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
		{
			Name:  "init",
			Usage: "Initialize the configuration",
			Action: func(c *cli.Context) {
				// TODO
			},
		},
	}
}
