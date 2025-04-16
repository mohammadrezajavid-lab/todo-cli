package main

import (
	"bufio"
	"net"
	"os"
	"time"
)

func main() {
	connection, _ := net.Dial("tcp", "127.0.0.1:1999")
	defer connection.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputData := scanner.Text()
	connection.Write([]byte(inputData))
	time.Sleep(50 * time.Second)

	//conn, err := net.Dial("tcp", "127.0.0.1:1999")
	//if err != nil {
	//	panic("error established connection, make sure the server is up")
	//}
	//
	//defer func() {
	//	if err := conn.Close(); err != nil {
	//		_ = fmt.Errorf("close connection is error: %s", err.Error())
	//	}
	//}()
	//
	//log.Printf("Connected to a tcp server on %s\n", conn.RemoteAddr())
	//
	//fmt.Println("Write something to send to server")
	//
	//var message string
	//scanner := bufio.NewScanner(os.Stdin)
	//for {
	//	fmt.Print(">> ")
	//
	//	scanner.Scan()
	//	message = scanner.Text()
	//
	//	if message == "exit" {
	//		return
	//	}
	//
	//	fmt.Fprint(conn, message)
	//}
}
