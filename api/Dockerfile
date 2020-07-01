FROM golang:alpine3.11

COPY . /app
WORKDIR /app

RUN go get ./...
ENV GOOS linux
RUN go build -o api cmd/main.go
EXPOSE 2004
ENTRYPOINT [ "./api" ]
