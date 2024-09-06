package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Model 封装公共的方法，供 Controller / 模板文件 / main.go 使用
// 供模板文件使用，需要在加载模板上方注册
// r := gin.Default()
// r.SetFuncMap(template.FuncMap{
// 	"unixToDate": models.UnixToDate,
// })

// 时间戳转换为日期
func UnixToDate(timestamp int) string {
	// fmt.Println("打印时间戳：", timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	// Interprets a time as UTC
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取时间
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// md5 加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// string 转 int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// int 转 string
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}

// 上传图片到指定目录
func UploadImg(c *gin.Context, picName string) (string, error) {
	// 1. 获取上传的文件
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	// 2. 获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}

	// 3. 创建图片保存目录  static/upload/20210624
	day := GetDay()
	dir := "./static/upload/" + day
	err1 := os.MkdirAll(dir, 0666)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}

	// 4. 生成文件名称和文件保存的目录   111111111111.jpeg
	fileName := strconv.FormatInt(GetUnix(), 10) + extName
	fmt.Println("保存图片到：",fileName)

	// 5. 执行上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}
