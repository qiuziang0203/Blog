package utils

import (
	"crypto/md5"
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取当前时间戳(毫秒)
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前时间戳(纳秒)
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// 获取当前日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// md5加密
func Md5(str string) string {
	//data := []byte(str)
	//return fmt.Sprintf("%x\n", md5.Sum(data))

	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 把字符串解析成html
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

// 表示把string字符串转换成int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// 表示把string字符串转换成Float
func Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

// 表示把int转换成string字符串
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}

func Upload(c *app.RequestContext) (string, error) {
	//1.获取上传文件
	file, err := c.FormFile("file")
	//判断上传文件上否存在
	if err != nil { //说明上传文件不存在
		return "", nil
	}
	//2.获取后缀名,判断后缀是否正确: .jpg,.png,.gif,.jpeg
	extName := path.Ext(file.Filename)
	//设置后缀map
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	//判断后缀是否合法
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀名不合法")
	}
	file.Filename = strconv.FormatInt(time.Now().UnixNano(), 10) + extName
	//3.创建图片保存目录 ./static/upload/20230203
	//获取日期
	day := GetDay()
	//拼接目录
	dir := "./static/" + day
	//创建目录:MkdirAll 目录不存在,会一次性创建多层
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		return "", err
	}
	//4.生成文件名称和文件保存目录: models.GetUnixNano() 获取时间戳(int64) 纳秒:防止速度过快而上传图片失败; strconv.FormatInt() 把时间戳(int64)转换成字符串
	filename := strconv.FormatInt(GetUnixNano(), 10) + extName
	//5.执行上传
	dst := path.Join(dir, filename)
	//上传文件到指定目录
	c.SaveUploadedFile(file, dst)
	return dst, nil
}
