build: clean install style
	go install -v ./src/brandy/...

build-linux: clean style
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v ./src/brandy/...

install:
	cd src; glide install
  
clean:
	cd src; go clean -i ./...
	rm -rf pkg bin

style:
	go vet ./src/brandy/...

test:
	go test ./src/brandy/...

glide:
	cd src; glide $(filter-out $@,$(MAKECMDGOALS))

%:
	@: