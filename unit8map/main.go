package main

import "fmt"

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

}
