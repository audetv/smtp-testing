.PHONY:
.SILENT:

build:
	go build -o ./.bin/mail cmd/cli/main.go

run: build
	./.bin/mail