package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		// 读取客户端发送的数据
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			return
		}
		fmt.Printf("Received from client: %s\n", string(buffer[:n]))

		// 向客户端回写数据
		_, err = conn.Write([]byte("Message received"))
		if err != nil {
			fmt.Println("Error sending data to client:", err)
			return
		}
	}
}

func main() {
	// 监听指定端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")
	for {
		// 等待并接受新的客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 处理客户端连接
		go handleConnection(conn)
	}
}
