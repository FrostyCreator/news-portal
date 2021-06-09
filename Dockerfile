FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o bin/news-portal cmd/news-portal/main.go

ENTRYPOINT ["bin/news-portal"]