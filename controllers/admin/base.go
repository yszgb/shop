package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// 打印信息，跳转到目标页面
func (con BaseController) Success(c *gin.Context, message string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,     // 打印信息
		"redirectUrl": redirectUrl, // 跳转地址
	})
}

func (con BaseController) Error(c *gin.Context, message string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
