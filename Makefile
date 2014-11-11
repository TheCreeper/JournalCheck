#!/usr/bin/make -f

SHELL=/bin/sh
bin=bin
name=notify

all: build

build:
	go build -v -o bin/$(name)

clean:
	go clean -x

remove:
	go clean -i
