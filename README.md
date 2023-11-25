# Install dependencies
- Install protoc https://grpc.io/docs/protoc-installation/
- Install plugins for GO
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

# Generate new protobuf interfaces
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative agify/agify.proto
```