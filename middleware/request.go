package middleware

import (
	"encoding/json"
	"gateway/config"
	"gateway/nacos"
	"gateway/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		groupName := config.File.MustValue("nacos_server", "groupName", "DEFAULT_GROUP")
		//convertNewUrl(c)
		service := nacos.GetService("lspos-finance", groupName, "")
		host := utils.WeightRandom(service.Hosts)
		marshal, _ := json.Marshal(host)
		log.Println(string(marshal))
	}
}

func convertNewUrl(c *gin.Context) {
	err := convertUrl(c.Request.URL)
	if err != nil {
		log.Println("转换转发请求地址有误", err)
	}
	println(c.Request.URL.String())
	req, err := http.NewRequestWithContext(c, c.Request.Method, c.Request.URL.String(), c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, make(map[string]string))
	c.Abort()
}

/**
根据nacos配置 获取实际访问ip:host
*/
func convertUrl(rawUrl *url.URL) error {
	config, _ := json.Marshal(nacos.GateWayConfig)
	println(string(config))
	proxyUrl := "http://localhost:8503/lspos-finance/test"
	u, err := url.Parse(proxyUrl)
	if err != nil {
		return err
	}

	rawUrl.Scheme = u.Scheme
	rawUrl.Host = u.Host
	rawUrl.Path = "/lspos-finance/test"
	ruq := rawUrl.Query()
	rawUrl.RawQuery = ruq.Encode()
	return nil
}
