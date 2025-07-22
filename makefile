BIN=epos-plugin-populator
VERSION?=makefile
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

.PHONY: build build-release clean vet test

build: generate
	go build -ldflags "-X epos-plugin-populator/cmd.Version=$(VERSION)" -o $(BIN) .

# Build for specific platform (used by CI)
build-release:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X epos-plugin-populator/cmd.Version=$(VERSION)" -o $(BIN)-$(GOOS)-$(GOARCH)$(EXT) .

clean:
	rm -f $(BIN)*

generate:
	go generate ./...

vet:
	go vet ./...

test:
	go test ./... -v
