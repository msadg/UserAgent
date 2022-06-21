package UserAgent

import (
	"math/rand"
	"time"

	"github.com/msadg/UserAgent/comm"
	"github.com/msadg/UserAgent/entity"
)

// 全局数据
var data *entity.UaData

// 初始化函数
// 从文本中读取信息，并写入一个变量中
func init() {
	rand.Seed(time.Now().Unix())
	uas, err := comm.LoadData() // 加载数据
	checkError(err)

	data = &entity.UaData{}
	data.Browsers = uas.Bs

	data.Res = make([]string, 0)
	data.Names = make([]string, 0)
	// 填充 UserAgent list
	for k, v := range data.Browsers {
		data.Names = append(data.Names, k)
		for _, info := range v {
			data.Res = append(data.Res, info)
		}
	}
}

// 所有浏览器列表
func ListBrowsers() []string {
	if data.Browsers == nil {
		return nil
	}
	if len(data.Browsers) > 0 {
		return data.Names
	}
	return nil
}

// 随机获取一条 User-Agent
func Rand() string {
	if data.Res == nil {
		return ""
	}
	if length := len(data.Res); length > 0 {
		index := randInt(data.Res)
		if index == length {
			index -= 1
		}
		if index < 0 {
			index = 0
		}
		return data.Res[index]
	}
	return ""
}

// 浏览器列表名字中随机获取一条 User-Agent
func RandBs(name string) string {
	if name == "" {
		return ""
	}
	bsData := data.Browsers[name]
	if len(bsData) > 0 {
		return bsData[randInt(bsData)]
	}
	return ""
}

// 获取某个浏览器所属的所有 User-Agent
func BrowserAll(name string) []string {
	if data.Browsers == nil {
		return nil
	}
	return data.Browsers[name]
}

// 获取所有的 User-Agent
func All() []string {
	return data.Res
}

// 根据数组长度获取随机数
func randInt(d []string) int {
	length := len(d)
	index := rand.Intn(length)
	return index
}

// 错误检查
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
