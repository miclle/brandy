package main

import (
	"brandy/conf"
	"fmt"
	"log"
	"net/http"
	"os"
	// "gopkg.in/fsnotify.v1"
)

func main() {

	conf.Init()

	var str string

	if len(os.Args) > 1 {
		str = os.Args[1]

		switch str {
		case "help", "h":
			log.Println("help task")

		case "new", "n":
			log.Println("new task")

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

	log.Printf("Listening on localhost:%d, CTRL+C to stop\n", conf.Port)
	log.Println("Serving files from", conf.Dir)

	addr := fmt.Sprintf(":%d", conf.Port)
	handler := http.FileServer(http.Dir(conf.Dir))
	panic(http.ListenAndServe(addr, handler))
}
