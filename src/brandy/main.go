package main

import (
	"brandy/command"
	"brandy/conf"
	"fmt"
	"log"
	"net/http"
	"os"
	// "gopkg.in/fsnotify.v1"
)

func main() {

	conf.InitArgs()

	var str string

	if len(os.Args) > 1 {
		str = os.Args[1]

		switch str {
		case "help", "h":
			command.Help()

		case "new", "n":
			command.NewApp()
			os.Exit(0)

		case "clean", "c":
			log.Println("clean task")

		case "build", "b":
			log.Println("build task")

		case "server", "s":
			log.Println("server task")

		case "verbose", "v":
			log.Println("verbose task")

		case "version", "V":
			log.Println("version task")

		default:
			log.Println("default task")
		}
	}

	log.Printf("Listening on localhost:%d, CTRL+C to stop\n", conf.Config.Port)
	log.Println("Serving files from", conf.Config.Dir)

	addr := fmt.Sprintf(":%d", conf.Config.Port)
	handler := http.FileServer(http.Dir(conf.Config.Dir))
	panic(http.ListenAndServe(addr, handler))
}
