package config

import (
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

const configFile = "/conf/conf.ini"

var File *goconfig.ConfigFile

var NacosServerConfig = &NacosServer{}

type NacosServer struct {
	Host         string
	Port         int64
	NamespaceId  string
	GroupName    string
	ServiceName  string
	ConfigDataId string
}

var WebServerConfig = &WebServer{}

type WebServer struct {
	Host string
	Port int64
}

//加载此文件的时候 会先走初始化方法
func init() {
	//拿到当前的程序的目录
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configPath := currentDir + configFile

	if !fileExist(configPath) {
		log.Printf("路径:%s,未读取到配置文件", configPath)
	}
	//参数  mssgserver.exe  D:/xxx
	length := len(os.Args)
	if length > 1 {
		dir := os.Args[1]
		log.Printf("读取路径:%s配置文件", dir)
		if dir != "" {
			configPath = dir + configFile
		}
	}
	//文件系统的读取
	File, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		log.Fatal("读取配置文件出错:", err)
	}
	initServer()
}

func initServer() {
	webHost := File.MustValue("web_server", "host", "127.0.0.1")
	webPort := File.MustInt64("web_server", "port", 8080)
	nacosHost := File.MustValue("nacos_server", "host", "127.0.0.1")
	nacosPort := File.MustInt64("nacos_server", "port", 8848)
	namespaceId := File.MustValue("nacos_server", "namespaceId", "8848")
	serviceName := File.MustValue("nacos_server", "serviceName", "defaultServiceName")
	groupName := File.MustValue("nacos_server", "groupName", "DEFAULT_GROUP")
	configDataId := File.MustValue("nacos_server", "configDataId", "")
	NacosServerConfig = &NacosServer{
		Host:         nacosHost,
		Port:         nacosPort,
		NamespaceId:  namespaceId,
		ServiceName:  serviceName,
		GroupName:    groupName,
		ConfigDataId: configDataId,
	}
	WebServerConfig = &WebServer{
		Host: webHost,
		Port: webPort,
	}

}

func fileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
