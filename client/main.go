package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// 连接到服务端
	addr := "localhost:8080"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}
	// 连接到服务端
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		// 向服务端发送数据
		message := "Hello from client"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending data to server:", err)
			return
		}

		// 接收服务端回传的数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Printf("Received from server: %s\n", string(buffer[:n]))

		// 每隔2秒发送一次消息
		time.Sleep(2 * time.Second)
	}
}
