package action

import (
	"github.com/miclle/brandy/logger"
)

// Build the app
func Build() error {
	config, err := ReadConfig()

	if err != nil {
		logger.Log.Error(err)
		return err
	}

	logger.Log.Info(config)
	return nil
}
