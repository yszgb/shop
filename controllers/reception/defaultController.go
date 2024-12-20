package reception

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hunterhug/go_image"
	qrcode "github.com/skip2/go-qrcode"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(200, "首页")
}

func (con DefaultController) Thumbnail1(c *gin.Context) {
	// 按宽度进行比例缩放，输入输出都是文件
	// 参数：filename string, savepath string, width int
	filename := "static/upload/0.png"
	savepath := "static/upload/0_600.png"
	err := go_image.ScaleF2F(filename, savepath, 600)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail1 成功")
}

func (con DefaultController) Thumbnail2(c *gin.Context) {
	filename := "static/upload/1.jpg"
	savepath := "static/upload/1_400.png"
	//按宽度和高度进行比例缩放，输入和输出都是文件
	err := go_image.ThumbnailF2F(filename, savepath, 400, 400)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail2 成功")
}

func (con DefaultController) Qrcode1(c *gin.Context) {
	var png []byte
	png, err := qrcode.Encode("https://github.com/yszgb", qrcode.Medium, 256)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	c.String(200, string(png)) // 页面上显示二维码
}

func (con DefaultController) Qrcode2(c *gin.Context) {
	savepath := "static/upload/qrcode.png"
	err := qrcode.WriteFile("https://github.com/yszgb", qrcode.Medium, 556, savepath)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	file, _ := os.ReadFile(savepath)
	c.String(200, string(file))
}
