FROM golang:1.16.5-alpine

WORKDIR /code/
ADD ./ /code/

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o .bin/news-portal cmd/news-portal/main.go
CMD [".bin/news-portal"]