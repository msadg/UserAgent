package comm

import (
	"encoding/json"
	"errors"

	"github.com/msadg/UserAgent/db"
	"github.com/msadg/UserAgent/entity"
	"github.com/msadg/UserAgent/net"
)

const (
	// 文件位置
	// Json_Agent_File = "../../data/UserAgent.json"
	Json_Agent_File = "../UserAgent/data/UserAgent.json"
	// 链接地址
	User_Agent_Url = "https://fake-useragent.herokuapp.com/browsers/0.1.11"
)

// https://fake-useragent.herokuapp.com/browsers/0.1.11
// 从网络更新 UserAgent
func Update() error {
	data, err := net.Get(User_Agent_Url)
	checkError(err)

	// 更新数据
	return db.Update(Json_Agent_File, data)
}

// 加载数据
func LoadData() (*entity.UasJsonData, error) {
	// 加载数据
	ds, err := db.Load(Json_Agent_File)
	checkError(err)

	if ds == nil {
		return nil, errors.New("获取数据出错")
	}

	if len(ds) == 0 {
		return nil, errors.New("获取数据出错")
	}

	uasJsonData := &entity.UasJsonData{}

	// 解析成 JSON 数据
	err = json.Unmarshal(ds, uasJsonData)
	return uasJsonData, err
}

// 检查错误
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
