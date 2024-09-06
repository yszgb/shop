package admin

import (
	"encoding/json"
	"net/http"

	"shop/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

// 向后台的 index.html 传输数据
func (con MainController) Index(c *gin.Context) {
	// 获取 userinfo 对应的 session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	userinfoStr, ok := userinfo.(string) // 类型断言，判断 userinfo 是不是 string
	if ok {
		// 参看 role.go Auth ，权限判断
		// 1. 获取用户信息
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct) //将 json 格式数据，传入到结构体 userinfoStruct

		// 2. 获取所有的权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList) // 在 AccessItem 中找到 module_id=0 的，存到 accessList

		// 3. 获取当前职能的权限，将权限id 放在 map 里
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		// 4. 遍历权限数据，判断权限id 是否在 map 里，是则加上 checked 属性
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		// 5. 传输数据
		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   userinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    userinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "admin/login")
	}
}

// 加载 welcome.html
func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共修改状态，不需刷新
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	// 通过 ABS() 求绝对值，使 0 变成 1， 1 变成 0
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

// 公共修改排序，不需刷新
func (con MainController) ChangeNum(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")
	err1 := models.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}
}