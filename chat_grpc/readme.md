# 用Go语言和gRPC实现一个控制台聊天程序

## 定义消息，编译proto
```
protoc --go_out=plugins=grpc:. -I . chat.proto
```
## 编译server,client
```
go build client.go chat.pb.go
go build server.go chat.pb.go
```

## 运行server
```
./sever
```

## 运行client
```
./client
```