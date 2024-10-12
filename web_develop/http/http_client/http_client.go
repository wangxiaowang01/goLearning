package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.shicicn.com/zuozhe/129.html")
	if err != nil {
		fmt.Printf("err = %v", err)
	}
	// body, _ := ioutil.ReadAll(response.Body)
	body, _ := io.ReadAll(response.Body)
	fmt.Printf("body = %s", string(body))
	response.Body.Close()
}
