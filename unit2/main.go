package main

import "fmt"

var q = f()

func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func main() {
	v := 1
	incr(&v)
	fmt.Printf("v2=%v\n", v)
	fmt.Printf("q=%v\n", q)
	fmt.Println(f() == f()) // "false"

	fmt.Println("hhh")
	x := 1
	fmt.Printf("x = %v\n", x) // x = 1

	p := &x // p是什么？ p就是x的指针，指针就是这个x在内存中的地址，是一个物理存在的位置

	fmt.Printf("p = %v\n", p)   // p = 0xc0000120d0 这里的p是什么？ 是x = 1在内存中的地址，是一个物理存在的位置
	fmt.Printf("*p = %v\n", *p) // *p = 1 这里*p是什么？ 表示的是p这个地址上的值 也就是1
	*p = 2                      // 等于号= 就是赋值，现在我们把*p也就是p指针上的值重新赋值成了2，所以，p这个指针的位置0x什么的没有变，但是值变了 所以x的值也就变了，
	// 当您执行 *p = 2 时，您实际上修改了 p 指针所指向的内存地址上的值，将其更新为 2。由于 p 指向的是 x 的地址，因此这个更新也会影响到 x 的值。
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)
	fmt.Printf("x = %v\n", x)
	// 在这个过程中，x的值是发生了改变，但是x的指针是永远没有变的
	// 所以，通过修改 *p，也就是通过指针间接地修改了 x 的值。x 的指针 p 的物理位置没有改变，它仍然指向 x 在内存中的地址，但是 x 的值发生了改变。
	var z, y int
	fmt.Println(&z == &z, &z == &y, &z == nil) // "true false false"

}
