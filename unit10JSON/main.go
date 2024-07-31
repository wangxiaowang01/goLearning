package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

// Factorial 递归计算阶乘
func Factorial(n int) int {
	// 基准情形
	if n == 0 {
		return 1
	}
	// 递归步骤
	return n * Factorial(n-1)
}

func main() {

	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Printf("json err %v\n", err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(Factorial(5))

}
