FROM golang:alpine3.11

COPY . /app
WORKDIR /app

RUN go get ./...
ENV GOOS linux
RUN go build -o destroyer cmd/main.go
EXPOSE 30001
ENTRYPOINT [ "./destroyer" ]
