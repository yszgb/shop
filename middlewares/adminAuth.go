package middlewares

import (
	"encoding/json"
	"fmt"
	"os"
	"shop/models"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

// 用户权限判断，判断是否登录 未登录跳转到登陆页面
func InitAdminAuthMiddleware(c *gin.Context) {
	// 1. 获取 Url 访问的地址
	// admin/login 页面目标 Url 是 /admin/captcha?t=xxx
	// 后面的 t=xxx 不需要
	// pathname := c.Request.Url.String() // 得到 /admin/captcha?t=xxx
	pathname := strings.Split(c.Request.URL.String(), "?")[0]

	// 2. 获取 session 里的信息
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	// 3. 判断 session 里的信息是否正确
	// 类型断言 判断 userinfo 是不是字符串，不是说明传入失败
	userinfoStr, ok := userinfo.(string)
	if ok {
		// 判断 userinfo 里的信息是否存在
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		} else { // 用户登录成功，权限判断，不能跳转到没有访问权限的页面
			urlPath := strings.Replace(pathname, "/admin/", "", 1) // 将第 1 次出现的 "/admin/" 替换为空
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				// 1. 获取当前职能的权限，将 权限id 放入 map
				roleAccess := []models.RoleAccess{}
				models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccess {
					roleAccessMap[v.AccessId] = v.AccessId
				}

				// 2. 获取访问 url 需要的 权限
				access := models.Access{}
				models.DB.Where("url = ?", urlPath).Find(&access)

				// 3. 判断权限在不在权限表
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}
	} else { // session 不存在，未登录，跳转到 login 页面
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}
}

// 判断要排除的页面
func excludeAuthPath(urlPath string) bool {
	// 加载 ini 配置文件
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",") // 按 "," 分隔

	// 传入的 urlPath 和 app.ini 里不用权限判断的地址比较
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
