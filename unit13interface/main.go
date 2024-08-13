package main

import "fmt"

// 定义一个接口 animal
type animal interface {
	say()
}

type cat struct {
	name string
}

// cat这个结构体实现了say方法
func (c cat) say() {
	fmt.Printf("%v叫了,喵喵喵\n", c.name)
}

type dog struct {
	name string
}

// dog这个结构体实现了say方法
func (d dog) say() {
	fmt.Printf("%v叫了,旺旺旺\n", d.name)
}

func aaa(a animal) {
	a.say()
}

func main() {

	c := cat{"蓝猫1"}
	// d := dog{"哮 天犬"}

	// c.say()
	// d.say()
	var a animal
	// a.say()

	// fmt.Println(a)        // <nil>
	// fmt.Printf("%T\n", a) // <nil>  因为a是没有意义的1
	a = c
	// fmt.Println(a)        // {蓝猫}
	// fmt.Printf("%T\n", a) // main.cat
	a.say()
	// aaa(c)
	// aaa(d)
	// aaa(a)

}
