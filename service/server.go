package service

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:1999")

	defer listener.Close()

	fmt.Println("server is started, ready to receive connection!")

	connection, _ := listener.Accept()
	defer connection.Close()

	fmt.Println("local address in server: ", connection.LocalAddr())
	connection.Write([]byte("iam server, we are received your data!"))

	fmt.Println("remote address in server: ", connection.RemoteAddr())
	buffer := make([]byte, 1024)
	n, _ := connection.Read(buffer)
	time.Sleep(10 * time.Second)
	fmt.Println("data is received from client: ", buffer[:n])

}

//func main() {
//	listener, err := net.Listen("tcp", "127.0.0.1:1999")
//	if err != nil {
//		panic("can not create a listener, error:" + err.Error())
//	}
//
//	defer listener.Close()
//
//	fmt.Println("Server started, ready to receive connection")
//
//	var i = 1
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			panic("error in accepting connection, error:" + err.Error())
//		}
//
//		go handleConnection(conn, i)
//		i++
//	}
//}
//
//func handleConnection(conn net.Conn, i int) {
//
//	fmt.Println("A new connection established number assigned:", i)
//
//	defer conn.Close()
//
//	buffer := make([]byte, 1024)
//	for {
//		n, err := conn.Read(buffer)
//		if err != nil {
//			fmt.Println("error in reading the connection, error:" + err.Error())
//			return
//		}
//		fmt.Printf("Client %d: %s\n", i, string(buffer[:n]))
//	}
//}
