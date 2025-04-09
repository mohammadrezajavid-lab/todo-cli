package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Password []uint8
}

type Task struct {
	Title    string
	DueDate  string
	Category uint
	IsDone   bool
	UserId   uint
}

type Category struct {
	ID     uint
	Title  string
	Color  string
	UserId uint
}

var (
	authenticatedUser *User
	userTasks         []*Task
	userStorage       []*User
	categories        []*Category
	tasks             []*Task

	scanner *bufio.Scanner
)

const (
	usersFile      = "users.txt"
	tasksFile      = "tasks.txt"
	categoriesFile = "categories.txt"
	permFile       = 0777
)

func init() {

	scanner = bufio.NewScanner(os.Stdin)

	func() {

		fUsers, _ := os.OpenFile(usersFile, os.O_CREATE, permFile)
		fTasks, _ := os.OpenFile(tasksFile, os.O_CREATE, permFile)
		fCategories, _ := os.OpenFile(categoriesFile, os.O_CREATE, permFile)

		_ = fUsers.Close()
		_ = fTasks.Close()
		_ = fCategories.Close()
	}()

	// load users
	func() {

		usersDataByte := readFile(usersFile)
		usersStr := strings.Split(string(usersDataByte), "\n")
		var u *User = new(User)

		for _, us := range usersStr {
			if us == "" {

				continue
			}
			if err := json.Unmarshal([]byte(us), u); err != nil {

				panic(err)
			}

			userStorage = append(userStorage, u)
			u = new(User)
		}
	}()

	// load tasks
	func() {

		var t *Task = &Task{}

		tasksDataByte := readFile(tasksFile)
		tasksStr := strings.Split(string(tasksDataByte), "\n")

		for _, ts := range tasksStr {

			if ts == "" {

				continue
			}
			if err := json.Unmarshal([]byte(ts), t); err != nil {

				panic(err)
			}

			tasks = append(tasks, t)
			t = &Task{}
		}
	}()

	// load categories
	func() {

		var c *Category = &Category{}

		categoriesDataByte := readFile(categoriesFile)
		categoriesStr := strings.Split(string(categoriesDataByte), "\n")

		for _, cs := range categoriesStr {

			if cs == "" {

				continue
			}
			if err := json.Unmarshal([]byte(cs), c); err != nil {

				panic(err)
			}

			categories = append(categories, c)
			c = &Category{}
		}
	}()
}

func newCategory() uint {

	fmt.Print("enter title: ")
	var title string = readInput()

	fmt.Print("enter color: ")
	var color string = readInput()

	cId := uint(len(categories) + 1)
	c := &Category{
		ID:     cId,
		Title:  title,
		Color:  color,
		UserId: authenticatedUser.ID,
	}

	categories = append(categories, c)

	writeToFile(*serializedData(*c), categoriesFile)

	fmt.Printf("category [%s] is create!\n", c.Title)

	return cId
}

func newTask() {

	fmt.Print("enter title: ")
	var title string = readInput()

	fmt.Print("enter due date: ")
	var dueDate string = readInput()

	fmt.Print("enter category: ")
	var category, _ = strconv.Atoi(readInput())

	t := &Task{
		Title:    title,
		DueDate:  dueDate,
		Category: uint(category),
		IsDone:   false,
		UserId:   authenticatedUser.ID,
	}
	tasks = append(tasks, t)

	writeToFile(*serializedData(*t), tasksFile)

	fmt.Printf("task [%s] is create!\n", t.Title)
}

func writeToFile(object []byte, fileName string) {

	// create object of file
	file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, permFile)

	// defer file close
	defer func(f *os.File) {
		cErr := f.Close()
		if cErr != nil {
			panic(cErr)
		}
	}(file)

	object = append(object, '\n')

	func(f *os.File) {
		_, wErr := file.Write(object)
		if wErr != nil {
			panic(wErr)
		}
	}(file)
}

func listTask() []*Task {

	for _, task := range tasks {
		if task.UserId == authenticatedUser.ID {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks
}

func tasksByDate() []*Task {

	fmt.Print("enter date: ")
	var date string = readInput()

	if userTasks == nil {
		listTask()
	}

	var tbd []*Task
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

func registeredUser() {

	fmt.Print("register user!\n")

	fmt.Print("enter name: ")
	var name string = readInput()

	fmt.Print("enter email: ")
	var email string = readInput()

	fmt.Print("enter password: ")
	var password []uint8 = hashPassword(readInput())

	user := &User{
		ID:       uint(len(userStorage) + 1),
		Name:     name,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	writeToFile(*serializedData(*user), usersFile)

	fmt.Printf("%s is registerd!\n", user.Email)
}

func serializedData(vStruct any) *[]byte {

	var data, jErr = json.Marshal(vStruct)
	if jErr != nil {
		fmt.Printf("can't marshal user struct to json %v\n", jErr)
		return nil
	}

	return &data
}

func hashPassword(password string) []uint8 {

	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	return bs
}

func runCommand(command string) {

	if command != "login" && command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		return
	}

	switch command {
	case "login":
		login()
	case "register-user":
		registeredUser()
	case "new-category":
		newCategory()
	case "new-task":
		newTask()
	case "list-task":
		printTask(listTask())
	case "tasks-date":
		fmt.Println(tasksByDate())
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("invalid command input!")
	}
}

func printTask(tasks []*Task) {
	for _, t := range tasks {
		fmt.Printf(
			"title: %s, userId: %d, dueDate: %s, isDone: %v, cat: %d\n",
			t.Title, t.UserId, t.DueDate, t.IsDone, t.Category,
		)
	}
}

func readInput() string {
	scanner.Scan()

	return scanner.Text()
}

func readFile(fileName string) []byte {
	file, _ := os.OpenFile(fileName, os.O_RDONLY, permFile)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	bs, rErr := io.ReadAll(file)
	if rErr != nil {
		panic(rErr)
	}

	return bs
}

func main() {

	var command string
	flag.StringVar(&command, "command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(command)

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
