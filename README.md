# Creating simplest API in GO

## Download and Install Go
https://golang.org/dl/

## Directory Structure
Create a folder 'go' at your home directory as $HOME/go is the default $GOPATH. Next, create the following directory structure:

```cgo
bin/
src/
    github.com/shoeb240/	
```

Clone this project from github with the following git command at src/github.com/shoeb240/
```cgo
git clone http://github.com/shoeb240/first-go-api
```

Now, you will have the following directory structure:
```cgo
bin/
    hello                          # command executable
src/
    github.com/shoeb240/
	    first-go-api/
	        main.go               # command source
	        route.go               
	        handlers.go               
```

## route.go
Creating api route '/users' and opening port 8081 to listen to request using gorilla mux

```go
package main

import (
    "github.com/gorilla/mux"
    "net/http"
)

func RequestHandle() {
    myRouter := mux.NewRouter()
    myRouter.HandleFunc("/users", getUsers).Methods("GET")

    http.ListenAndServe(":8081", myRouter)
}
```

## handlers.go
Creating getUsers handler method to respond to the route

```
package main

import (
    "fmt"
    "net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Showing users list")
}
```

## main.go
Calling RequestHandle() from main

```go
package main

import (
    "fmt"
)

func main(){
    fmt.Println("Waiting for request to serve...")
    RequestHandle()
}
```

Run the following command
```
go run github.com/shoeb240/first-go-api
```

It will output
```
Waiting for request to serve...
```

Open your browser to open url 'localhost:8081/users'. It will show..
```
Showing users list
```
