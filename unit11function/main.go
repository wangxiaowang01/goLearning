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

// sum 函数接受任意数量的整数参数并返回它们的总和
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

type Point struct {
	x uint64
	y uint64
}

func (p Point) ScaleBy(factor uint64) {
	p.x = factor
}
func (p *Point) ScaleBy2(factor uint64) {
	p.x = factor
}
func main() {

	fmt.Println(Factorial(5))
	f := squares()               // squares的返回值是一个int类型返回值的匿名函数fun
	fmt.Println(f())             // "1"
	fmt.Println(f())             // "4"
	fmt.Println(f())             // "9"
	fmt.Println(f())             // "16"
	fmt.Println(sum(1, 2, 3))    // 输出：6
	fmt.Println(sum(4, 5, 6, 7)) // 输出：22
	fmt.Println(sum())           // 输出：0，没有参数时返回0

	testPoint := Point{
		1,
		2,
	}
	fmt.Printf("修改前的1  x= %v\n", testPoint.x)
	testPoint.ScaleBy(3)
	fmt.Printf("使用值方法ScaleBy修改后的x = %v\n", testPoint.x)
	testPoint.ScaleBy2(3)
	fmt.Printf("使用指针方法ScaleBy修改后的x= %v\n", testPoint.x)

}
