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
	Json_Agent_File = "../../UserAgent/data/UserAgent.json" // ChinaArea
	// Json_Agent_File = "../../data/UserAgent.json" // test
	// Json_Agent_File = "../UserAgent/data/UserAgent.json"

	// 链接地址
	// https://fake-useragent.herokuapp.com/browsers/0.1.11
	User_Agent_Url = "https://fake-useragent.herokuapp.com/browsers/0.1.11"
)

// 从网络更新 UserAgent
func Update() error {

	data, err := net.Get(User_Agent_Url)
	if err != nil {
		return err
	}

	// 更新数据
	return db.Update(Json_Agent_File, data)
}

// 加载数据
func LoadData() (*entity.UasJsonData, error) {

	// pwd, err := os.Getwd()
	// checkError(err)
	// log.Println(pwd)

	// 加载数据
	ds, err := db.Load(Json_Agent_File)
	if err != nil {
		return nil, err
	}

	if ds == nil {
		return nil, errors.New("获取数据出错")
	}

	if len(ds) == 0 {
		return nil, errors.New("获取不到任何内容")
	}

	uasJsonData := &entity.UasJsonData{}

	// 解析成 JSON 数据
	err = json.Unmarshal(ds, uasJsonData)
	if err != nil {
		return nil, err
	}
	return uasJsonData, nil
}
