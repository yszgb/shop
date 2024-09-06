package routers

import (
	"shop/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		// api 首页
		apiRouters.GET("/", api.ApiController{}.Index)

		// 获取用户列表
		apiRouters.GET("/userlist", api.ApiController{}.UserList)

		// plist 属性列表文件
		apiRouters.GET("/plist", api.ApiController{}.Plist)

		// 购物车列表
		apiRouters.GET("/cartlist", api.ApiController{}.CartList)
	}
}
