package handler

import (
	"covid19/homepage"
	"covid19/indiacovid"
	"covid19/vaccine"
	"covid19/worldcovid"
	"covid19/worldstatescovid"
	"log"
	"net/http"
)

func HandlerMethod() {
	log.Println("Server started on: http://localhost:8080")

	http.HandleFunc("/indiaStates/", indiacovid.GetIndiaStatesdCovidDetails)
	http.HandleFunc("/world/", worldcovid.GetWorldCovidDetails)
	http.HandleFunc("/vaccinated/", vaccine.GetWorldVaccinationDetails)
	http.HandleFunc("/states/", worldstatescovid.GetWorldStatesCovidDetails)
	http.HandleFunc("/", homepage.Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
