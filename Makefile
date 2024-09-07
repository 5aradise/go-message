build:
	CGO_ENABLED=1 go build -C cmd/go-message/ -o ../../bin/go-message

debug: build
	GIN_MODE=debug ./bin/go-message

run: build
	GIN_MODE=release ./bin/go-message