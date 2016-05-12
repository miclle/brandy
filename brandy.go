package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/miclle/brandy/action"
	"github.com/miclle/brandy/logger"
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

func main() {

	app := cli.NewApp()
	app.Name = "brandy"
	app.Usage = usage
	app.Version = version
	app.Author = "Miclle Zheng"
	app.Email = "miclle.zheng@gmail.com"

	app.CommandNotFound = func(c *cli.Context, command string) {
		logger.Log.Errorf("Command %s does not exist.", command)
	}

	app.Before = startup

	app.Commands = commands()

	if err := app.Run(os.Args); err != nil {
		logger.Log.Error(err.Error())
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
				action.Build()
			},
		},
		{
			Name:  "init",
			Usage: "Initialize the configuration",
			Action: func(c *cli.Context) {
				action.Init()
			},
		},
	}
}
