package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
creting a REST APi that allows to perform CRUD services on our website
talk about Crud apis , so we are referring to an Api that an
handle all of these tasks:-
Creating,Reading,Updating,Deleting
*/

//create an Article struct that have features
type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`

}
//lets declare a global Articles array
//to simulate a database

var Articles []Article

func returnAllArticles(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func handleRequestss(){
	http.HandleFunc("/",returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequestss()
}