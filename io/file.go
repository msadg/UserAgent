package io

import (
	"io/fs"
	"io/ioutil"
)

// 读取文件内容
func ReadFile(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

// 写入文件
func WriteFile(file string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(file, data, perm)
}
