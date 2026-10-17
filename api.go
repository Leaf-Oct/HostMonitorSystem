package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应html资源的
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func exceptionHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "exception.html", gin.H{})
}

func aboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}

// 响应json格式的数据
func queryHandler(c *gin.Context) {
	table_name := c.Param("table")
	if table_name == "" {

	}
}
