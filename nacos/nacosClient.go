package nacos

import (
	"gateway/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
)

var Client naming_client.INamingClient
var ConfigClient config_client.IConfigClient

func InitNacosServer() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.NacosServerConfig.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			config.NacosServerConfig.Host,
			uint64(config.NacosServerConfig.Port),
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}

	Client, _ = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	ConfigClient, _ = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	RegisterInstance(getIp(), config.WebServerConfig.Port, config.NacosServerConfig.ServiceName, config.NacosServerConfig.GroupName)
	GetConfig(config.NacosServerConfig.ConfigDataId, config.NacosServerConfig.GroupName)
	ListenConfig(config.NacosServerConfig.ConfigDataId, config.NacosServerConfig.GroupName)
}

func getIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
