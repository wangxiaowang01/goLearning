package main

import (
	"fmt"
	// "math/big"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 28,
	}
	fmt.Println(ages["bob"])
	ages["bob"] = ages["bob"] + 1 // happy birthday!

	fmt.Println(ages["bob"])
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// a:=big.NewInt(10000000)
	// b:=big.NewInt(20000000)
	// result := new(big.Int)
	// c=a+b

}
