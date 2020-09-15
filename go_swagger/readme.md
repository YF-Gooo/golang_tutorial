go env -w GOPROXY=https://goproxy.cn,direct
go get -u github.com/swaggo/swag/cmd/swag

swag init

go run main.go

http://localhost:8080/swagger/index.html