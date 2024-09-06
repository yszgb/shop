package reception

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(c *gin.Context) {
	c.SetCookie("username", "zhangsan", 3600, "/", "localhost", false, false)

	// 设置 session
	session := sessions.Default(c)

	// 设置过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 6 hrs
	})

	session.Set("username1", "zhangsan111")
	session.Save() // 设置 session 时 必须调用

	// 如果要传入结构体，要改成 JSON 字符串 （每一项后面加 `JSON）
	c.SetCookie("hobby", "吃饭 睡觉", 3600, "/", "localhost", false, false)

	// 渲染模板文件
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		// 传入数据
		// "a":1,
		"title":     "首页",
		"timestamp": 1629788418,
		"score":     85,
	})
}

func (con DefaultController) News(c *gin.Context) {
	// 获取 cookie
	// 需要先访问首页
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")

	c.String(http.StatusOK, "新闻页面 /news"+"\n")
	c.String(http.StatusOK, "cookie = "+username+"\n")
	c.String(http.StatusOK, "cookie = "+hobby+"\n")

	session := sessions.Default(c)
	username1 := session.Get("username1")
	c.String(http.StatusOK, "username1 = %v\n", username1)
}

func (con DefaultController) shop(c *gin.Context) {
	// 获取 cookie
	username, _ := c.Cookie("username")
	hobby, _ := c.Cookie("hobby")

	c.String(http.StatusOK, "/shop"+"\n")
	c.String(http.StatusOK, "cookie = "+username+"\n")
	c.String(http.StatusOK, "cookie = "+hobby+"\n")
}
