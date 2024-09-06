package admin

import (
	"fmt"
	"net/http"
	"os"
	"shop/models"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

// 加载轮播图首页
func (con FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

// 加载增加轮播图页面
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

// 执行增加轮播图，跳转
func (con FocusController) DoAdd(c *gin.Context) {
	// 获取 focus 表单数据
	title := c.PostForm("title")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))
	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
	}

	// 上传文件
	focusImgSrc, err4 := models.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
		con.Error(c, "上传文件出错", "/admin/focus/add")
	} else {
		focus := models.Focus{
			Title:     title,
			FocusType: focusType,
			FocusImg:  focusImgSrc,
			Link:      link,
			Sort:      sort,
			Status:    status,
			AddTime:   int(models.GetUnix()),
		}
		err5 := models.DB.Create(&focus).Error
		if err5 != nil {
			con.Error(c, "增加轮播图失败", "/admin/focus/add")
		} else {
			con.Success(c, "增加轮播图成功", "/admin/focus")
		}
	}
}

// 加载编辑轮播图页面
func (con FocusController) Edit(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "传入参数错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

// 执行编辑轮播图，跳转
func (con FocusController) DoEdit(c *gin.Context) {
	// 获取数据
	id, err1 := models.Int(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))
	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
	}

	// 上传文件
	focusImg, _ := models.UploadImg(c, "focus_img")
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+models.String(id))
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}

// 执行删除轮播图，跳转
func (con FocusController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id")) // 从 HTTP 请求中获取 id
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
	} else {
		focus := models.Focus{Id: id}
		models.DB.Where("id = ?", id).First(&focus) // 查询单条语句

		// 删除图片
		err1 := os.Remove(focus.FocusImg)
		if err1 != nil {
			con.Error(c, "删除图片失败", "/admin/focus/edit")
		} else {
			// 删除数据库数据
			models.DB.Delete(&focus)
			con.Success(c, "删除数据成功", "/admin/focus")
		}
	}
}
