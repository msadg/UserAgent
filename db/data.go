package db

import "github.com/msadg/UserAgent/io"

// 获得文件内容
func Load(file string) ([]byte, error) {
	return io.ReadFile(file)
}

// 更新资料
func Update(file string, data []byte) error {
	return io.WriteFile(file, data, 0777)
}
