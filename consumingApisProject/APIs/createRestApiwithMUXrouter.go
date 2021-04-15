package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"


)
//create an Article struct that have features
type Articlee struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`

}
//lets declare a global Articles array
//to simulate a database

var Articlees []Articlee

func ReturnAllArticles(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"returnAllArticles")
	json.NewEncoder(w).Encode(Articlees)
}
// Existing code from above
func handleRequestts() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/all", ReturnAllArticles)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
	//router := mux.NewRouter()

/*
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
*/
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articlees = []Articlee{
		Articlee{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Articlee{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequestts()
}
