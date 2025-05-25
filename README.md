# Todo-CLI Application 📝

Todo-CLI is a simple command-line interface (CLI) application for managing your tasks. It allows users to register, log in, and manage categories and tasks through a TCP client-server architecture. This project was developed as a practice exercise to get acquainted with the Go programming language and related concepts. Data is persisted in JSON files on the server-side.

---

## 🌟 Core Features

* **User Management**: Secure user registration and login.
* **Category Management**: Create and list personalized categories for each user.
* **Task Management**:
    * Create new tasks and assign them to categories.
    * List all tasks for a user.
    * List tasks by due date.
    * List tasks by their completion status (Done/UnDone).
* **Client-Server Communication**: Interaction via TCP/IP protocol with data exchanged in JSON format.
* **Data Persistence**: User, task, and category data are stored in JSON files (`users.json`, `tasks.json`, `categories.json`) on the server for persistence.
* **Command-Line Interface (CLI)**: Interact with the application using text-based commands in the terminal.

---

## 🛠️ Getting Started

### Prerequisites

* [Go (version 1.18 or higher recommended)](https://golang.org/dl/)

### Installation & Setup

1.  **Clone the Repository (Optional - if you have the code locally):**
    ```bash
    git clone [https://github.com/mohammadrezajavid-lab/todo-cli.git](https://github.com/mohammadrezajavid-lab/todo-cli.git)
    cd todo-cli-master
    ```

2.  **Build the Client and Server:**
    First, you need to build the executable files for the client and server. Run the following commands from the project's root directory (`todo-cli-master`):

    * **Build Client:**
        ```bash
        cd ./delivery/tcpclient
        go build -o client-todocli client.go
        cd ../../
        ```

    * **Build Server:**
        ```bash
        cd ./delivery/tcpserver
        go build -o server-todocli server.go
        cd ../../
        ```
    After these commands, `client-todocli` and `server-todocli` executables will be created in their respective directories (`delivery/tcpclient` and `delivery/tcpserver`).

3.  **Run the Server:**
    The server must be running before the client can connect.
    ```bash
    ./delivery/tcpserver/server-todocli
    ```
    By default, the server starts and listens on `127.0.0.1:1999`.
    You can change the IP address and port using the `-ip` flag:
    ```bash
    ./delivery/tcpserver/server-todocli -ip <ip_address:port>
    # Example: ./delivery/tcpserver/server-todocli -ip 0.0.0.0:2000
    ```
    The data files (`users.json`, `tasks.json`, `categories.json`) will be automatically created in the `delivery/tcpserver` directory if they don't exist.

4.  **Run the Client and Use Commands:**
    Once the server is running, you can start the client to send commands.

    * **Show Client Help:**
        ```bash
        ./delivery/tcpclient/client-todocli -h
        ```

    * **Run Client with an Initial Command:**
        ```bash
        ./delivery/tcpclient/client-todocli --command <command_name>
        ```
        After executing the initial command, the client will prompt you for subsequent commands.
        You can specify the server's IP address and port using the `--ip` flag (if the server is not running on the default address):
        ```bash
        ./delivery/tcpclient/client-todocli --command <command_name> --ip <server_ip_address:port>
        ```

### Available Client Commands

Once connected to the server, you can use the following commands (as prompted by the client):

* `register-user`: Registers a new user.
    * Prompts for name, email, and password.
* `login-user`: Logs into an existing user account.
    * Prompts for email and password.
* `new-category`: Creates a new category (requires login).
    * Prompts for category title and color.
* `list-category`: Displays the user's list of categories (requires login).
* `new-task`: Creates a new task (requires login).
    * Prompts for title, due date, and category ID.
* `list-task`: Displays all tasks for the logged-in user (requires login).
* `tasks-date`: Displays user's tasks filtered by due date (requires login).
    * Prompts for the due date.
* `list-task-status`: Displays user's tasks filtered by completion status (Done/UnDone) (requires login).
    * Prompts for status ("Done" or "UnDone").
* `edit-task`: Edits an existing task (currently not implemented).
* `exit`: Exits the client application and closes the connection.

**Important Note:** If you attempt to run a command that requires login before you have logged in, the client will first direct you to the login process.

---

## 🏗️ Project Structure

The project is organized into the following main directories:

* **`constant/`**: Contains application constants like data storage filenames and permissions.
* **`delivery/`**: Communication layer; includes:
    * `tcpclient/`: TCP client implementation.
    * `tcpserver/`: TCP server implementation.
    * `deliveryparam/`: Data Transfer Objects (DTOs) for client-server communication.
* **`entity/`**: Defines the core domain models (`User`, `Task`, `Category`).
* **`pkg/`**: Shared utility packages (input reading, hashing).
* **`repository/`**: Data access layer; includes:
    * `filestore/`: Generic logic for reading and writing entities to JSON files.
    * `memoryStore/`: Specific implementations for each entity, using `filestore` and also caching data in memory.
    * `repositorycontract/`: Interfaces related to data storage.
* **`service/`**: Business logic layer; includes:
    * `user/`, `task/`, `category/`: Services for each entity, including their specific request/response parameters.
    * `servicecontract/`: Interfaces expected by services from the repository layer.

---

## 🔌 Server-Side Operations (Conceptual Endpoints)

While this is a TCP-based application and not HTTP, the server handles commands that are conceptually similar to REST API endpoints. The client sends a command string, followed by JSON data if required.

* **Authentication**:
    * `register-user` (Payload: `deliveryparam.RegisterUserRequest`) -> Creates a new user.
    * `login-user` (Payload: `deliveryparam.UserRequest`) -> Authenticates a user and returns their ID.
* **Categories**:
    * `new-category` (Payload: `deliveryparam.CategoryRequest`) -> Creates a new category for the authenticated user.
    * `list-category` (Payload: `deliveryparam.CategoryListRequest`) -> Lists categories for the authenticated user.
* **Tasks**:
    * `new-task` (Payload: `deliveryparam.TaskRequest`) -> Creates a new task for the authenticated user.
    * `list-task` (Payload: `deliveryparam.ListTaskRequest`) -> Lists all tasks for the authenticated user.
    * `tasks-date` (Payload: `deliveryparam.ListTaskByDateRequest`) -> Lists tasks by due date for the authenticated user.
    * `list-task-status` (Payload: `deliveryparam.ListTaskByStatusRequest`) -> Lists tasks by status for the authenticated user.
* **Control**:
    * `exit` -> Signals the server that the client is disconnecting.

---
