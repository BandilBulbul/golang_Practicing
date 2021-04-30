package vaccine

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
	//Administered                float64
	People_Vaccinated string
	//People_Partially_Vaccinated float64
	Country       string
	Population    string
	Population_Id float64
}

func GetWorldVaccinationDetails(w http.ResponseWriter, r *http.Request) {
	restUrl := util.ReadUrl().UrlVaccine
	response, err := http.Get(restUrl)
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles(constant.VaccineTemplate)
	if err != nil {
		log.Fatal(err)
	}
	vaccineDatails := []VaccinatationCountryDetails{} //create Slice of Structure

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var msg map[string]interface{}  //create map with string as key and interface for values
	json.Unmarshal(bodyBytes, &msg) //map the values into msg from json

	bodyString := string(bodyBytes) //Convert into String

	//to get the all countries key values
	var p fastjson.Parser //using package for iterate get the key and values
	//May parse array containing values with distinct types (aka non-homogenous types).
	v, err := p.Parse(bodyString)
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
			if keyy == constant.ALLKey && i != constant.GlobalKey && i != constant.WorldKey { // condition should be satisfied
				allV := value.(map[string]interface{})                       //create another one
				vaccinatationCountryDetails := VaccinatationCountryDetails{} //Array of Struct
				var peopleVaccinated string
				var country, population string
				var population_id float64
				country = i //pass the country name
				for k1, v1 := range allV {
					if k1 == constant.People_VaccinatedKey && v1 != nil {
						peopleVaccinated = strconv.FormatFloat(v1.(float64), 'f', 0, 64)
					}
					if k1 == constant.CountryKey && v1 != nil {
						country = v1.(string)
					}
					if k1 == constant.PopulationKey && v1 != nil {
						population = strconv.FormatFloat(v1.(float64), 'f', 0, 64)
					}
					if k1 == constant.People_VaccinatedKey && v1 != nil {
						population_id = v1.(float64)
					}
				}
				vaccinatationCountryDetails = VaccinatationCountryDetails{People_Vaccinated: peopleVaccinated, Country: country, Population: population, Population_Id: population_id} //save data into detail variable
				vaccineDatails = append(vaccineDatails, vaccinatationCountryDetails)                                                                                                   //append into result slice                                                                                                                  //appending it

			}
		} //states values
	}
	slice.Sort(vaccineDatails, func(i, j int) bool {
		return vaccineDatails[i].Population_Id > vaccineDatails[j].Population_Id
	})
	webPage.Execute(w, vaccineDatails) //return  to  web page
}
