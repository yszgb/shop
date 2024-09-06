package main

import (
	"html/template"
	"shop/models"
	"shop/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToDate,
	})

	// 加载模板
	r.LoadHTMLGlob("templates/**/**/*")

	// 加载静态 web 目录
	r.Static("/static", "./static")

	// 创建基于 cookie 的存储引擎， secret111 为密钥
	store := cookie.NewStore([]byte("secret111"))

	// 配置 session 中间件在 Redis 上
	// store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	// 80 端口，直接 localhost 访问
	r.Run(":80")
}
