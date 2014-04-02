all: build

build:
	go build -v -o notify

clean:
	go clean -x

remove:
	go clean -i

install:
	cp notify /usr/bin/
	cp notify.service /usr/lib/systemd/system

deinstall:
	rm /usr/bin/notify
	rm /usr/lib/systemd/system/notify.service