package main

import (
	"fmt"

	"os"
)

func main() {
	fmt.Print("test")
	// var w io.Writer
	// w = os.Stdout
	// f, ok := w.(*os.File)
	// if ok != true {
	// 	fmt.Print(ok)
	// }

	_, err := os.Open("/test/test")
	// if err!=nil && os.IsNotExist(err){
	// 	fmt.Print("文件不存在")
	// }
	if err != nil {
		fmt.Print(err.Error())
	}
}
