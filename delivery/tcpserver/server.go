package main

import (
	"encoding/json"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/delivery/deliveryparam"
	"gocasts.ir/go-fundamentals/todo-cli/repository/memoryStore"
	"gocasts.ir/go-fundamentals/todo-cli/service/task"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1999")
	if err != nil {
		panic("can not create a listener, error:" + err.Error())
	}

	defer func() {
		if cErr := listener.Close(); cErr != nil {
			log.Println("listener is not closed", cErr.Error())
		}
	}()

	fmt.Printf("Server started, ready to receive connection on %s\n", listener.Addr())

	var i = 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error in accepting connection, error:" + err.Error())
		}

		go handleConnection(conn, i)
		i++
	}
}

func handleConnection(connection net.Conn, i int) {

	fmt.Println("A new connection established number assigned:", i)

	// definition buffer for read data from client
	buffer := make([]byte, 1024)
	//for {

	numberOfReadBytes, rErr := connection.Read(buffer)
	if rErr != nil {
		log.Println("error in reading the buffer connection, error:", rErr)

		//break
	}

	// deserializing data for use in server
	req := deliveryparam.NewRequest("")
	if uErr := json.Unmarshal(buffer[:numberOfReadBytes], req); uErr != nil {
		log.Println("bad request", uErr)

		//continue
	}

	runCommand(connection, req)

	//}
}

func runCommand(connection net.Conn, request *deliveryparam.Request) {

	//defer connection.Close()
	defer func() {
		if cErr := connection.Close(); cErr != nil {
			log.Println("connection is not closed", cErr.Error())
		}
		log.Println("connection close ...")
	}()

	taskMemoryRepo := memoryStore.NewTaskMemory()
	categoryMemoryRepo := memoryStore.NewCategoryMemory()
	taskService := task.NewService(taskMemoryRepo, categoryMemoryRepo)

	switch request.GetCommand() {
	case "create-task":

		responseCreatedTask, cErr := taskService.CreateTask(task.NewRequest("title", "dueDate", 0, 0))
		if cErr != nil {
			if _, wErr := connection.Write([]byte(cErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)

				return
				//continue
			}
			return
		}

		data, mErr := json.Marshal(responseCreatedTask.GetTask())
		if mErr != nil {
			log.Println("can't marshal responseCreatedTask:", mErr)

			//continue
		}

		if _, wErr := connection.Write(data); wErr != nil {
			log.Println("can't write data to connection", wErr)

			//continue
		}
	case "list-task":
		responseListTask, lErr := taskService.ListTask(task.NewListRequest(1999))
		if lErr != nil {
			if _, wErr := connection.Write([]byte(lErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)
			}
		}

		data, mErr := json.Marshal(responseListTask.GetTasks())
		if mErr != nil {
			log.Println("can't marshal responseCreatedTask:", mErr)

			//continue
		}

		if _, wErr := connection.Write(data); wErr != nil {
			log.Println("can't write data to connection", wErr)

			//continue
		}
	}
}
