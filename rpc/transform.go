package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"shorturl-service/common/configcenter"
	_const "shorturl-service/common/const"
	"shorturl-service/rpc/internal/config"
	"shorturl-service/rpc/internal/server"
	"shorturl-service/rpc/internal/svc"
	"shorturl-service/rpc/transform"
	"sigs.k8s.io/yaml"
)

var configFile = flag.String("f", "etc/transform.yaml", "the config file")

func main() {
	flag.Parse()
	// 进行绑定的操作
	var c_dev config.ConfigDev
	conf.MustLoad(*configFile, &c_dev)
	var c config.Config
	// 进行服务的在线注册操作
	bootstrapConfig := configcenter.BootstrapConfig{
		NacosConfig: configcenter.NacosConfig{
			DataId: _const.Nacos_DataId,
			Group:  _const.Nacos_Group,
			ServerConfigs: []configcenter.NacosServerConfig{
				{
					IpAddr: _const.Nacos_Host,
					Port:   _const.Nacos_Port,
				},
			},
			ClientConfig: configcenter.NacosClientConfig{
				NamespaceId: _const.Nacos_Names_PaceId,
			},
		},
	}
	// 进行链接的操作
	factory := configcenter.NacosFactory(bootstrapConfig)
	// 进行服务的注册
	factory.ServiceRegistration(c_dev.Name, c_dev.ListenOn)
	// 返回在获取在线的配置信息
	factory.CreateConfigClient(func(data string) {
		body := bytes.TrimPrefix([]byte(data), []byte("\xef\xbb\xbf"))
		yamlstring, err := yaml.YAMLToJSON(body)
		// 再用json转为结构体
		json.Unmarshal(yamlstring, &c)
		if err != nil {
			fmt.Println("error:", err)
		}
	})
	// 进行和系统的赋值
	ctx := svc.NewServiceContext(c)
	s := zrpc.MustNewServer(c_dev.RpcServerConf, func(grpcServer *grpc.Server) {
		transform.RegisterTransformServer(grpcServer, server.NewTransformServer(ctx))
		if c_dev.Mode == service.DevMode || c_dev.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
