package main

import (
	"GoAPI/handler"
	"net/http"
)

func main() {

	http.HandleFunc("/users", handler.UsersRouters)
	http.HandleFunc("/users/", handler.UsersRouters)
	http.HandleFunc("/", handler.RootHandler)

	err := http.ListenAndServe("localhost:1111", nil)

	if err != nil {
		panic(err)
	}
}
