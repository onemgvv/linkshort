include .env

build:
	go build -o .bin/main cmd/main.go
run: build
	./.bin/main
dev:
	go run cmd/main.go