BIN=epos-plugin-populator
VERSION?=makefile

.PHONY: build build-release clean

build:
	go build -ldflags "-X epos-plugin-populator/cmd.Version=$(VERSION)" -o $(BIN) .

build-release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X epos-plugin-populator/cmd.Version=$(VERSION)" -o $(BIN) .

clean:
	rm -f $(BIN)
