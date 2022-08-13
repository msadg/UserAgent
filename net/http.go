package net

import (
	"io/ioutil"
	"net/http"
)

// 请求资源
func Get(url string) ([]byte, error) {
	// 设置请求
	req, err := http.NewRequest(http.MethodGet, url, nil)
	ua := ""
	checkError(err)
	req.Header.Set("userAgent", ua)
	checkError(err)

	// 建立链接
	cli := &http.Client{}
	resp, err := cli.Do(req)
	checkError(err)
	defer resp.Body.Close()

	// 读取数据
	return ioutil.ReadAll(resp.Body)
}

// 检查错误
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
