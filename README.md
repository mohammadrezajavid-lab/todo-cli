# Todo-cli App

A practice project implementing a simple todo app while attending the [gocasts](https://github.com/gocasts-bootcamp) bootcamp course to get acquainted with golang

## Usage

```bash
# build client project
cd ./delivery/client
go build -o client-todocli client.go

# build server project
cd ./delivery/tcpserver
go build -o server-todocli server.go

# show help client project
./client-todocli -h

# show help server project
./server-todocli -h

# run client
./client-todocli --command [login | register-user | new-task| new-category | list-task | list-category | tasks-date | exit] --ip [ipAddr:port]

# run server (The server is available by default at the following address: 127.0.0.1:1999)
./server-todocli

```
## Entities

Category

    Properties:
        Title
        Color
    Behaviors:
        Create a new Category
        List User Categories with progress status
        Edit a category

Task

    Properties:
        Title
        DueDate
        Category
        IsDone
    Behaviors:
        Create a new Task
        List User Today Tasks
        List User Tasks By Date
        Change Task status (done/undone)
        Edit a task

User

    Properties:
        ID
        Email
        Password
    Behaviors:
        Register a user
        Log in user

## Use Cases

    User should be registered [*]
    User should be able to log in to the application [*]
    User can create a new category [*]
    User can add a new task [*]
    User can see the list of categories with progress status
    User can see the List of its tasks [*]
    User can see the Today’s Tasks
    User can see the Tasks by date [*]
    User can Done/Undone a task
    User can Edit a task
    User can Edit a category
