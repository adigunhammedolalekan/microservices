export GOOS=linux
go build -o api/api api/cmd/main.go
go build -o dealthstar/dealthstar dealthstar/cmd/main.go
go build -o destroyer/destroyer destroyer/cmd/main.go
