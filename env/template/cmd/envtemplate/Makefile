.PHONY: all
all: envtemplate

envtemplate: Makefile main.go
	CGO_ENABLED=0 go build -ldflags="-s -w"
