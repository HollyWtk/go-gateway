package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func RegisterInstance(host string, port int64, serviceName string, groupName string) {
	success, err := Client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		//Metadata:    map[string]string{"idc": "shanghai"},
		GroupName: groupName,
	})
	if success {
		log.Printf("nacos注册成功,ip:%s,host:%d,serviceName:%s,groupName:%s \n", host, port, serviceName, groupName)
	}
	if err != nil {
		log.Println("nacos注册失败", err)
		panic(err)
	}
}

func GetConfig(dataId string, groupName string) {
	content, err := ConfigClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName})
	if err != nil {
		log.Println("配置文件获取失败", err)
	} else {
		convert(content)
	}
}

func ListenConfig(dataId string, groupName string) {
	_ = ConfigClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			convert(data)
		},
	})
}

func SelectAllInstances(serviceName string, groupName string, clusters ...string) []model.Instance {
	if len(clusters) == 0 {
		clusters[1] = "DEFAULT"
	}
	instances, err := Client.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: serviceName,
		GroupName:   groupName, // default value is DEFAULT_GROUP
		Clusters:    clusters,  // default value is DEFAULT
	})
	if err != nil {
		log.Printf("获取服务实例失败,服务名:%s,组名:%s,集群:%s", serviceName, groupName, clusters)
	}
	return instances
}

func convert(content string) {
	err := ConvertConfig(content)
	if err != nil {
		log.Println("配置文件数据转换错误", err)
	}
}

func GetService(serviceName string, groupName string, clusters ...string) model.Service {
	if len(clusters) == 0 {
		clusters[1] = "DEFAULT"
	}
	services, err := Client.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
		GroupName:   groupName, // default value is DEFAULT_GROUP
		Clusters:    clusters,  // default value is DEFAULT
	})
	if err != nil {
		log.Printf("获取服务实例失败,服务名:%s,组名:%s,集群:%s", serviceName, groupName, clusters)
	}
	return services
}
