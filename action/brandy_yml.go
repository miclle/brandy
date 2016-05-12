package action

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
