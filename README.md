
### Note ✍🏻

1) Paths 🎏
```
go get google.golang.org/protobuf
go get google.golang.org/grpc

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*user*.proto
```

¯\_(ツ)_/¯