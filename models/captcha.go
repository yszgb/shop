package models

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 创建 store
// var store = base64Captcha.DefaultMemStore
// 配置 RedisStore，实现接口，类型为 base64Captcha.Store
var store base64Captcha.Store = RedisStore{}

// 获取验证码
func MakeCaptcha() (string, string, string, error) {
	var driver base64Captcha.Driver

	// 配置验证码信息，创建字符串验证码
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          2,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 创建实例，ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()

	// 创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)

	// Generate 生成随机 id、base64 图像字符串
	id, b64s, answer, err := captcha.Generate()
	return id, b64s, answer, err
}

// 验证验证码是否正确
func VerifyCaptcha(id string, VerifyValue string) bool {
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
