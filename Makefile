# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
NAME=step
TARGETS=darwin linux windows
VERSION=0.1.0

default: build

build:
	packr build -mod=vendor -o "bin/$(NAME)"

fmt:
	gofmt -w $(GOFMT_FILES)

test:
	packr test -v

init:
	packr get ./...

targets: $(TARGETS)

$(TARGETS):
	GOOS=$@ GOARCH=amd64 CGO_ENABLED=0 packr build \
		-mod=vendor \
		-o "dist/$@/$(NAME)" \
		-a -ldflags '-extldflags "-static"'
	zip \
		-j "dist/$(NAME)_v$(VERSION)_$@_amd64.zip" \
		"dist/$@/$(NAME)"

.PHONY: build fmt test init targets $(TARGETS)
