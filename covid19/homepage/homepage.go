package homepage

import (
	"covid19/constant"
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	webPageHome, err := template.ParseFiles(constant.HomePage)
	if err != nil {
		log.Fatal(err)
	}
	webPageHome.Execute(w, "home")

}
