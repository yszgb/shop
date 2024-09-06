package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"shop/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, answer, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println("验证码错误")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"captchaId":     id,
			"captchaImg":    b64s,
			"captchaAnswer": answer,
		})
	}
}

// 验证验证码
func (con LoginController) DoLogin(c *gin.Context) {
	// 接收从前端、客户端传到后端，用 PostForm
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")

	// 1. 验证验证码，防止攻击
	verifyValue := c.PostForm("verifyValue")

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {

		// 2. 验证码正确，查询数据库
		userinfoList := []models.Manager{}
		password = models.Md5(password)
		models.DB.Where("username=? AND password=?", username, password).Find(&userinfoList)

		// 验证码正确
		if len(userinfoList) > 0 {

			// 3. session 保存用户信息
			session := sessions.Default(c)
			// 结构体转换为 JSON 字符串
			uesrinfoSlice, _ := json.Marshal(userinfoList)
			session.Set("userinfo", string(uesrinfoSlice)) // Set 只能保存字符串，无法保持切片。切片转为字符串
			session.Save()

			// 4. 跳转
			con.Success(c, "输入正确，登陆成功", "/admin/")
		} else {
			con.Error(c, "用户名或密码输入错误", "/admin/login")
		}
	} else {
		con.Error(c, "验证码输入错误", "/admin/login")
	}
}

// 清除 session 后，跳转
func (con LoginController) LoginOut(c *gin.Context) {
	// 清除 session
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()

	con.Success(c, "退出登陆", "/admin/login")
}
