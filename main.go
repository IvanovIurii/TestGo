package main

import (
	"net/http"

	"github.com/IvanovYura/restApi/rest"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/login", rest.GetToken).Methods("GET")

	router.Handle("/posts", rest.IsAuthorized(rest.GetPosts)).Methods("GET")
	router.Handle("/posts", rest.IsAuthorized(rest.CreatePost)).Methods("POST")
	router.Handle("/posts/{id}", rest.IsAuthorized(rest.GetPost)).Methods("GET")
	router.Handle("/posts/{id}", rest.IsAuthorized(rest.DeletePost)).Methods("DELETE")
	router.Handle("/posts", rest.IsAuthorized(rest.UpdatePost)).Methods("UPDATE")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
