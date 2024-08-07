package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行参数
	var name string
	var age int
	var married bool

	// 给命令行参数绑定对应的变量
	flag.StringVar(&name, "name", "", "Your name")
	flag.IntVar(&age, "age", 0, "Your age")
	flag.BoolVar(&married, "married", false, "Are you married?")

	// 解析命令行参数
	flag.Parse()

	// 输出解析结果
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Married:", married)
}