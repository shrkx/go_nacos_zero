package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//DataSource string          //增加代码
	//Table      string          //增加代码
	//Cache      cache.CacheConf //增加代码
	Transform zrpc.RpcClientConf //新增代码
}
