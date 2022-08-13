package file

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

// 检查文件是否存在
func IsExist(file string) bool {
	_, err := os.Stat(file)
	// 不存在
	if err != nil {
		// 错误中的信息指示存在与否，这里检测不存在的信息
		if os.IsNotExist(err) {
			return false
		}
		// 打印错误
		log.Fatal(err)
		return false
	}
	// 存在
	return true
}

// 读取文件内容
func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// 写入文件
func WriteFile(file string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(file, data, perm)
}

// 创建文件
func CreateFile(file string) (*os.File, error) {
	return os.Create(file)
}

// 重命名
func Rename(file, newName string) error {
	return os.Rename(file, newName)
}

// 删除文件
func RemoveFile(file string) error {
	return os.Remove(file)
}
