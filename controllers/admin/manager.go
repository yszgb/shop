package admin

import (
	"net/http"
	"strings"

	"shop/models"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	// fmt.Printf("%#v", managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	//获取所有的职能
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入参数错误", "/admin/manager/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ") // 获取页面传入的字符串数据，去掉空格
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")

	// 错误情况判断 长度
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "长度不合法", "/admin/manager/add")
		return
	}

	// 错误情况判断 是否存在
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "管理员已存在", "/admin/manager/add")
		return
	}

	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(models.GetUnix()),
	}
	err = models.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
		return
	}
	con.Success(c, "增加管理员成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {
	// 获取管理员
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/manager")
		return
	}
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)

	// 获取职能
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/manager")
		return
	}
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入数据失败", "/admin/manager")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	if len(mobile) != 11 {
		con.Error(c, "传入手机号不合法", "/admin/manager")
		return
	}

	//执行修改
	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	manager.Username = username
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId
	if password != "" {
		//判断密码长度是否合法
		if len(password) < 6 {
			con.Error(c, "密码的长度不合法 密码长度不能小于6位", "/admin/manager/edit?id="+models.String(id))
			return
		}
		manager.Password = models.Md5(password)
	}
	err = models.DB.Save(&manager).Error
	if err != nil {
		con.Error(c, "修改数据失败", "/admin/manager/edit?id="+models.String(id))
		return
	}
	con.Success(c, "修改数据成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "删除数据失败", "/admin/manager")
	} else {
		manager := models.Manager{Id: id}
		models.DB.Delete(&manager)
		con.Success(c, "删除成功", "/admin/manager")
	}
}
