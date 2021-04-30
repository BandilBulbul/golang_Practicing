package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type AllInfo struct {
	All    CountryData `json:"All"`
	States map[string]string
}

type CountryData struct {
	Confirmed           int       `json:"confirmed"`
	Recovered           int       `json:"recovered"`
	Deaths              int       `json:"deaths"`
	Country             string    `json:"country"`
	Population          int       `json:"population"`
	Sq_Km_Area          int       `json:"sq_km_area"`
	LifeExpectancy      string    `json:"life_expectancy"`
	Elevation_In_Meters int       `json:"elevation_in_meters"`
	Continent           string    `json:"continent"`
	Abbreviation        string    `json:"abbreviation"`
	Location            string    `json:"location"`
	ISO                 int       `json:"iso"`
	Capital_City        string    `json:"capital_city"`
	Lat                 string    `json:"lat"`
	Long                string    `json:"long"`
	Updated             time.Time `json:"updated"`
}

type State struct {
	Lat       string    `json:"lat"`
	Long      string    `json:"long"`
	Confirmed int       `json:"confirmed"`
	Recovered int       `json:"recovered"`
	Deaths    int       `json:"deaths"`
	Updated   time.Time `json:"updated"`
}

func get(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//bodyString := string(bodyBytes)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var all AllInfo
	json.Unmarshal(bodyBytes, &all)
	values := CountryData{Confirmed: all.All.Confirmed, Recovered: all.All.Recovered, Deaths: all.All.Deaths, Country: all.All.Country, Capital_City: all.All.Capital_City, Updated: all.All.Updated}
	p, err := template.ParseFiles("html\\covidWriteFile.html")
	if err != nil {
		log.Fatal(err)
	}
	p.Execute(w, values)
}
func showFrontPage(w http.ResponseWriter, r *http.Request) {
	p, err := template.ParseFiles("html\\frontPage.html")
	if err != nil {
		log.Fatal(err)
	}
	values := "hello"
	p.Execute(w, values)

}

func handlerMethod() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", get)
	http.HandleFunc("/covid", showFrontPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handlerMethod()
}

//testing.go
