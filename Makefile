BLANDY_GO_EXECUTABLE ?= go
VERSION := $(shell git describe --tags)
DIST_DIRS := find * -type d -exec

build:
	${BLANDY_GO_EXECUTABLE} build -o blandy -ldflags "-X main.version=${VERSION}" main.go

install: build
	install -d ${DESTDIR}/usr/local/bin/
	install -m 755 ./blandy ${DESTDIR}/usr/local/bin/blandy

test:
	${BLANDY_GO_EXECUTABLE} test . ./command ./conf

clean:
	rm -f ./blandy.test
	rm -f ./blandy
	rm -rf ./dist

bootstrap-dist:
	${BLANDY_GO_EXECUTABLE} get -u github.com/mitchellh/gox

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
	$(DIST_DIRS) tar -zcf blandy-${VERSION}-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r blandy-${VERSION}-{}.zip {} \; && \
	cd ..


.PHONY: build test install clean bootstrap-dist build-all dist