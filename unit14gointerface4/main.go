package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var w io.Writer

func main() {
	// 类型断言
	// var w io.Writer
	w = os.Stdout
	// f := w.(*os.File)
	// fmt.Print(f)
	// c:=w.(*bytes.Buffer)
	c, ok := w.(*bytes.Buffer)
	if ok {
		fmt.Println("断言成功,值为：", c)
	} else {
		fmt.Println("断言失败")
	}
	r, err := os.Open("/no/such/file")
	// r.Read()
	// r.Read()
	// if err != nil {
	// 	fmt.Print(err) //open /no/such/file: no such file or directory
	// }
	fmt.Println(err)                //open /no/such/file: no such file or directory
	fmt.Println(os.IsNotExist(err)) // true
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	}
	// byteTest,err:=io.ReadAll(r)
	// io.Writer
}
