# method 1
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto

# method 2
protoc -I. --go_out=plugins=grpc:. helloworld.proto