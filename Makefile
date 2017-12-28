BINARY := stream-set
VERSION := v0.0.1

.PHONY: windows
windows:
    mkdir -p release
    GOOS=windows GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-windows-amd64.exe

.PHONY: linux
linux:
    mkdir -p release
    GOOS=linux GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-linux-amd64

.PHONY: darwin
darwin:
    mkdir -p release
    GOOS=darwin GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-darwin-amd64

.PHONY: release
release: windows linux darwin