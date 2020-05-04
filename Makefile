## Makefile

# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GOOS ?= $(call lc, $(shell uname -s))
GOARCH ?= amd64

SOURCES = $(wildcard */cmd/api/main.go)
PROJECTS = $(foreach p, $(dir $(SOURCES)), $(p:/=))
BINARIES = target/app

print_vars:
	echo $(SOURCES)
	echo $(PROJECTS)
	echo $(BINARIES)
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $(BINARIES) $(SOURCES)
clean:
	@rm -vf $(BINARIES)/$(BINARIES)-bin

.PHONY: build clean $(BINARIES)/$(BINARIES)-bin
