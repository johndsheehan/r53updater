.DEFAULT_GOAL := build

TAG = $(shell date +"%Y%m%d-%H%M%S")

build: clean
	cd cmd/r53updater ; go get -u ; go build

clean:
	rm -rf cmd/r53updater/r53updater

docker:
	docker build  --force-rm  -t r53updater:$(TAG)  -f Dockerfile  .

install:
	cp cmd/r53updater/r53updater  /usr/local/bin

uninstall:
	rm /usr/local/bin/r53updater
