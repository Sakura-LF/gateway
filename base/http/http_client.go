package http

import (
	"fmt"
	"io"
	"net/http"
)

func HttpClient() {
	// 1.创建客户端
	client := &http.Client{}
	// 2.创建请求
	req, err := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	if err != nil {
		panic(err)
		return
	}
	// 3.发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 4.处理服务器响应resp
	all, err := io.ReadAll(resp.Body)
	fmt.Println(string(all))
}
