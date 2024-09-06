package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
}

func (con ApiController) Index(c *gin.Context) {
	c.String(http.StatusOK, "接口 /api")
}

func (con ApiController) UserList(c *gin.Context) {
	c.String(http.StatusOK, "接口 /api/userlist")
}

func (con ApiController) Plist(c *gin.Context) {
	c.String(http.StatusOK, "接口 /api/plist")
}

func (con ApiController) CartList(c *gin.Context) {
	c.String(http.StatusOK, "接口 /api/cartlist")
}
