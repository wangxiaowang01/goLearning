package unit17gotest

import "fmt"

func AddFun(x int64, y int64) int64 {
	z := x + y
	fmt.Printf(" x + y = z = %v\n", z)
	if z == 3 {
		fmt.Print("hhhh")
	}
	return z
}
