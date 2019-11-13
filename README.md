
> 以前的实例请到 v0.1 版本： https://github.com/laixhe/php_grpc/tree/v0.1

## 让 PHP 搞定 gRPC 不是难事
> https://gitbook.cn/gitchat/activity/5dc9559d0b93ef713b21928c

#### server 是服务端 127.0.0.1:50052
#### grpc_php_plugin 是 php 的 grpc 生成插件

#### 相关：
```
protoc --php_out=. --grpc_out=. --plugin=protoc-gen-grpc=/xxx/xxx/grpc_php_plugin *.proto
```