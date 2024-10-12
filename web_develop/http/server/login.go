package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("匿名函数写响应"))
	})

	fmt.Println("服务启动 ")
	// http.ListenAndServe("127.0.0.1:8080", nil)
	http.ListenAndServe(":8080", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	// 客户端要http://localhost:8080/login?username=zhangsan&password=123456 这样请求才对
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username == "zhangsan" && password == "123456" {
		w.Write([]byte("登陆成功"))
	} else {
		w.Write([]byte("登陆错误，帐号密码不匹配"))

	}
}
