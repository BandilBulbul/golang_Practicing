package handler

import (
	"covid19/worldcovid"
	"log"
	"net/http"
)

func HandlerMethod() {
	log.Println("Server started on: http://localhost:8080")

	http.HandleFunc("/indiaStates/", worldcovid.GetIndiaStatesdCovidDetails)
	http.HandleFunc("/world/", worldcovid.GetWorldCovidDetails)
	http.HandleFunc("/vaccinated/", worldcovid.GetWorldVaccinationDetails)
	http.HandleFunc("/states/", worldcovid.GetWorldStatesCovidDetails)
	http.HandleFunc("/", worldcovid.Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
