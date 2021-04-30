package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func printAtHtmlPage(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "Ravi", Age: 30}

	webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\Assignments\\src\\myFirstPage.html")
	if err != nil {
		log.Print("some path issue")
	}
	webPage.Execute(w, person)

}

func main() {
	http.HandleFunc("/", printAtHtmlPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
