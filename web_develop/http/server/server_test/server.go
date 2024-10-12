package server_test

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello 有请求了")
	w.Write([]byte("hello "))
}
func test() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("匿名函数写响应"))
	})

	fmt.Println("服务启动 ")
	// http.ListenAndServe("127.0.0.1:8080", nil)
	http.ListenAndServe(":8080", nil)
}
