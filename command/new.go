package command

import (
	"log"

	"github.com/miclle/brandy/conf"
)

func NewApp() {
	log.Println("new app command:", conf.Config.Dir)
}
