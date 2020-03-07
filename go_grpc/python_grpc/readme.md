pip install grpcio grpcio-tools protobuf -i https://pypi.tuna.tsinghua.edu.cn/simple
protoc -I . –python_out=tmp –grpc_out=tmp –plugin=protoc-gen-grpc=tmp/grpc_python_plugin helloworld.proto