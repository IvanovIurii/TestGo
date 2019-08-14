package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/IvanovYura/restApi/config"
	"github.com/IvanovYura/restApi/dao"
	"github.com/IvanovYura/restApi/model"
)

var d = dao.PostDao{}
var c = config.Config{}

func init() {
	c.Read()

	d.Server = c.Server
	d.Database = c.Database
	d.Connect()
	d.DropDb()
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	token, err := GetJWT()
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, map[string]string{"result": "not authorized"})
		return
	}
	writeResponse(w, http.StatusOK, map[string]string{"token": token})
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(findAllPosts())
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range findAllPosts() {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Post{})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	json.NewDecoder(r.Body).Decode(&post)

	if err := d.Save(post); err != nil {
		fmt.Println("Problem occured while saving post: %v "+err.Error(), post)
		return
	}
	writeResponse(w, http.StatusCreated, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range findAllPosts() {
		if item.Id == params["id"] {
			if err := d.Delete(item); err != nil {
				fmt.Println("Problem occured while deleting post: " + err.Error())
				return
			}
			writeResponse(w, http.StatusNotFound, map[string]string{"result": "success"})
			return
		}
		writeResponse(w, http.StatusNotFound, map[string]string{"result": "post not found"})
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	json.NewDecoder(r.Body).Decode(&post)

	if err := d.Update(post); err != nil {
		fmt.Println("Problem occured while updating post: " + err.Error())
		return
	}
	writeResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func writeResponse(w http.ResponseWriter, statusCode int, responseMessage interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMessage)
	w.WriteHeader(statusCode)
	w.Write(response)
}

func findAllPosts() []model.Post {
	var posts []model.Post

	posts, err := d.FindAll()
	if err != nil {
		fmt.Println("Problem occured while getting posts: " + err.Error())
	}
	return posts
}
