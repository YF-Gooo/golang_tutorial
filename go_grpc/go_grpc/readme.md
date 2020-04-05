go mod init main
https://github.com/protocolbuffers/protobuf/releases 下载protoc
export GOPROXY=https://goproxy.cn
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

protoc --go_out=plugins=grpc:. ./helloworld.proto