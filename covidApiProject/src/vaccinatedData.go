package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type VaccinatedData struct {
	All Values `json:"All"`
}
type Values struct {
	Administered                int       `json:"adminstered"`
	People_Vaccinated           int       `json:"people_vaccinated"`
	People_Partially_Vaccinated int       `json:"people_partially_vaccinated"`
	Country                     string    `json:"country"`
	Population                  int       `json:"population"`
	Sq_Km_Area                  int       `json:"sq_km_area"`
	Life_Expectancy             string    `json:"life_expectancy"`
	Elevation_In_Meters         int       `json:"elevation_in_meters"`
	Continent                   string    `json:"continent"`
	Abbreviation                string    `json:"abbreviation"`
	Location                    string    `json:"location"`
	ISO                         int       `json:"iso"`
	Capital_City                string    `json:"capital_city"`
	Lat                         string    `json:"lat"`
	Long                        string    `json:"long"`
	Updated                     time.Time `json:"updated"`
}

func get() {
	resp, err := http.Get("https://covid-api.mmediagroup.fr/v1/vaccines?country=India")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var vaccinatedData VaccinatedData
	json.Unmarshal(bodyBytes, &vaccinatedData)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	//fmt.Printf("%+v", vaccinatedData)
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	//fmt.Println(len(vaccinatedData.All))
	fmt.Println(vaccinatedData.All)

}
func main() {
	get()
}
