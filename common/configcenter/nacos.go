package configcenter

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
	"strings"
)

// 定义参数
var (
	TimeoutMs           uint64 = 5000 // 定义超时时间
	NotLoadCacheAtStart bool   = true // 开始时不加载缓存
	LogDir              string = "/logs/nacos/logger"
	CacheDir            string = "/logs/nacos/cacheDir"
	LogLevel            string = "debug"
)

type (
	NacosService interface {
		// 1. 服务的链接
		serviceLink() (naming_client.INamingClient, config_client.IConfigClient)
		// 2. 服务的注册
		ServiceRegistration(serverName, ListenOn string)
		// 3. 获取所有的服务
		FindInstance() *model.Instance
		// 4. 获取服务的配置信息
		CreateConfigClient(listenConfigCallback ListenConfig) string
		// 5. 获取服务的列表
	}
	// 定义服务的链接
	NacosServerConfig struct {
		IpAddr string
		Port   uint64
	}
	// 定义链接器的操作
	NacosClientConfig struct {
		NamespaceId string
	}

	// 定义返回的操作
	ListenConfig func(data string)

	// 定义配置
	NacosConfig struct {
		ServerConfigs []NacosServerConfig
		ClientConfig  NacosClientConfig
		DataId        string
		Group         string
	}

	ServerInfo struct {
	}

	// 默认Nacos
	defaultNacos struct {
		cfg NacosConfig
	}
	// 引导配置
	BootstrapConfig struct {
		NacosConfig NacosConfig
	}
)

// 定义新的构造方法
func NacosFactory(config BootstrapConfig) NacosService {
	return &defaultNacos{config.NacosConfig}
}

// 1. 服务的链接
func (d defaultNacos) serviceLink() (naming_client.INamingClient, config_client.IConfigClient) {
	//nacos server
	var sc []constant.ServerConfig
	if len(d.cfg.ServerConfigs) == 0 {
		panic("nacos server no set")
	}

	// 定义服务的地址
	for _, serveConfig := range d.cfg.ServerConfigs {
		sc = append(sc, constant.ServerConfig{
			Port:   serveConfig.Port,
			IpAddr: serveConfig.IpAddr,
		},
		)
	}
	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         d.cfg.ClientConfig.NamespaceId, //当namespace是public时，此处填空字符串。
		TimeoutMs:           TimeoutMs,                      // 定义超时时间
		NotLoadCacheAtStart: NotLoadCacheAtStart,            // 开始时不加载缓存
		LogDir:              LogDir,                         //定义日志的位置
		CacheDir:            CacheDir,                       // 定义缓存的日志目录
		LogLevel:            LogLevel,                       //定义日志的输出格式
	}
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}
	return client, configClient
}

// 接口以实现
func (d defaultNacos) ServiceRegistration(serverName, ListenOn string) {
	// 进行数据的切割
	split := strings.Split(ListenOn, ":")
	IP := split[0]
	port, err := strconv.ParseUint(split[1], 10, 64)
	// 使用服务的链接
	client, _ := d.serviceLink()
	ok, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          IP,
		Port:        port,
		ServiceName: serverName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{},
		GroupName:   d.cfg.Group,
		Ephemeral:   true,
	})
	if err != nil {
		panic(err)
	}
	if !ok {
		panic(errors.New("注册本服务发生错误"))
	}
}

// 进行服务的发现 => 获取指定服务的信息
func (d defaultNacos) FindInstance() *model.Instance {
	// 使用服务的链接
	client, _ := d.serviceLink()
	ins, err := client.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		GroupName:   d.cfg.Group,
		ServiceName: "transform.rpc",
	})
	if err != nil {
		panic(err.Error())
	}
	return ins
}

// 获取配置的信息
func (d defaultNacos) CreateConfigClient(listenConfigCallback ListenConfig) string {
	_, client := d.serviceLink()
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: d.cfg.DataId,
		Group:  d.cfg.Group,
	})
	if err != nil {
		panic(err)
	}
	// 监听数据的编号
	err = client.ListenConfig(vo.ConfigParam{
		DataId: d.cfg.DataId,
		Group:  d.cfg.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件发生了变化...")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			listenConfigCallback(data)
		},
	})
	listenConfigCallback(content)
	return content
}
