<?php
//引入 composer 的自动载加
include "vendor/autoload.php";

// 用于连接 服务端
$client = new Simple\UserClient('127.0.0.1:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);


// 实例化 GetUserRequest 请求类
$request = new Simple\GetUserRequest();
$request->setUserid(1);

// 调用远程服务
// 获取某个 user 数据 - 返回数组
$get = $client->GetUser($request)->wait();

// $user   是 GetUserResponse 对象
// $status 是 记录 grpc 错误信息 对象
list($user, $status) = $get;

if($status->code == 0){
    echo 'user_id=', $user->getUserid(), "\n";
    echo 'user_name=', $user->getUsername(), "\n";
    echo 'sex=', $user->getSex(), "\n";
}

