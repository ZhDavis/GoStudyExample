package main

import (
"fmt"
"html"
"log"
"net/http"
"github.com/gorilla/mux"
"encoding/json"
)

type User struct {
    Name string `json:"name"`
    Age  int `json:"age"`
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/",HandleIndex)
    router.HandleFunc("/Domain", HandleDomainIndex)
    router.HandleFunc("/Domain/{id}",HandleDomainShow)
    fmt.Println("Main tasks:")

   log.Fatal(http.ListenAndServe(":80",router))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Router test: Hello, %q \n", html.EscapeString(r.URL.Path))
    fmt.Fprintf(w,"Router test: Hello, %q", html.EscapeString(r.URL.Path))
}

func HandleDomainIndex(w http.ResponseWriter, r *http.Request) {
    userInfo := User {
        Name:"zhanghui",
        Age:18,
    }
    fmt.Printf("Router test:  %q \n", userInfo)
    json.NewEncoder(w).Encode(userInfo)
}

func HandleDomainShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Printf("Router test:  %q \n", id)
    fmt.Fprintf(w,"Domain show:", id)
}
