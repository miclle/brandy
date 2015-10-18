package command

import (
	"brandy/conf"
	"log"
)

func NewApp() {
	log.Println("new app command:", conf.Config.Dir)
}
