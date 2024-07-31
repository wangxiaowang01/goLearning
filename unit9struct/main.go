package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Address string
}

func ReturnStudent(student Student) string {
	return student.Address
}
func EditStudent(student Student) {
	student.Age = 100
	fmt.Printf("函数内s.Age= %v\n", student.Age)

}
func PointEditStudent(student *Student) {
	student.Age = 100
	fmt.Printf("函数内s.Age= %v\n", student.Age)

}
func main() {
	s := &Student{
		Name:    "alice",
		Age:     10,
		Address: "beijing",
	}
	fmt.Printf("s = %#v\n", s)
	fmt.Printf("s.Age= %v\n", s.Age)
	// EditStudent(s)
	fmt.Printf("函数外s.Age= %v\n", s.Age)
	PointEditStudent(s)
	fmt.Printf("第二次函数外s.Age= %v\n", s.Age)

}
