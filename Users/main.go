package main

import (
	"fmt"
	"net/http"
)

func users(reswt http.ResponseWriter, req *http.Request) {
	fmt.Fprint(reswt, "hello world")
}

func main() {
	http.HandleFunc("/users", users)
	http.ListenAndServe(":7070", nil)

}
