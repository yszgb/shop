package admin

import (
	"net/http"
	"strings"

	"shop/models"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload(("AccessItem")).Find(&accessList)
	// fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) Add(c *gin.Context) {
	// 获取顶级分类
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {
	// 获取表单数据
	moduleName := c.PostForm("module_name")
	accessType, err1 := models.Int(c.PostForm("type")) // type 是保留字，改名叫 accessType
	actionName := c.PostForm("action_name")
	url := c.PostForm("url")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "parameter wrong!", "/access/add")
		return
	}

	// 无错误，实例化数据库结构体
	access := models.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err5 := models.DB.Create(&access).Error
	if err5 != nil {
		con.Error(c, "增加数据失败", "/admin/access/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/access")
}

// Load access/edit.html
func (con AccessController) Edit(c *gin.Context) {
	// 获取页面传来的 id
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "传入数据失败", "/admin/access")
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)

	// 获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})

}

// Do Edit Access
func (con AccessController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err2 := models.Int(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err3 := models.Int(c.PostForm("module_id"))
	sort, err4 := models.Int(c.PostForm("sort"))
	status, err5 := models.Int(c.PostForm("status"))
	description := c.PostForm("description")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		con.Error(c, "传入参数错误", "/admin/access")
		return
	}
	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "/admin/access/edit?id="+models.String(id))
		return
	}

	// 向数据库传入数据
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status

	err := models.DB.Save(&access).Error
	if err != nil {
		con.Error(c, "修改数据", "/admin/access/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/access/edit?id="+models.String(id))
	}
}

// Delete Access 操作
func (con AccessController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/access")
	} else {
		// 获取要删除的数据
		access := models.Access{Id: id}
		models.DB.Find(&access)
		if access.ModuleId == 0 { // 顶级模块
			accessList := []models.Access{}
			models.DB.Where("module_id = ?", access.Id).Find(&accessList)
			if len(accessList) > 0 {
				con.Error(c, "当前模块下有菜单或操作，请删除子数据后再删除这个数据", "/admin/access")
			} else {
				models.DB.Delete(&access)
				con.Success(c, "删除数据成功", "/admin/access")
			}
		} else { //操作或菜单
			models.DB.Delete(&access)
			con.Success(c, "删除数据成功", "/admin/access")
		}
	}
}
