.PHONY: clean build deps

build:
	go build

clean:
	go clean -i github.com/thehungry-dev/log...

deps:
	go get && go mod tidy
