package action

import "io/ioutil"

var template = `
# Be sure to restart brandy when you modify this file.

default: &default
  app: src
  dist: dist
  port: 8080
  # proxy: 
  # target: 

production:
  <<: *default

development:
  <<: *default

test:
  <<: *default
`

// Init the configuration
func Init() {
	data := []byte(template)
	err := ioutil.WriteFile("brandy.yml", data, 0644)
	if err != nil {
		panic(err)
	}
}
