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

var authenticatedUserId uint

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
		log.Fatalf("can't marshal data command login-user: %v", mErr)
	}

	if _, wErr := connection.Write(serializedUser); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawResponse = make([]byte, 1024)
	numberOfReadBytes, rErr := connection.Read(rawResponse)
	if rErr != nil {
		log.Fatalf("can't read data from connection in login-user, error: %v", rErr)
	}
	var userResponse *deliveryparam.UserResponse = deliveryparam.NewUserResponse(0)
	if uErr := json.Unmarshal(rawResponse[:numberOfReadBytes], userResponse); uErr != nil {
		log.Fatalf("can't unmarshal data in login-user response %v", uErr)
	}

	authenticatedUserId = userResponse.GetAuthenticateUserId()

	if authenticatedUserId == 0 {
		fmt.Print("invalid email or password!\n")
	} else {
		fmt.Println("welcome to todo-cli, successful login")
	}
}

func registerUser(command string, connection net.Conn) {
	fmt.Print("register user!\n")

	fmt.Print("enter name: ")
	var name string = pkg.ReadInput()

	fmt.Print("enter email: ")
	var email string = pkg.ReadInput()

	fmt.Print("enter password: ")
	var password string = pkg.ReadInput()

	sendCommand(command, connection)

	requestRegisterUser := deliveryparam.NewRegisterUserRequest(name, email, password)
	serializedUser, mErr := json.Marshal(requestRegisterUser)
	if mErr != nil {
		log.Fatalf("can't marshal data command register-user: %v", mErr)
	}
	if _, wErr := connection.Write(serializedUser); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawResponse = make([]byte, 1024)
	numberOfReadBytes, rErr := connection.Read(rawResponse)
	if rErr != nil {
		log.Fatalf("can't read data from connection in register-user, error: %v", rErr)
	}
	var userRegisterResponse *deliveryparam.RegisterUserResponse = deliveryparam.NewRegisterUserResponse("", nil)
	if uErr := json.Unmarshal(rawResponse[:numberOfReadBytes], userRegisterResponse); uErr != nil {
		log.Fatalf("can't unmarshal data in register-user response %v", uErr)
	}
	if userRegisterResponse.GetError() != nil {
		log.Fatalf("can't register user, email: %s\nerror: %v", userRegisterResponse.GetEmail(), userRegisterResponse.GetError())
	}
	fmt.Printf("%s is registerd!\n", userRegisterResponse.GetEmail())
}

func newCategory(command string, connection net.Conn) {

	fmt.Println("Create New Category!")

	fmt.Print("enter title: ")
	var title string = pkg.ReadInput()

	fmt.Print("enter color: ")
	var color string = pkg.ReadInput()

	sendCommand(command, connection)

	requestCreateCategory := deliveryparam.NewCategoryRequest(title, color, authenticatedUserId)
	serializedCategory, mErr := json.Marshal(requestCreateCategory)
	if mErr != nil {
		log.Fatalf("can't marshal data new-category: %v", mErr)
	}
	if _, wErr := connection.Write(serializedCategory); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawResponse = make([]byte, 1024)
	numberOfReadBytes, rErr := connection.Read(rawResponse)
	if rErr != nil {
		log.Fatalf("can't read data from connection in new-category, error: %v", rErr)
	}
	var responseCreateCategory = deliveryparam.NewCategoryResponse("", 0, nil)
	if uErr := json.Unmarshal(rawResponse[:numberOfReadBytes], responseCreateCategory); uErr != nil {
		log.Fatalf("can't unmarshal data in new-category response %v", uErr)
	}
	if responseCreateCategory.GetError() != nil {
		log.Fatalf("can't create this category, \nerror: %v", responseCreateCategory.GetError())
	}

	fmt.Printf("category [%s] by id: [%d] is create!\n", responseCreateCategory.GetTitle(), responseCreateCategory.GetCategoryId())
}

func listCategory(command string, connection net.Conn) {

	sendCommand(command, connection)

	var catListReq = deliveryparam.NewCategoryListRequest(authenticatedUserId)
	serializedCatListReq, mErr := json.Marshal(catListReq)
	if mErr != nil {
		log.Fatalf("can't marshal data list-category: %v", mErr)
	}
	if _, wErr := connection.Write(serializedCatListReq); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawCategories = make([]byte, 2048)
	numberOfReadBytes, rErr := connection.Read(rawCategories)
	if rErr != nil {
		log.Fatalf("can't read data from connection in list-category, error: %v", rErr)
	}
	var catListRes = deliveryparam.NewCategoryListResponse()
	if uErr := json.Unmarshal(rawCategories[:numberOfReadBytes], catListRes); uErr != nil {
		log.Fatalf("can't unmarshal data in list-category response %v", uErr)
	}

	fmt.Printf("your categoreis: \n%s\n", catListRes.String())
}

func newTask(command string, connection net.Conn) {

	fmt.Println("Create New Task!")

	fmt.Print("enter title: ")
	var title string = pkg.ReadInput()

	fmt.Print("enter due date: ")
	var dueDate string = pkg.ReadInput()

	fmt.Print("enter category: ")
	var categoryId, _ = strconv.Atoi(pkg.ReadInput())

	sendCommand(command, connection)

	var reqNewTask = deliveryparam.NewTaskRequest(title, dueDate, uint(categoryId), authenticatedUserId)
	marshalReqNewTask, mErr := json.Marshal(reqNewTask)
	if mErr != nil {
		log.Fatalf("can't marshal data new-task: %v", mErr)
	}
	if _, wErr := connection.Write(marshalReqNewTask); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawResponse = make([]byte, 1024)
	numberOfReadBytes, rErr := connection.Read(rawResponse)
	if rErr != nil {
		log.Fatalf("can't read data from connection in new-task, error: %v", rErr)
	}
	var resNewTask = deliveryparam.NewTaskResponse("", 0, nil)
	if uErr := json.Unmarshal(rawResponse[:numberOfReadBytes], resNewTask); uErr != nil {
		log.Fatalf("can't unmarshal data in new-category response %v", uErr)
	}
	if resNewTask.GetError() != nil {
		log.Fatalf("can't create this category, \nerror: %v", resNewTask.GetError())
	}

	fmt.Printf("task [%s] by id: [%d] is create!\n", resNewTask.GetTitle(), resNewTask.GetTaskId())
}

func listTask(command string, connection net.Conn) {

	sendCommand(command, connection)

	var taskListReq = deliveryparam.NewListTaskRequest(authenticatedUserId)
	marshalTaskListReq, mErr := json.Marshal(taskListReq)
	if mErr != nil {
		log.Fatalf("can't marshal data list-task: %v", mErr)
	}
	if _, wErr := connection.Write(marshalTaskListReq); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawTasks = make([]byte, 2048)
	numberOfReadBytes, rErr := connection.Read(rawTasks)
	if rErr != nil {
		log.Fatalf("can't read data from connection in list-task, error: %v", rErr)
	}
	var taskListRes = deliveryparam.NewListTaskResponse()
	if uErr := json.Unmarshal(rawTasks[:numberOfReadBytes], taskListRes); uErr != nil {
		log.Fatalf("can't unmarshal data in list-task response %v", uErr)
	}

	fmt.Printf("your tasks: \n%s\n", taskListRes.String())
}

func tasksByDate(command string, connection net.Conn) {

	fmt.Println("list tasks by date...")

	fmt.Print("enter due date: ")
	var dueDate string = pkg.ReadInput()

	sendCommand(command, connection)

	var taskListDateReq = deliveryparam.NewListTaskByDateRequest(authenticatedUserId, dueDate)
	marshalTaskListDateReq, mErr := json.Marshal(taskListDateReq)
	if mErr != nil {
		log.Fatalf("can't marshal data list-task: %v", mErr)
	}
	if _, wErr := connection.Write(marshalTaskListDateReq); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawTasks = make([]byte, 2048)
	numberOfReadBytes, rErr := connection.Read(rawTasks)
	if rErr != nil {
		log.Fatalf("can't read data from connection in task-date, error: %v", rErr)
	}
	var taskListDateRes = deliveryparam.NewListTaskByDateResponse()
	if uErr := json.Unmarshal(rawTasks[:numberOfReadBytes], taskListDateRes); uErr != nil {
		log.Fatalf("can't unmarshal data in list-task response %v", uErr)
	}

	fmt.Printf("your tasks: \n%s\n", taskListDateRes.String())
}

func listTaskStatus(command string, connection net.Conn) {

	fmt.Println("list tasks is Done/UnDone...")

	fmt.Println("enter Done/UnDone: ")
	var taskStatus string = pkg.ReadInput()

	if taskStatus != "Done" || taskStatus != "UnDone" {
		log.Fatalf("Input value is invalid.")
	}

	sendCommand(command, connection)

	var taskList = deliveryparam.NewListTaskByStatusRequest(authenticatedUserId, taskStatus)
	marshalTaskList, mErr := json.Marshal(taskList)
	if mErr != nil {
		log.Fatalf("can't marshal data list-task-status: %v", mErr)
	}
	if _, wErr := connection.Write(marshalTaskList); wErr != nil {
		log.Fatalf("can't write data to connection: %v", wErr)
	}

	var rawTasks = make([]byte, 2048)
	numberOfReadBytes, rErr := connection.Read(rawTasks)
	if rErr != nil {
		log.Fatalf("can't read data from connection in list-task-status, error: %v", rErr)
	}
	var tasksStatusRes = deliveryparam.NewListTaskByStatusResponse()
	if uErr := json.Unmarshal(rawTasks[:numberOfReadBytes], tasksStatusRes); uErr != nil {
		log.Fatalf("can't unmarshal data in list-task response %v", uErr)
	}

	fmt.Printf("your tasks: \n%s\n", tasksStatusRes.String())
}

func runCommand(command string, connection net.Conn) {

	if command != "login-user" && command != "register-user" && command != "exit" && authenticatedUserId == 0 {
		login(command, connection)
		return
	}
	switch command {
	case "login-user":
		login(command, connection)
	case "register-user":
		registerUser(command, connection)
	case "new-category":
		newCategory(command, connection)
	case "list-category":
		listCategory(command, connection)
	case "new-task":
		newTask(command, connection)
	case "list-task":
		listTask(command, connection)
	case "tasks-date":
		tasksByDate(command, connection)
	case "list-task-status":
		listTaskStatus(command, connection)
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

	for {
		runCommand(command, connection)

		fmt.Print("please enter another command: ")
		command = pkg.ReadInput()
	}
}
