.DEFAULT_GOAL := run

run:
	go run cmd/news-portal/main.go
lint:
	golangci-lint run



