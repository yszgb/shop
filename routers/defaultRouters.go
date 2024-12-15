package routers

import (
	"shop/controllers/reception"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", reception.DefaultController{}.Index)
		defaultRouters.GET("/thumbnail1", reception.DefaultController{}.Thumbnail1)
		defaultRouters.GET("/thumbnail2", reception.DefaultController{}.Thumbnail2)
		defaultRouters.GET("/qrcode1", reception.DefaultController{}.Qrcode1)
		defaultRouters.GET("/qrcode2", reception.DefaultController{}.Qrcode2)
	}
}
