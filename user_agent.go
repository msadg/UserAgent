package UserAgent

import (
	"math/rand"
	"time"

	"github.com/msadg/UserAgent/comm"
	"github.com/msadg/UserAgent/entity"
)

// 加载
// 从文本中读取信息，并写入一个变量中
func load() entity.UaData {

	var data entity.UaData

	rand.Seed(time.Now().Unix())

	uas, err := comm.LoadData() // 加载数据
	checkError(err)

	data = entity.UaData{}
	data.Browsers = uas.Bs // 浏览器分类

	data.Res = make([]string, 0)
	data.Names = make([]string, 0)
	// 填充 UserAgent list
	for k, v := range data.Browsers {
		data.Names = append(data.Names, k)
		for _, info := range v {
			data.Res = append(data.Res, info)
		}
	}
	return data
}

// 所有浏览器列表
func ListBrowsers() []string {
	data := load()
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
	data := load()
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
	data := load()
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
	data := load()
	if data.Browsers == nil {
		return nil
	}
	return data.Browsers[name]
}

// 获取所有的 User-Agent
func All() []string {
	data := load()
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
