package main

import "fmt"


func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	var numSlice = []string{}
	numSlice = append(numSlice, "1")
	numSlice = append(numSlice, "a")
	numSlice = append(numSlice, "hello")
	numSlice = append(numSlice, "你好")
	fmt.Printf("numSlice %v\n", numSlice)
	// numSlice = numSlice[2:]
	// fmt.Printf("numSlice %v\n", numSlice)

	// numSlice = append(numSlice[:0], numSlice[2:]...)
	// fmt.Printf("numSlice %v\n", numSlice)
	// numSlice = numSlice[:copy(numSlice, numSlice[2:])]
	// fmt.Printf("numSlice %v\n", numSlice)
	// fmt.Printf("numSlice %v\n", numSlice[:20])
	// newSlice := numSlice[:20]
	// fmt.Printf("newSlice %v\n", newSlice)
	numSlice1 := []string{"1", "a", "hello", "你好"}
	numSlice2 := []string{"1", "a", "hello", "你好"}
	
	

	fmt.Printf("判断是否相等%v\n", equal(numSlice1, numSlice2))
	
}
