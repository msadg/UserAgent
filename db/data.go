package db

import (
	"fmt"
	"strconv"
	"time"

	f "github.com/msadg/UserAgent/io/file"
)

// 获得文件内容
func Load(file string) ([]byte, error) {
	if !f.IsExist(file) {
		return nil, fmt.Errorf("没有【%s】这个文件，获取不到内容", file)
	}
	return f.ReadFile(file)
}

// 更新资料
func Update(file string, data []byte) error {
	// 如果文件存在
	if f.IsExist(file) {
		// // 先删除文件
		// err := f.RemoveFile(file)

		t := time.Now()
		now := t.Year() +
			int(t.Local().UTC().Month()) +
			t.Local().UTC().Day()
		// backup
		err := f.Rename(file, file+strconv.FormatUint(uint64(now), 0)+".bak")
		if err != nil {
			return err
		}
	}
	// 创建文件
	_, err := f.CreateFile(file)
	if err != nil {
		return err
	}
	// 写入文件
	return f.WriteFile(file, data, 0777)
}
