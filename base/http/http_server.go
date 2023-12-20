package http

import (
	"fmt"
	"log"
	"net/http"
)

func HttpServer() {
	// 一: 设置不同路由对应的不同处理器
	// 第一种方式: HandleFunc
	http.HandleFunc("/ping", HandlerPing)

	// 第二种方式: InfoHandleFunc
	// 需要结构体,以及实现方法ServerHTTP
	Info := InfoHandle{Info: "Hello , Sakura"}
	http.Handle("/info", Info)

	// 二:启动监听
	addr := ":8080"
	log.Println("HTTP 服务器正在监听: ", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}

func HandlerPing(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "pong")
}

// Handle第二个参数需要结构体,并且实现ServerHTTP
type InfoHandle struct {
	Info string
}

func (info InfoHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, info.Info)
}
