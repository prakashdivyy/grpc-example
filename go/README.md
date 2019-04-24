# Go Example

## How to run
```
# Install protoc and protoc-gen-go
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

# compile protobuff
protoc --proto_path=example --proto_path=../ --go_out=plugins=grpc:example example.proto

# run client
go run client.go

# run server
go run server.go
```