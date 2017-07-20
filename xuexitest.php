<?php
//引入 composer 的自动载加
require __DIR__ . '/vendor/autoload.php';

//用于连接 服务端
$client = new \Xuexitest\XuexitestClient('192.168.0.60:50052', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

//实例化 TestRequest 请求类
$request = new \Xuexitest\TestRequest();
$request->setTypeid(1);

//调用远程服务
$get = $client->SayTest($request)->wait();

//返回数组
//$reply 是 TestReply 对象
//$status 是数组
list($reply, $status) = $get;

//数组
$getdata = $reply->getGetdataarr();

foreach ($getdata as $k=>$v){
    echo $v->getId(),'=>',$v->getName(),"\n\r";
}

//
//echo "\n\r";
//$countid = $getdata->count();
//
//for ($i=0;$i < $countid;$i++){
//    echo $getdata->offsetGet($i)->getName(),"\n\r";
//}
//
////
//echo "\n\r";
//$getdatas = $reply->getGetdataarr()->getIterator();
//
//
//while (true){
//    $isid = $getdatas->valid();
//    if($isid){
//
//        echo $getdatas->current()->getName(),"\n\r";
//
//        $getdatas->next();
//    }else{
//        break;
//    }
//}
