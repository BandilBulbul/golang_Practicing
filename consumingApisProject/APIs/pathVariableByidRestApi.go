package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

/*just view one articles only
the gorilla mux router we can add variables to our paths and then pick and choose what articles we want to
return based on these variables
*/

type Product struct{
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
var Products []Product

func returnSingleArticle(w http.ResponseWriter,r *http.Request){
	vars :=mux.Vars(r)
	key:=vars["id"]
	//fmt.Fprintf(w,"key:"+key)
	for _,product:=range Products{
		if product.Id==key{
			json.NewEncoder(w).Encode(product)
		}
	}

}
func createNewProduct(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody,&product)
	Products=append(Products,product)
	json.NewEncoder(w).Encode(product)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
func deleteProductBYID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, product := range Products {
		if product.Id == id {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}

}


func handleReq(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/product/{id}", returnSingleArticle)
	myRouter.HandleFunc("/product", createNewProduct).Methods("POST")
	myRouter.HandleFunc("/product/{id}",deleteProductBYID).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
func main(){
	Products =[]Product{
		Product{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Product{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleReq()

}

/*reference:https://tutorialedge.net/golang/creating-restful-api-with-golang/*/
