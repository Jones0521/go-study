# 如何编译proto文件

```
# 进入当前probuf定义的目录
$ cd protocol
# 指定protobuf 文件搜索位置为 当前目录 -I=.
$ protoc -I=. -I=D:\Go\bin\include --go_out=. --go_opt=module="github.com/Jones0521/go-study/go6-rpc/grpc-demo/grpc/" hello.proto
# GRPC 
$ protoc -I=. -I=D:\Go\bin\include --go_out=. --go_opt=module="server/protocol" --go-grpc_out=. --go-grpc_opt=module="server/protocol" hello.proto

```