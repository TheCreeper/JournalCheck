all: build

build:
	go build -v -o notify

clean:
	go clean -x

remove:
	go clean -i