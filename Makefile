.DEFAULT_GOAL := run

build:
	go mod download && go build -o ./.bin/news-portal ./cmd/news-portal/main.go

run: build
	./.bin/news-portal

lint:
	golangci-lint run