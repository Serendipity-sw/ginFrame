package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
设置gin路由规则
创建人:邵炜
创建时间:2017年2月9日13:51:48
输入参数:gin engine
*/
func setGinRouter(r *gin.Engine) {
	g := &r.RouterGroup
	if rootPrefix != "" {
		g = r.Group(rootPrefix)
	}
	{
		g.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	}
}
