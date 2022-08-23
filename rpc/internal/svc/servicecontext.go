package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shorturl-service/model"
	"shorturl-service/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel //增加代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.Cache)
	return &ServiceContext{
		Config: c,
		Model:  model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), //增加代码
	}
}
