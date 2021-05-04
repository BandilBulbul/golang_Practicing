package route

import (
	"covid19/covidcaseapp"
	"log"
	"net/http"
)

func HandlerMethod() {
	log.Println("Server started on: http://localhost:8080")

	http.HandleFunc("/indiaStates/", covidcaseapp.GetIndiaStatesdCovidDetails)
	http.HandleFunc("/world/", covidcaseapp.GetWorldCovidDetails)
	http.HandleFunc("/vaccinated/", covidcaseapp.GetWorldVaccinationDetails)
	http.HandleFunc("/states/", covidcaseapp.GetWorldStatesCovidDetails)
	http.HandleFunc("/", covidcaseapp.Home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
