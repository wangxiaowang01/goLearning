package main

import "fmt"

func Factorial(n int) int {
	// 基准情形
	if n == 0 {
		return 1
	}
	// 递归步骤
	return n * Factorial(n-1)
}
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {

	fmt.Println(Factorial(5))
	f := squares()   // squares的返回值是一个int类型返回值的匿名函数fun
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

}
