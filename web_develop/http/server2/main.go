package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
func main() {
	http.ListenAndServe()
	m := MyHandler{}
	s := http.Server{
		Addr:    ":8080",
		Handler: &m,
	}
	s.ListenAndServe()
}
