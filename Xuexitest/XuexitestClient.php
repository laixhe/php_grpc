<?php
namespace Xuexitest;

/**
 * service Xuexitest{}
 * 编写 (gprc 定义 Xuexitest 服务)的客户端
 */
class XuexitestClient extends \Grpc\BaseStub{

    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * rpc SayTest(TestRequest) returns (TestReply) {}
     * 方法名尽量和 (gprc 定义 Xuexitest 服务)的方法一样
     * 用于请求和响应该服务
     */
    public function SayTest(\Xuexitest\TestRequest $argument,$metadata=[],$options=[]){
        // (/xuexitest.Xuexitest/SayTest) 是请求服务端那个服务和方法，基本和 proto 文件定义一样
        // (\Xuexitest\TestReply) 是响应信息（那个类），基本和 proto 文件定义一样
        return $this->_simpleRequest('/xuexitest.Xuexitest/SayTest',
            $argument,
            ['\Xuexitest\TestReply', 'decode'],
            $metadata, $options);
    }

}