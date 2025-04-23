package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/constant"
	"gocasts.ir/go-fundamentals/todo-cli/delivery/deliveryparam"
	"gocasts.ir/go-fundamentals/todo-cli/pkg"
	"gocasts.ir/go-fundamentals/todo-cli/repository/memoryStore"
	"gocasts.ir/go-fundamentals/todo-cli/service/category"
	"gocasts.ir/go-fundamentals/todo-cli/service/category/categoryparam"
	"gocasts.ir/go-fundamentals/todo-cli/service/task"
	"gocasts.ir/go-fundamentals/todo-cli/service/task/taskparam"
	"gocasts.ir/go-fundamentals/todo-cli/service/user"
	"gocasts.ir/go-fundamentals/todo-cli/service/user/userparam"
	"log"
	"net"
	"os"
)

func init() {

	func() {
		fUsers, _ := os.OpenFile(constant.UsersFile, os.O_CREATE, constant.PermFile)
		fTasks, _ := os.OpenFile(constant.TasksFile, os.O_CREATE, constant.PermFile)
		fCategories, _ := os.OpenFile(constant.CategoriesFile, os.O_CREATE, constant.PermFile)

		_ = fUsers.Close()
		_ = fTasks.Close()
		_ = fCategories.Close()
	}()
}

func main() {

	var userMemoryRepo *memoryStore.UserMemory = memoryStore.NewUserMemory()
	var categoryMemoryRepo *memoryStore.CategoryMemory = memoryStore.NewCategoryMemory()
	var taskMemoryRepo *memoryStore.TaskMemory = memoryStore.NewTaskMemory()

	var userService *user.Service = user.NewService(userMemoryRepo)
	var categoryService *category.Service = category.NewService(categoryMemoryRepo)
	var taskService *task.Service = task.NewService(taskMemoryRepo, categoryMemoryRepo)

	var ipAddr string
	flag.StringVar(&ipAddr, "ip", "127.0.0.1:1999", "set your ip address for listen in server")
	flag.Parse()

	listener, err := net.Listen("tcp", ipAddr)
	if err != nil {
		panic("can not create a listener, error:" + err.Error())
	}

	defer func() {
		if cErr := listener.Close(); cErr != nil {
			log.Println("listener is not closed", cErr.Error())
		}
	}()

	fmt.Printf("Server started, ready to receive connection on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error in accepting connection, error:" + err.Error())
		}

		go handleConnection(conn, userService, categoryService, taskService)
	}
}

func handleConnection(connection net.Conn, userService *user.Service, categoryService *category.Service, taskService *task.Service) {

	defer func() {
		if cErr := connection.Close(); cErr != nil {
			log.Println("connection is not closed", cErr.Error())
		}
		log.Println("connection close ...")
	}()

	fmt.Println("A new connection established")

	for {
		var command string = getCommand(connection)
		if command == "exit" {
			break
		}
		runCommand(connection, command, userService, categoryService, taskService)
	}
}

func runCommand(connection net.Conn, command string, userService *user.Service, categoryService *category.Service, taskService *task.Service) {

	switch command {
	case "login-user":

		var rawUser = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(rawUser)
		if rErr != nil {
			log.Printf("can't read data user from client in login, error: %v", rErr)
		}

		requestUser := deliveryparam.NewUserRequest("", "")
		if uErr := json.Unmarshal(rawUser[:numberOfReadBytes], requestUser); uErr != nil {
			log.Printf("can't unmarshal user in login, error: %v", uErr)
		}

		var res, lErr = userService.Login(userparam.NewRequestUser(requestUser.GetEmail(), requestUser.GetPassword()))
		if lErr != nil {
			log.Printf("user not found, email: %s\nerror: %v", requestUser.GetEmail(), lErr)

			responseUser := deliveryparam.NewUserResponse(0)
			serializedResUser, mErr := json.Marshal(responseUser)
			if mErr != nil {
				log.Printf("can't marshal data  serializedResUser login: %v", mErr)
			}

			if _, wErr := connection.Write(serializedResUser); wErr != nil {
				log.Printf("can't write data to connection: %v", wErr)
			}
			return
		}
		responseUser := deliveryparam.NewUserResponse(res.GetUserId())
		serializedResUser, mErr := json.Marshal(responseUser)
		if mErr != nil {
			log.Printf("can't marshal data  serializedResUser login: %v", mErr)
		}

		if _, wErr := connection.Write(serializedResUser); wErr != nil {
			log.Printf("can't write data to connection: %v", wErr)
		}

		log.Printf("login user by email: %s", res.GetEmail())
	case "register-user":

		var rawUser = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(rawUser)
		if rErr != nil {
			log.Printf("can't read data user from client in register-user, error: %v", rErr)
		}
		requestUser := deliveryparam.NewRegisterUserRequest("", "", "")
		if uErr := json.Unmarshal(rawUser[:numberOfReadBytes], requestUser); uErr != nil {
			log.Printf("can't unmarshal user in register-user, error: %v", uErr)
		}

		var responseRegisterUser, registerUserErr = userService.RegisterUser(userparam.NewRequestRegisterUser(requestUser.GetName(), requestUser.GetEmail(), pkg.HashPassword(requestUser.GetPassword())))
		if registerUserErr != nil {
			log.Printf("user can't register, email: %s\nerror: %v", responseRegisterUser.GetEmail(), registerUserErr)

			responseUser := deliveryparam.NewRegisterUserResponse(responseRegisterUser.GetEmail(), registerUserErr)
			serializedResUser, mErr := json.Marshal(responseUser)
			if mErr != nil {
				log.Printf("can't marshal data  in register-user: %v", mErr)
			}

			if _, wErr := connection.Write(serializedResUser); wErr != nil {
				log.Printf("can't write data to connection: %v", wErr)
			}
			return
		}
		responseUser := deliveryparam.NewRegisterUserResponse(responseRegisterUser.GetEmail(), nil)
		serializedResUser, mErr := json.Marshal(responseUser)
		if mErr != nil {
			log.Printf("can't marshal data  serializedResponseUser register-user: %v", mErr)
		}

		if _, wErr := connection.Write(serializedResUser); wErr != nil {
			log.Printf("can't write data to connection: %v", wErr)
		}
		log.Printf("registered user by email: %v", responseRegisterUser.GetEmail())
	case "new-category":

		var rawCategory = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(rawCategory)
		if rErr != nil {
			log.Printf("can't read data user from client in new-category, error: %v", rErr)
		}
		categoryRequest := deliveryparam.NewCategoryRequest("", "", 0)
		if uErr := json.Unmarshal(rawCategory[:numberOfReadBytes], categoryRequest); uErr != nil {
			log.Printf("can't unmarshal user in new-category, error: %v", uErr)
		}

		createCategoryResponse, cErr := categoryService.CreateCategory(categoryparam.NewRequest(categoryRequest.GetTitle(), categoryRequest.GetColor(), categoryRequest.GetAuthenticatedUserId()))
		if cErr != nil {
			log.Printf("can't create category \nerror: %v", cErr)

			responseCreateCategory := deliveryparam.NewCategoryResponse("", 0, cErr)
			serializedData, mErr := json.Marshal(responseCreateCategory)
			if mErr != nil {
				log.Printf("can't marshal data  in new-category: %v", mErr)
			}

			if _, wErr := connection.Write(serializedData); wErr != nil {
				log.Printf("can't write data to connection: %v", wErr)
			}
			return
		}

		responseCreateCategory := deliveryparam.NewCategoryResponse(createCategoryResponse.GetCategory().GetTitle(), createCategoryResponse.GetCategory().GetId(), nil)
		serializedData, mErr := json.Marshal(responseCreateCategory)
		if mErr != nil {
			log.Printf("can't marshal data  in new-category: %v", mErr)
		}

		if _, wErr := connection.Write(serializedData); wErr != nil {
			log.Printf("can't write data to connection: %v", wErr)
		}

		log.Printf("category [%s] by id: [%d] is create!\n", responseCreateCategory.GetTitle(), responseCreateCategory.GetCategoryId())
	case "list-category":

		var catListReq = deliveryparam.NewCategoryListRequest(0)
		var buffer = make([]byte, 500)
		numberOfReadBytes, rErr := connection.Read(buffer)
		if rErr != nil {
			log.Printf("can't read data from connection in list-category, error: %v", rErr)
		}
		if uErr := json.Unmarshal(buffer[:numberOfReadBytes], catListReq); uErr != nil {
			log.Printf("can't unmarshal data in list-category response %v", uErr)
		}

		responseListCategory, lErr := categoryService.ListCategory(categoryparam.NewListRequest(catListReq.GetAuthenticatedUserId()))
		if lErr != nil {
			if _, wErr := connection.Write([]byte(lErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)
			}
		}
		data, mErr := json.Marshal(responseListCategory)
		if mErr != nil {
			log.Println("can't marshal responseCreatedTask:", mErr)
		}
		if _, wErr := connection.Write(data); wErr != nil {
			log.Println("can't write data to connection", wErr)
		}
	case "new-task":

		var rawTask = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(rawTask)
		if rErr != nil {
			log.Printf("can't read data user from client in new-task, error: %v", rErr)
		}
		newTaskReq := deliveryparam.NewTaskRequest("", "", 0, 0)
		if uErr := json.Unmarshal(rawTask[:numberOfReadBytes], newTaskReq); uErr != nil {
			log.Printf("can't unmarshal user in new-task, error: %v", uErr)
		}

		newTaskRes, cErr := taskService.CreateTask(taskparam.NewRequest(newTaskReq.GetTitle(), newTaskReq.GetDueDate(), newTaskReq.GetCategoryId(), newTaskReq.GetAuthenticatedUserId()))
		if cErr != nil {
			log.Printf("can't create task \nerror: %v", cErr)

			taskResponse := deliveryparam.NewTaskResponse("", 0, cErr)
			serializedData, mErr := json.Marshal(taskResponse)
			if mErr != nil {
				log.Printf("can't marshal data  in new-task: %v", mErr)
			}

			if _, wErr := connection.Write(serializedData); wErr != nil {
				log.Printf("can't write data to connection: %v", wErr)
			}
			return
		}
		taskResponse := deliveryparam.NewTaskResponse(newTaskRes.GetTask().GetTitle(), newTaskRes.GetTask().GetId(), nil)
		serializedData, mErr := json.Marshal(taskResponse)
		if mErr != nil {
			log.Printf("can't marshal data  in new-task: %v", mErr)
		}

		if _, wErr := connection.Write(serializedData); wErr != nil {
			log.Printf("can't write data to connection: %v", wErr)
		}

		log.Printf("task [%s] by id: [%d] is create!\n", taskResponse.GetTitle(), taskResponse.GetTaskId())
	case "list-task":

		var taskListReq = deliveryparam.NewListTaskRequest(0)
		var buffer = make([]byte, 500)
		numberOfReadBytes, rErr := connection.Read(buffer)
		if rErr != nil {
			log.Printf("can't read data from connection in list-task, error: %v", rErr)
		}
		if uErr := json.Unmarshal(buffer[:numberOfReadBytes], taskListReq); uErr != nil {
			log.Printf("can't unmarshal data in list-category response %v", uErr)
		}

		responseListTask, lErr := taskService.ListTask(taskparam.NewListRequest(taskListReq.GetAuthenticatedUserId()))
		if lErr != nil {
			if _, wErr := connection.Write([]byte(lErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)
			}
		}
		data, mErr := json.Marshal(responseListTask)
		if mErr != nil {
			log.Println("can't marshal responseCreatedTask:", mErr)
		}
		if _, wErr := connection.Write(data); wErr != nil {
			log.Println("can't write data to connection", wErr)
		}
	case "tasks-date":

		var taskListReq = deliveryparam.NewListTaskByDateRequest(0, "")
		var buffer = make([]byte, 500)
		numberOfReadBytes, rErr := connection.Read(buffer)
		if rErr != nil {
			log.Printf("can't read data from connection in list-task, error: %v", rErr)
		}
		if uErr := json.Unmarshal(buffer[:numberOfReadBytes], taskListReq); uErr != nil {
			log.Printf("can't unmarshal data in list-category response %v", uErr)
		}

		taskListRes, lErr := taskService.ListTaskByDueDate(taskparam.NewListByDateRequest(taskListReq.GetAuthenticatedUserId(), taskListReq.GetDueDate()))
		if lErr != nil {
			if _, wErr := connection.Write([]byte(lErr.Error())); wErr != nil {
				log.Println("can't write data to connection", wErr)
			}
		}
		data, mErr := json.Marshal(taskListRes)
		if mErr != nil {
			log.Println("can't marshal taskListRes:", mErr)
		}
		if _, wErr := connection.Write(data); wErr != nil {
			log.Println("can't write data to connection", wErr)
		}
	}
}

func getCommand(connection net.Conn) string {

	buffer := make([]byte, 1024)

	numberOfReadBytes, rErr := connection.Read(buffer)
	if rErr != nil {
		log.Println("error in reading the buffer connection, error:", rErr)

		//break
	}

	commandRequest := deliveryparam.NewCommand("")
	if uErr := json.Unmarshal(buffer[:numberOfReadBytes], commandRequest); uErr != nil {
		log.Println("bad request", uErr)
	}
	return commandRequest.GetCommand()
}
