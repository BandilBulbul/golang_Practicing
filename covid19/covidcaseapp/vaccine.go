package covidcaseapp

import (
	"covid19/constant"
	"covid19/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/bradfitz/slice"
	"github.com/valyala/fastjson"
)

type VaccinatationCountryDetails struct {
	People_Vaccinated string
	Country           string
	Population        string
	Population_Id     float64
}

//get vaccination data
func GetWorldVaccinationDetails(w http.ResponseWriter, r *http.Request) {
	restUrl := util.ReadUrl().UrlVaccine
	response, err := http.Get(restUrl)
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles(constant.VaccineTemplate) //Templates
	if err != nil {
		log.Fatal(err)
	}
	vaccineDatails := []VaccinatationCountryDetails{} //create Slice of Structure

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var vaccineData map[string]interface{}  //create map with string as key and interface for values
	json.Unmarshal(bodyBytes, &vaccineData) //map the values into msg from json

	bodyString := string(bodyBytes) //Convert into String

	//to get the all countries key values
	var json_iterate fastjson.Parser //using package for iterate get the key and values
	//May parse array containing values with distinct types (aka non-homogenous types).
	iterationKeysValues, err := json_iterate.Parse(bodyString)
	if err != nil {
		log.Fatal(err)
	}
	var keyValues []string //create  slice string
	// Visit all the items in the top object
	iterationKeysValues.GetObject().Visit(func(keys []byte, iterationKeysValues *fastjson.Value) { //Visit all the items in the top object
		keyValues = append(keyValues, string(keys)) //Append into keyValues

	})
	//iterate the keyValues to get inside  values
	for _, country_Name := range keyValues { //i having country name ==key
		all := vaccineData[country_Name].(map[string]interface{}) //create another one interface to map with inside valuea and keys
		for keyy, value := range all {
			if keyy == constant.ALLKey && country_Name != constant.GlobalKey && country_Name != constant.WorldKey && country_Name != constant.ContinentalAF && country_Name != constant.ContinentalUS && country_Name != constant.ContinentalEU && country_Name != constant.ContinentalSA && country_Name != constant.ContinentalAS && country_Name != constant.ContinentalNA && country_Name != constant.ContinentalE { // condition should be satisfied
				allV := value.(map[string]interface{})                       //create another one
				vaccinatationCountryDetails := VaccinatationCountryDetails{} //Array of Struct
				var peopleVaccinated string
				var country, population string
				var population_id float64
				country = country_Name //pass the country name
				for datakey, datavalue := range allV {
					if datakey == constant.People_VaccinatedKey && datavalue != nil {
						peopleVaccinated = strconv.FormatFloat(datavalue.(float64), 'f', 0, 64)
					}
					if datakey == constant.CountryKey && datavalue != nil {
						country = datavalue.(string)
					}
					if datakey == constant.PopulationKey {
						//population = strconv.FormatFloat(datavalue.(float64), 'f', 0, 64)
						if datavalue != nil {
							population = strconv.FormatFloat(datavalue.(float64), 'f', 0, 64)
						} else {
							population = "Not Available"
						}
					}
					if datakey == constant.People_VaccinatedKey && datavalue != nil {
						population_id = datavalue.(float64)
					}
				}
				vaccinatationCountryDetails = VaccinatationCountryDetails{People_Vaccinated: peopleVaccinated, Country: country, Population: population, Population_Id: population_id} //save data into detail variable
				vaccineDatails = append(vaccineDatails, vaccinatationCountryDetails)                                                                                                   //append into result slice                                                                                                                  //appending it

			}
		}
	}
	//sorting the slice of structures
	slice.Sort(vaccineDatails, func(i, j int) bool {
		return vaccineDatails[i].Population_Id > vaccineDatails[j].Population_Id
	})
	webPage.Execute(w, vaccineDatails) //return  to  web page
}
