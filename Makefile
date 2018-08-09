GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

setup:
	$(GOGET) -u github.com/golang/dep/cmd/dep
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./interface/server/router/main.go
test:
	$(GOTEST) -cover ./...
deps:
	dep ensure -v
