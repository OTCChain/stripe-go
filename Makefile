SHELL=PATH='$(PATH)' /bin/sh

GOBUILD=CGO_ENABLED=0 go build -ldflags '-w -s'

PLATFORM := $(shell uname -o)

NAME := chord.exe
OS := windows

ifeq ($(PLATFORM), Msys)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else ifeq ($(PLATFORM), Cygwin)
    INCLUDE := ${shell echo "$(GOPATH)"|sed -e 's/\\/\//g'}
else
	INCLUDE := $(GOPATH)
	NAME=chord
	OS=linux
endif

# enable second expansion
.SECONDEXPANSION:

.PHONY: all
.PHONY: pbs
.PHONY: test

BINDIR=$(INCLUDE)/bin

all: pbs build

build:
	GOOS=$(OS) GOARCH=amd64 $(GOBUILD) -o $(BINDIR)/$(NAME)

pbs:
	cd pbs/ && $(MAKE)

mac:
	GOOS=darwin go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).mac
arm: pbs
	GOOS=linux GOARM=7 GOARCH=arm go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).arm
linux: pbs
	GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).lnx
win: pbs
	GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' -o $(BINDIR)/$(NAME).exe

sol:
	cd eth/generated/ && $(MAKE)

clean:
	rm $(BINDIR)/$(NAME)
