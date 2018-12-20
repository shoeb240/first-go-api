# Creating simplest API in GO

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
