BIN=epos-plugin-populator
VERSION?=makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

.PHONY: build build-release clean

build:
	go build -ldflags "-X epos-plugin-populator/cmd.Version=$(VERSION)" -o $(BIN) .

build-release:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X epos-plugin-populator/cmd.Version=$(VERSION)" -o $(BIN)-$(GOOS)-$(GOARCH)$(EXT) .

clean:
	rm -f $(BIN)*
