package main

import (
	"fmt"
	"net/http"
	"log"
)

//to create a very simple server which can handle HTTP requests

/* A homePage function that will handle all requests to
our root URL, a handleRequests function */

func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcom to the HomePage!")
	fmt.Println("Endpoint Hit:homePage")
}
func handleRequests(){
	http.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe("8080",nil))
}
func main(){
	handleRequests()
}