package main

import (
	"encoding/json"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/delivery/deliveryparam"
	"gocasts.ir/go-fundamentals/todo-cli/repository/memoryStore"
	"gocasts.ir/go-fundamentals/todo-cli/service/task"
	"gocasts.ir/go-fundamentals/todo-cli/service/task/taskparam"
	"log"
	"net"
)

func main() {

	taskMemoryRepo := memoryStore.NewTaskMemory()
	categoryMemoryRepo := memoryStore.NewCategoryMemory()
	taskService := task.NewService(taskMemoryRepo, categoryMemoryRepo)

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

		go handleConnection(conn, taskService, i)
		i++
	}
}

func handleConnection(connection net.Conn, taskService *task.Service, i int) {

	//defer connection.Close()
	defer func() {
		if cErr := connection.Close(); cErr != nil {
			log.Println("connection is not closed", cErr.Error())
		}
		log.Println("connection close ...")
	}()

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
	req := deliveryparam.NewEmptyRequest()
	if uErr := json.Unmarshal(buffer[:numberOfReadBytes], req); uErr != nil {
		log.Println("bad request", uErr)

		//continue
	}

	runCommand(connection, taskService, req)

	//}
}

func runCommand(connection net.Conn, taskService *task.Service, request *deliveryparam.Request) {

	switch request.GetCommand() {
	case "create-task":

		newTask := request.GetTask()
		responseCreatedTask, cErr := taskService.CreateTask(taskparam.NewRequest(newTask.GetTitle(), newTask.GetDueDate(), newTask.GetCategoryId(), 0))
		if cErr != nil {
			if _, wErr := connection.Write([]byte(cErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)

				return
				//continue
			}
			return
		}

		data, mErr := json.Marshal(responseCreatedTask)
		if mErr != nil {
			log.Println("can't marshal responseCreatedTask:", mErr)

			//continue
		}

		if _, wErr := connection.Write(data); wErr != nil {
			log.Println("can't write data to connection", wErr)

			//continue
		}
	case "list-task":

		responseListTask, lErr := taskService.ListTask(taskparam.NewListRequest(0))
		if lErr != nil {
			if _, wErr := connection.Write([]byte(lErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)
			}
		}

		data, mErr := json.Marshal(responseListTask)
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
