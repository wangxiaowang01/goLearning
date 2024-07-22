package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Go!")
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}
}
