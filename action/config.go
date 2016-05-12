package action

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/miclle/brandy/logger"
)

// DefaultConfig is default configuration
type DefaultConfig struct {
	App  string `yaml:"app"`
	Dist string `yaml:"dist"`
}

// Config is brandy configuration
type Config struct {
	Development DefaultConfig `yaml:"production"`
	Production  DefaultConfig `yaml:"development"`
	Test        DefaultConfig `yaml:"test"`
}

// Init the configuration
func Init() {
	data := []byte(template)
	err := ioutil.WriteFile("brandy.yml", data, 0644)
	if err != nil {
		panic(err)
	}
}

// ReadConfig read the configuration file
func ReadConfig() (*Config, error) {
	data, err := ioutil.ReadFile("brandy.yml")
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	config := &Config{}

	err = yaml.Unmarshal([]byte(data), config)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	return config, nil
}
