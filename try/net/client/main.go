package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	/*
	   Dial:
	       在网络network上连接地址address，并返回一个Conn接口。
	       可用的网络类型有："tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"
	       对TCP和UDP网络，地址格式是host:port或[host]:port
	*/
	conn, err := net.Dial("tcp", "localhost:10001")
	if err != nil {
		fmt.Println("conn err :", err)
		return
	}
	fmt.Println("dial success")
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please enter a message? 'quit' exit")
		//读取消息
		input, _ := inputReader.ReadString('\n')
		msg := strings.Trim(input, "\r\n")
		//quit 退出
		if msg == "quit" {
			fmt.Println("quit")
			conn.Write([]byte("client quit "))
			return
		}
		_, err = conn.Write([]byte(msg))
	}

}
