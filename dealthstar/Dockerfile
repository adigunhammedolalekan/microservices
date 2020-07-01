FROM golang:alpine3.11

COPY . /app
WORKDIR /app

RUN go get ./...
ENV GOOS linux
RUN go build -o dealthstar cmd/main.go
EXPOSE 30002
ENTRYPOINT [ "./dealthstar" ]
