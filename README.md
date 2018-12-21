# Creating simplest API in GO

## Download and Install Go
https://golang.org/dl/

## Directory Structure
Create a folder 'go' at your home directory as $HOME/go is the default $GOPATH. Create 'src' and 'bin' directory within 'go' directory. Packages reside inside 'src' directory. 

Now, run the follwoing command to get this project with all dependency packages:
```cgo
go get github.com/shoeb240/first-go-api
```

OR

Create the following directory structure within $HOME/go directory:

```cgo
bin/
src/
    github.com/shoeb240/	
```

cd to src/github.com/shoeb240/ and clone this project from github with the following git command:
```cgo
git clone http://github.com/shoeb240/first-go-api
```

To download Gorilla Mux pacjake run the following command:
```cgo
go get github.com/gorilla/mux
```


Now, you will have the following directory structure:
```cgo
bin/
src/
    github.com/gorilla/
	    mux/
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

```go
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
I am responding to your API call
```
