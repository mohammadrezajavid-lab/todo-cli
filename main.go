package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/constant"
	"gocasts.ir/go-fundamentals/todo-cli/contract"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/filestore"
	"os"
	"strconv"
)

var (
	authenticatedUser *entity.User
	userTasks         []*entity.Task
	userStorage       []*entity.User
	categories        []*entity.Category
	tasks             []*entity.Task

	scanner *bufio.Scanner
)

func init() {

	scanner = bufio.NewScanner(os.Stdin)

	func() {
		fUsers, _ := os.OpenFile(constant.UsersFile, os.O_CREATE, constant.PermFile)
		fTasks, _ := os.OpenFile(constant.TasksFile, os.O_CREATE, constant.PermFile)
		fCategories, _ := os.OpenFile(constant.CategoriesFile, os.O_CREATE, constant.PermFile)

		_ = fUsers.Close()
		_ = fTasks.Close()
		_ = fCategories.Close()
	}()
}

func registeredUser(store contract.Store[entity.User]) {

	fmt.Print("register user!\n")

	fmt.Print("enter name: ")
	var name string = readInput()

	fmt.Print("enter email: ")
	var email string = readInput()

	fmt.Print("enter password: ")
	var password []uint8 = hashPassword(readInput())

	user := &entity.User{
		ID:       uint(len(userStorage) + 1),
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	store.Save(user)

	fmt.Printf("%s is registerd!\n", user.Email)
}

func newCategory(store contract.Store[entity.Category]) uint {

	fmt.Print("enter title: ")
	var title string = readInput()

	fmt.Print("enter color: ")
	var color string = readInput()

	cId := uint(len(categories) + 1)
	c := &entity.Category{
		ID:     cId,
		Title:  title,
		Color:  color,
		UserId: authenticatedUser.ID,
	}

	categories = append(categories, c)

	store.Save(c)

	fmt.Printf("category [%s] is create!\n", c.Title)

	return cId
}

func newTask(store contract.Store[entity.Task]) {

	fmt.Print("enter title: ")
	var title string = readInput()

	fmt.Print("enter due date: ")
	var dueDate string = readInput()

	fmt.Print("enter category: ")
	var category, _ = strconv.Atoi(readInput())

	t := &entity.Task{
		Title:    title,
		DueDate:  dueDate,
		Category: uint(category),
		IsDone:   false,
		UserId:   authenticatedUser.ID,
	}
	tasks = append(tasks, t)

	store.Save(t)

	fmt.Printf("task [%s] is create!\n", t.Title)
}

func listTask() []*entity.Task {

	for _, task := range tasks {
		if task.UserId == authenticatedUser.ID {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks
}

func tasksByDate() []*entity.Task {

	fmt.Print("enter date: ")
	var date string = readInput()

	if userTasks == nil {
		listTask()
	}

	var tbd []*entity.Task
	for _, task := range userTasks {
		if task.DueDate == date {
			tbd = append(tbd, task)
		}
	}

	return tbd
}

func login() {

	fmt.Println("login process...")

	fmt.Print("enter email: ")
	var email string = readInput()

	fmt.Print("enter password: ")
	var password string = readInput()
	hPass := hashPassword(password)

	for _, user := range userStorage {
		if user.Email == email && string(user.Password) == string(hPass) {
			authenticatedUser = user

			fmt.Printf("welcome %s\n", user.Email)

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Print("invalid email or password!\n")
	}
}

func hashPassword(password string) []uint8 {

	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return bs
}

func runCommand(command string, uStore contract.Store[entity.User], tStore contract.Store[entity.Task], cStore contract.Store[entity.Category]) {

	if command != "login" && command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		return
	}

	switch command {
	case "login":
		login()
	case "register-user":
		registeredUser(uStore)
	case "new-category":
		newCategory(cStore)
	case "new-task":
		newTask(tStore)
	case "list-task":
		printTasks(listTask())
	case "tasks-date":
		fmt.Println(tasksByDate())
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("invalid command input!")
	}
}

func printTasks(tasks []*entity.Task) {
	for _, t := range tasks {
		fmt.Print(t.String())
	}
}

func readInput() string {
	scanner.Scan()

	return scanner.Text()
}

func main() {

	var command string
	flag.StringVar(&command, "command", "no-command", "command to run")
	flag.Parse()

	// load data from storage
	var userStore = &filestore.FileStore[entity.User]{FilePath: constant.UsersFile, PermFile: constant.PermFile}
	userStorage = append(userStorage, userStore.Load(new(entity.User))...)

	var taskStore = &filestore.FileStore[entity.Task]{FilePath: constant.TasksFile, PermFile: constant.PermFile}
	tasks = append(tasks, taskStore.Load(new(entity.Task))...)

	var categoryStore = &filestore.FileStore[entity.Category]{FilePath: constant.CategoriesFile, PermFile: constant.PermFile}
	categories = append(categories, categoryStore.Load(new(entity.Category))...)

	for {
		runCommand(command, userStore, taskStore, categoryStore)

		fmt.Print("please enter another command: ")
		command = readInput()
	}
}

/*
	// create object of file
	file, _ := os.OpenFile("users.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)

	// defer file close
	defer func(f *os.File) {
		cErr := f.Close()
		if cErr != nil {
			panic(cErr)
		}
	}(file)

	// write in file
	for i := 0; i < 10; i++ {

		s := fmt.Sprintf("paragraf%d\n", i)

		_, wErr := file.Write([]byte(s))
		if wErr != nil {

			panic(wErr)
		}
	}

	//readFile("users.txt")

	// seek in file for read
	func(f *os.File) {
		_, sErr := f.Seek(0, io.SeekStart)
		if sErr != nil {
			panic(sErr)
		}
	}(file)

	// read as file
	resultOfReadFile := func(f *os.File) []byte {
		bs, rErr := io.ReadAll(file)
		if rErr != nil {
			panic(rErr)
		}
		return bs
	}(file)

	fmt.Println(string(resultOfReadFile))
*/
