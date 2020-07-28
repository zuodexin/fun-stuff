# 用Go语言和gRPC实现一个控制台聊天程序

## 定义消息，编译proto
```
protoc --go_out=plugins=grpc:. -I . chat.proto
```

## 编写server端

具体步骤如下：

- 实现gRPC函数
- 绑定grpc server
- 开启服务
  
## 编写client端
具体步骤如下：

- 客户端绑定连接
- 完成gRPC调用


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