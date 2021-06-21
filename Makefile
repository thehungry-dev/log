.PHONY: clean build deps test

build:
	go build

clean:
	go clean -i github.com/thehungry-dev/log...

deps:
	go get && go mod tidy

test:
	LOG_COLOR=on \
	LOG_DEVICE=stderr \
	LOG_FILTERS=level,tag \
	LOG_LEVEL=_all \
	LOG_OUTPUT_FORMAT=text \
	LOG_TAGS_OUTPUT=on \
	LOG_TAGS=_all \
	go test
