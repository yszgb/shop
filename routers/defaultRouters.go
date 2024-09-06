package routers

import (
	"shop/controllers/reception"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", reception.DefaultController{}.Index)
		defaultRouters.GET("/news", reception.DefaultController{}.News)
	}
}
