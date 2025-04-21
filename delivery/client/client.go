package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/delivery/deliveryparam"
	"gocasts.ir/go-fundamentals/todo-cli/pkg"
	"log"
	"net"
	"os"
	"strconv"
)

var authenticatedUserId int

func sendCommand(command string, connection net.Conn) {

	commandRequest := deliveryparam.NewCommand(command)
	serializedCommandRequest, mErr := json.Marshal(commandRequest)
	if mErr != nil {
		log.Fatalf("can't marshal data command login: %v", mErr)
	}

	if _, wErr := connection.Write(serializedCommandRequest); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}
}

func login(command string, connection net.Conn) {
	fmt.Println("login process...")

	fmt.Print("enter email: ")
	var email string = pkg.ReadInput()

	fmt.Print("enter password: ")
	var password string = pkg.ReadInput()

	sendCommand(command, connection)

	requestUser := deliveryparam.NewUserRequest(email, password)
	serializedUser, mErr := json.Marshal(requestUser)
	if mErr != nil {
		log.Fatalf("can't marshal data command login: %v", mErr)
	}

	if _, wErr := connection.Write(serializedUser); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawResponse = make([]byte, 1024)
	numberOfReadBytes, rErr := connection.Read(rawResponse)
	if rErr != nil {
		log.Fatalf("can't read data from connection, error: %v", rErr)
	}

	authenticatedUserId, _ = strconv.Atoi(string(rawResponse[:numberOfReadBytes]))

	if authenticatedUserId == 0 {
		fmt.Print("invalid email or password!\n")
	}
}

func runCommand(command string, connection net.Conn) {
	if command != "login-user" && command != "register-user" && command != "exit" && authenticatedUserId == 0 {
		//ToDo login function

		return
	}
	switch command {
	case "login-user":
		login(command, connection)
	case "register-user":
		// ToDo registeredUser
	case "new-category":
		// ToDo newCategory
	case "list-category":
		// ToDo list category
	case "new-task":
		// ToDo newTask
	case "list-task":
		// ToDo list task
	case "tasks-date":
		// ToDo tasksByDate
	case "exit":
		sendCommand(command, connection)
		os.Exit(0)
	default:
		fmt.Println("invalid command input!")
	}
}

func main() {

	var command string
	var ipAddr string
	flag.StringVar(&command, "command", "no-command", "command to run")
	flag.StringVar(&ipAddr, "ip", "127.0.0.1:1999", "ip address for connect to server")
	flag.Parse()

	connection, dErr := net.Dial("tcp", ipAddr)
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

	// load data from storage
	for {
		runCommand(command, connection)

		fmt.Print("please enter another command: ")
		command = pkg.ReadInput()
	}

	// create one request and serialized data for sent to server
	//request := deliveryparam.NewTaskRequest(message, "test title", "test time due date", 1999)
	//serializedData, mErr := json.Marshal(request)
	//if mErr != nil {
	//	log.Fatalf("can't serialized data, error: %v\n", mErr)
	//}
	//
	//send serialized data to server
	//if _, wErr := connection.Write(serializedData); wErr != nil {
	//	log.Fatalf("can't write data to connection: %s\n Error: %v\n", connection.RemoteAddr(), wErr)
	//}
	//
	//var responseData = make([]byte, 1024)
	//numberOfByteRead, rErr := connection.Read(responseData)
	//if rErr != nil {
	//	log.Printf("can't read data from connection, error: %v", rErr)
	//}
	//fmt.Println("\nResponse, this task is create: ", string(responseData[:numberOfByteRead]))

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
