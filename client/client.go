package main

import (
	"encoding/json"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/delivery/deliveryparam"
	"log"
	"net"
	"os"
)

func main() {

	fmt.Println("command", os.Args[0])

	var message = "default message"

	if len(os.Args) > 1 {
		message = os.Args[1]
	}

	// create one connection to server
	connection, dErr := net.Dial("tcp", "127.0.0.1:1999")
	if dErr != nil {
		panic("error established connection, make sure the server is up")
	}

	// defer close this connection
	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatalf("close connection is error: %s\n", err.Error())
		}
	}()

	fmt.Printf("Connected to a tcp server on %s\n", connection.RemoteAddr())

	// create one request and serialized data for sent to server
	request := deliveryparam.NewRequest(message)
	serializedData, mErr := json.Marshal(request)
	if mErr != nil {
		log.Fatalf("can't serialized data, error: %v\n", mErr)
	}

	// send serialized data to server
	if _, wErr := connection.Write(serializedData); wErr != nil {
		log.Fatalf("can't write data to connection: %s\n Error: %v\n", connection.RemoteAddr(), wErr)
	}

	var responseData = make([]byte, 1024)
	numberOfByteRead, rErr := connection.Read(responseData)
	if rErr != nil {
		log.Printf("can't read data from connection, error: %v", rErr)
	}
	fmt.Println("\nResponse, this task is create: ", string(responseData[:numberOfByteRead]))

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
	//	fmt.Fprint(connection, message)
	//}
}
