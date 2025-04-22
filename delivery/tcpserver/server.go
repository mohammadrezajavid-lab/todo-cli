package main

import (
	"encoding/json"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/delivery/deliveryparam"
	"gocasts.ir/go-fundamentals/todo-cli/pkg"
	"gocasts.ir/go-fundamentals/todo-cli/repository/memoryStore"
	"gocasts.ir/go-fundamentals/todo-cli/service/category"
	"gocasts.ir/go-fundamentals/todo-cli/service/category/categoryparam"
	"gocasts.ir/go-fundamentals/todo-cli/service/user"
	"gocasts.ir/go-fundamentals/todo-cli/service/user/userparam"
	"log"
	"net"
)

func main() {

	//taskMemoryRepo := memoryStore.NewTaskMemory()
	//taskService := task.NewService(taskMemoryRepo, categoryMemoryRepo)
	var categoryMemoryRepo *memoryStore.CategoryMemory = memoryStore.NewCategoryMemory()
	var userMemoryRepo *memoryStore.UserMemory = memoryStore.NewUserMemory()
	var categoryService *category.Service = category.NewService(categoryMemoryRepo)
	var userService *user.Service = user.NewService(userMemoryRepo)

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

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("error in accepting connection, error:" + err.Error())
		}

		go handleConnection(conn, userService, categoryService)
	}
}

func handleConnection(connection net.Conn, userService *user.Service, categoryService *category.Service) {

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
		runCommand(connection, command, userService, categoryService)
	}
}

func runCommand(connection net.Conn, command string, userService *user.Service, categoryService *category.Service) {

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

		var catListReq *deliveryparam.CategoryListRequest = deliveryparam.NewCategoryListRequest(0)
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

		/*case "list-task":

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
		case "create-category":

			var newCategory *deliveryparam.Category = categoryRequest.GetCategory()
			responseCreateCategory, cErr := categoryService.CreateCategory(categoryparam.NewRequest(newCategory.GetColor(), newCategory.GetTitle(), 0))
			if cErr != nil {
				if _, wErr := connection.Write([]byte(cErr.Error())); wErr != nil {
					log.Println("can't write data to connection", wErr)

					return
					//continue
				}
				return
			}
			data, mErr := json.Marshal(responseCreateCategory)
			if mErr != nil {
				log.Println("can't marshal responseCreatedTask:", mErr)

				//continue
			}

			if _, wErr := connection.Write(data); wErr != nil {
				log.Println("can't write data to connection", wErr)

				//continue
			}
		*/
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
