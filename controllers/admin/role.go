package admin

import (
	"net/http"
	"strings"

	"shop/models"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	// 职能不多，可以不做分页，直接 Find 打印全部职能
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	// fmt.Println(roleList) // 打印数据库信息
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ") // Trim 去除空格
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" { // 要继承 BaseController
		con.Error(c, "职能名不能为空", "/admin/role/add")
		return
	}
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())

	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "增加职能失败", "/admin/role/add")
	} else {
		con.Success(c, "增加职成功", "/admin/role/")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	// 获取页面传来的 id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/role")
	} else {
		// 显示数据库中的信息
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}
}

// 按下“提交”按钮，执行 DoEdit 方法
func (con RoleController) DoEdit(c *gin.Context) {
	// 获取页面传来的 id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/role")
	}
	title := strings.Trim(c.PostForm(("title")), " ")           // 获取 title，去掉空格
	description := strings.Trim(c.PostForm("description"), " ") // 获取表单中 description，去掉空格
	if title == "" {
		con.Error(c, "标题不能为空", "/admin/role/edit")
		return
	}
	// 传到数据库
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	err = models.DB.Save(&role).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}
}

func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(c, "删除成功", "/admin/role")
	}
}

// 传输 role access 数据到前端
func (con RoleController) Auth(c *gin.Context) {
	// 1. 获取职能Id roleId
	roleId, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "Auth Get role id wrong! ", "/admin/role")
		return
	}

	// 2. 获取权限 Auth
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList) // query association relationship

	// 3. 获取当前职能的权限，放在 map 里
	// Using map to quickly retrieve and store all AccessId corresponding to role_id
	// Auth can be used then
	roleAccess := []models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int) // map is fast & can
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	// 4. 判断当前职能是否在 map 里
	// Determine current role is in roleAccessMap
	// Loop through access, if current id is in the Map, add "Checked" attribute
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ { // Check
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}

	// 向前端传输信息
	// Transmit to HTML
	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})
}

func (con RoleController) DoAuth(c *gin.Context) {
	// Get roleId
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "DoAuth Get roleId wrong! ", "/admin/role")
		return
	}

	// Get accessIds, from array
	accessIds := c.PostFormArray("access_node[]")

	// 删除当前职能的权限，避免数据库中数据重复
	// Delete access of current role firstly to avoid duplicate data
	roleAccess := models.RoleAccess{}
	models.DB.Where("role_id=?", roleId).Delete(&roleAccess)

	// 授予权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := models.Int(v)
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}

	// 跳转
	con.Success(c, "DoAuth success! ", "/admin/role/auth?id="+models.String(roleId))
}
