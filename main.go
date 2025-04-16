package main

import (
	"flag"
	"fmt"
	"gocasts.ir/go-fundamentals/todo-cli/constant"
	"gocasts.ir/go-fundamentals/todo-cli/contract"
	"gocasts.ir/go-fundamentals/todo-cli/entity"
	"gocasts.ir/go-fundamentals/todo-cli/filestore"
	"gocasts.ir/go-fundamentals/todo-cli/pkg"
	"os"
	"strconv"
)

var authenticatedUser *entity.User

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

func registeredUser(store contract.Store[entity.User]) {

	fmt.Print("register user!\n")

	fmt.Print("enter name: ")
	var name string = pkg.ReadInput()

	fmt.Print("enter email: ")
	var email string = pkg.ReadInput()

	fmt.Print("enter password: ")
	var password []uint8 = pkg.HashPassword(pkg.ReadInput())

	user := entity.NewUser(uint(len(store.GetObjectsStore())+1), name, email, password)

	store.SetObjectsStore(append(store.GetObjectsStore(), user))

	store.Save(user)

	fmt.Printf("%s is registerd!\n", user.GetEmail())
}

func newCategory(store contract.Store[entity.Category]) uint {

	fmt.Print("enter title: ")
	var title string = pkg.ReadInput()

	fmt.Print("enter color: ")
	var color string = pkg.ReadInput()

	cId := uint(len(store.GetObjectsStore()) + 1)
	c := entity.NewCategory(cId, title, color, authenticatedUser.GetId())

	store.SetObjectsStore(append(store.GetObjectsStore(), c))

	store.Save(c)

	fmt.Printf("category [%s] is create!\n", c.GetTitle())

	return cId
}

func newTask(store contract.Store[entity.Task]) {

	fmt.Print("enter title: ")
	var title string = pkg.ReadInput()

	fmt.Print("enter due date: ")
	var dueDate string = pkg.ReadInput()

	fmt.Print("enter category: ")
	var category, _ = strconv.Atoi(pkg.ReadInput())

	t := entity.NewTask(uint(len(store.GetObjectsStore()))+1, title, dueDate, uint(category), authenticatedUser.GetId())
	store.SetObjectsStore(append(store.GetObjectsStore(), t))

	store.Save(t)

	fmt.Printf("task [%s] is create!\n", t.GetTitle())
}

func listTask(store contract.Store[entity.Task]) []*entity.Task {
	ut := make([]*entity.Task, 0)

	for _, task := range store.GetObjectsStore() {
		if task.GetUserId() == authenticatedUser.GetId() {
			ut = append(ut, task)
		}
	}

	return ut
}

func listCategory(store contract.Store[entity.Category]) []*entity.Category {
	uc := make([]*entity.Category, 0)

	for _, cat := range store.GetObjectsStore() {
		if cat.GetUserId() == authenticatedUser.GetId() {
			uc = append(uc, cat)
		}
	}

	return uc
}

func tasksByDate(store contract.Store[entity.Task]) []*entity.Task {

	fmt.Print("enter date: ")
	var date string = pkg.ReadInput()

	var tbd = make([]*entity.Task, 0)

	for _, t := range store.GetObjectsStore() {
		if t.GetDueDate() == date && t.GetUserId() == authenticatedUser.GetId() {
			tbd = append(tbd, t)
		}
	}

	return tbd
}

func login(store contract.Store[entity.User]) {

	fmt.Println("login process...")

	fmt.Print("enter email: ")
	var email string = pkg.ReadInput()

	fmt.Print("enter password: ")
	var password string = pkg.ReadInput()
	hPass := pkg.HashPassword(password)

	for _, user := range store.GetObjectsStore() {
		if user.GetEmail() == email && string(user.GetPassword()) == string(hPass) {
			authenticatedUser = user

			fmt.Printf("welcome %s\n", user.GetEmail())

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Print("invalid email or password!\n")
	}
}

func runCommand(command string, uStore contract.Store[entity.User], tStore contract.Store[entity.Task], cStore contract.Store[entity.Category]) {

	if command != "login" && command != "register-user" && command != "exit" && authenticatedUser == nil {
		login(uStore)

		return
	}
	switch command {
	case "login":
		login(uStore)
	case "register-user":
		registeredUser(uStore)
	case "new-category":
		newCategory(cStore)
	case "new-task":
		newTask(tStore)
	case "list-task":
		PrintObjects[entity.Task](listTask(tStore))
	case "list-category":
		PrintObjects[entity.Category](listCategory(cStore))
	case "tasks-date":
		fmt.Println(tasksByDate(tStore))
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("invalid command input!")
	}
}

func PrintObjects[T any](objects []*T) {
	for _, obj := range objects {
		fmt.Print(obj)
	}
}

func main() {

	var command string
	flag.StringVar(&command, "command", "no-command", "command to run")
	flag.Parse()

	// load data from storage
	var userStore = filestore.NewStore[entity.User](constant.UsersFile, constant.PermFile)
	userStore.SetObjectsStore(append(userStore.GetObjectsStore(), userStore.Load(new(entity.User))...))

	var taskStore = filestore.NewStore[entity.Task](constant.TasksFile, constant.PermFile)
	taskStore.SetObjectsStore(append(taskStore.GetObjectsStore(), taskStore.Load(new(entity.Task))...))

	var categoryStore = filestore.NewStore[entity.Category](constant.CategoriesFile, constant.PermFile)
	categoryStore.SetObjectsStore(append(categoryStore.GetObjectsStore(), categoryStore.Load(new(entity.Category))...))

	for {
		runCommand(command, userStore, taskStore, categoryStore)

		fmt.Print("please enter another command: ")
		command = pkg.ReadInput()
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
