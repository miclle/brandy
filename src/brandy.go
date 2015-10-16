package main

import (
	"log"
	"os"
	// "gopkg.in/fsnotify.v1"
)

func main() {
	var str string
	if len(os.Args) > 1 {
		str = os.Args[1]
	}
	log.Println(str)

	switch str {
	case "help", "h":
		log.Println("help task")

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
