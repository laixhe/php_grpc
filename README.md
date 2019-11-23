
> 以前的实例请到 v0.1 版本： https://github.com/laixhe/php_grpc/tree/v0.1

## 让 PHP 搞定 gRPC 不是难事
> https://gitbook.cn/gitchat/activity/5dc9559d0b93ef713b21928c

#### server 是服务端 127.0.0.1:50051
#### grpc_php_plugin 是 php 的 grpc 生成插件 (下载和编译比较长这里提拱了)
#### 这两个文件 必须赋于执行权限，可使用命令 chmod 0755 server 和 chmod 0755 grpc_php_plugin

#### 相关：
```
protoc --php_out=. --grpc_out=. --plugin=protoc-gen-grpc=/xxx/xxx/grpc_php_plugin *.proto
```