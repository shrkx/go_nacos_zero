package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"shorturl-service/api/internal/config"
	"shorturl-service/rpc/transformclient"
)

type ServiceContext struct {
	Config config.Config
	//Model  model.ShorturlModel //增加代码
	Transformer transformclient.Transform //新增代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//Model:  model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), //增加代码
		Transformer: transformclient.NewTransform(zrpc.MustNewClient(c.Transform)), //新增代码
	}

}
