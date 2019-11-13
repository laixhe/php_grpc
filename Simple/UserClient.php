<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Simple;

/**
 * 定义 User 服务
 */
class UserClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * 定义 GetUser 方法 - 获取某个 user 数据
     * @param \Simple\GetUserRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetUser(\Simple\GetUserRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/simple.User/GetUser',
        $argument,
        ['\Simple\GetUserResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * 定义 GetUserList 方法 - 获取 user 所有数据
     * @param \Simple\GetUserListRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetUserList(\Simple\GetUserListRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/simple.User/GetUserList',
        $argument,
        ['\Simple\UserListResponse', 'decode'],
        $metadata, $options);
    }

}
