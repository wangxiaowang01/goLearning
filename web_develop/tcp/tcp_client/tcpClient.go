package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const ipAndProt = "127.0.0.1:8080"

func main() {
	// 1. 建立连接,拨号，dial就是拨号意思
	conn, err := net.Dial("tcp", ipAndProt)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		fmt.Println("执行defer语句 关闭连接")
		conn.Close()
	}()
	// 2。发送消息
	reader := bufio.NewReader(os.Stdin)
	for {
		// 先发消息
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		conn.Write([]byte(input))
		// 3、接收回复
		buf := make([]byte, 100)
		conn.Read(buf)
		fmt.Println("收到:", string(buf))

	}

}
