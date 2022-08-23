package nacos

import (
	"encoding/json"
	"log"
)

var GateWayConfig = &Config{}

type Config struct {
	RefreshGatewayRoute bool
	RouteList           *[]RouteList
}

type RouteList struct {
	Id         string
	Predicates *[]Predicate
	Filters    *[]Filter
	Uri        string
	Order      int
}

type Predicate struct {
	Name string
	Args map[string]any `json:"args"`
}
type Filter struct {
	Name string
	Args map[string]any `json:"args"`
}

func ConvertConfig(content string) {
	err := json.Unmarshal([]byte(content), GateWayConfig)
	if err != nil {
		log.Println("gateway配置数据转换失败", err)
	}
}
