package middleware

import (
	"errors"
	"gateway/config"
	"gateway/nacos"
	"gateway/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		convertNewUrl(c)
	}
}

func convertNewUrl(c *gin.Context) {
	var scheme string
	if c.Request.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}
	err := convertUrl(c.Request.URL, scheme)
	if err != nil {
		log.Println("转换转发请求地址有误", err)
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "service not found",
		})
		return
	}
	println(c.Request.URL.String())
	req, err := http.NewRequestWithContext(c, c.Request.Method, c.Request.URL.String(), c.Request.Body)
	req.Header = c.Request.Header
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
func convertUrl(rawUrl *url.URL, scheme string) error {
	groupName := config.NacosServerConfig.GroupName
	println(rawUrl.Path)
	realPath := rawUrl.Path[1:]
	var serviceName string
	index := strings.Index(realPath, "/")
	if index > 0 {
		serviceName = realPath[:index]
	} else {
		serviceName = realPath
	}
	service := nacos.GetService(serviceName, groupName, "")
	if len(service.Hosts) == 0 {
		return errors.New("service not found")
	} else {
		host := utils.WeightRandom(service.Hosts)
		rawUrl.Scheme = scheme
		rawUrl.Host = host.Ip + ":" + strconv.FormatUint(host.Port, 10)
		ruq := rawUrl.Query()
		rawUrl.RawQuery = ruq.Encode()
	}
	return nil
}
