BRANDY_GO_EXECUTABLE ?= go
VERSION := $(shell git describe --tags)
DIST_DIRS := find * -type d -exec

build:
	${BRANDY_GO_EXECUTABLE} build -o brandy -ldflags "-X main.version=${VERSION}" brandy.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./brandy ${DESTDIR}/usr/local/bin/brandy

test:
	${BRANDY_GO_EXECUTABLE} test ./...

clean:
	rm -f ./brandy.test
	rm -f ./brandy
	rm -rf ./dist

bootstrap-dist:
	${BRANDY_GO_EXECUTABLE} get -u github.com/mitchellh/gox

build-all:
	gox -verbose \
	-ldflags "-X main.version=${VERSION}" \
	-os="linux darwin windows " \
	-arch="amd64 386" \
	-output="dist/{{.OS}}-{{.Arch}}/{{.Dir}}" .

dist: build-all
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf brandy-${VERSION}-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r brandy-${VERSION}-{}.zip {} \; && \
	cd ..


.PHONY: build test install clean bootstrap-dist build-all dist