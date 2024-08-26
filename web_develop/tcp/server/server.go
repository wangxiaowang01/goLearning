package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// 1.准备手机卡
const ipAndProt = "127.0.0.1:8080"

func process(conn net.Conn) {
	// 4.电话来了，接听
	fmt.Println("电话来了，接听，了")

	defer func() {
		fmt.Println("执行defer语句 关闭连接")
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		// req := make([]byte, 100)
		// reader.Read(req)
		// 这里传一个\n是因为要判别，客户端如果发来的消息是\n结尾的话就读取
		reqStr, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connnection closed")
			} else {
				log.Println(err)
			}
			return
		}
		fmt.Println("服务端接收到的消息是：", reqStr)
		// 5. 回消息，
		fmt.Println("回消息")
		// 再发消息
		// strings
		conn.Write([]byte(strings.ToUpper(reqStr)))
		// 5. 挂电话
		// fmt.Println("挂电话了")
		// conn.Close()
		// fmt.Println("挂电话了")
	}

}

func main() {

	// 2.插卡开机，绑定电话卡
	// net.ListenTCP("tcp",ipAndProt)
	fmt.Println("建立一个通信，了")

	listen, err := net.Listen("tcp", ipAndProt)
	if err != nil {
		fmt.Print(err)
	}
	for {
		// 3.等电话
		fmt.Println("等电话，了")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print(err)
		}
		go process(conn)
	}

	// process(conn)

}
