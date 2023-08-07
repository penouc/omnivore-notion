package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// Omnibus是一个 post 接口
// 输出所有post数据
func Omnibus(c *gin.Context) {
	// 获取请求体
	requestBody, err := c.GetRawData()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 解析请求体中的 JSON 数据
	var data map[string]interface{}
	if err := json.Unmarshal(requestBody, &data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 使用解析后的数据进行操作
	// ...

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
