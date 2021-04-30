package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/valyala/fastjson"
)

type VaccinatationCountryDetails struct {
	Administered                float64
	People_Vaccinated           float64
	People_Partially_Vaccinated float64
	Country                     string
	Population                  float64
}

//func getWorldVaccinationDetails(w http.ResponseWriter, r *http.Request) {
func getWorldVaccinationDetails() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/vaccines")
	if err != nil {
		log.Fatalln(err)
	}
	// webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\webPage.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	res := []VaccinatationCountryDetails{} //create Slice of Structure

	bodyBytes1, _ := ioutil.ReadAll(resp1.Body)
	defer resp1.Body.Close()

	var msg map[string]interface{}   //create map with string as key and interface for values
	json.Unmarshal(bodyBytes1, &msg) //map the values into msg from json

	bodyString1 := string(bodyBytes1) //Convert into String

	//to get the all countries key values
	var p fastjson.Parser //using package for iterate get the key and values
	//May parse array containing values with distinct types (aka non-homogenous types).
	v, err := p.Parse(bodyString1)
	if err != nil {
		log.Fatal(err)
	}
	var keyValues []string //create  slice string
	// Visit all the items in the top object
	v.GetObject().Visit(func(k []byte, v *fastjson.Value) { //Visit all the items in the top object
		keyValues = append(keyValues, string(k)) //Append into keyValues

	})
	//iterate the keyValues to get inside  values
	for _, i := range keyValues { //i having country name ==key
		all := msg[i].(map[string]interface{}) //create another one interface to map with inside valuea and keys
		for keyy, value := range all {
			if keyy == "All" { // condition should be satisfied
				allV := value.(map[string]interface{})   //create another one
				details := VaccinatationCountryDetails{} //Array of Struct
				var administered float64
				var peopleVaccinated float64
				var peoplePartiallyVaccinated float64
				var country string
				var population float64
				country = i //pass the country name
				for k1, v1 := range allV {
					if k1 == "administered" && v1 != nil {
						administered = v1.(float64)
					}
					if k1 == "people_vaccinated" && v1 != nil {
						peopleVaccinated = v1.(float64)
					}
					if k1 == "people_partially_vaccinated" && v1 != nil {
						peoplePartiallyVaccinated = v1.(float64)
					}
					if k1 == "country" && v1 != nil {
						country = v1.(string)
					}

					if k1 == "population" && v1 != nil {
						population = v1.(float64)
					}

				}
				details = VaccinatationCountryDetails{Administered: administered, People_Vaccinated: peopleVaccinated, People_Partially_Vaccinated: peoplePartiallyVaccinated, Population: population, Country: country} //save data into detail variable
				res = append(res, details)                                                                                                                                                                               //append into result slice                                                                                                                  //appending it

			}
		} //states values
	}
	fmt.Println(res) //print result of structure details
	//create csv file
	//createCSVfile(res)
	//webPage.Execute(w, res) //return  to  web page

}
func main() {
	getWorldVaccinationDetails()
}
