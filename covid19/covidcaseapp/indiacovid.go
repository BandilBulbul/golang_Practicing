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

type StatesDetails struct {
	Confirmed_Id float64 `json:"confirmed_id"`
	Confirmed    string  `json:"confirmed"`
	Recovered    string  `json:"recovered"`
	Deaths       string  `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	Updated      string  `json:"updated"`
}

//get india's states data
func GetIndiaStatesdCovidDetails(w http.ResponseWriter, r *http.Request) {
	restUrl := util.ReadUrl().UrlIndia
	response, err := http.Get(restUrl) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	webPageStates, err := template.ParseFiles(constant.IndiaTemplate)
	if err != nil {
		log.Fatal(err)
	}
	states := []StatesDetails{} //create Slice of Structure
	details := StatesDetails{}  //Array of Struct

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var statesCovidDetails map[string]interface{}  //create map with string as key and interface for values
	json.Unmarshal(bodyBytes, &statesCovidDetails) //map the values into msg from json

	bodyString := string(bodyBytes) //Convert into String

	//to get the all countries key values
	var json_iterate fastjson.Parser //using package for iterate get the key and values
	//May parse array containing values with distinct types (aka non-homogenous types).
	StatesCovidValues, err := json_iterate.Parse(bodyString)
	if err != nil {
		log.Fatal(err)
	}
	var keyValues []string //create  slice string
	// Visit all the items in the top object
	StatesCovidValues.GetObject().Visit(func(key []byte, values *fastjson.Value) { //Visit all the items in the top object
		keyValues = append(keyValues, string(key)) //Append into keyValues

	})
	//iterate the keyValues to get inside  values
	for _, country_Name := range keyValues { //i having country name ==key
		var confirmed_id float64
		var confirmed, recovered, deaths, country, capital_city, updated string
		all := statesCovidDetails[country_Name].(map[string]interface{}) //create another one interface to map with inside valuea and keys
		if country_Name != constant.ALLKey {                             // condition should be satisfied
			for dataKey, dataValue := range all {

				country = country_Name //pass the country name
				//for k1, v1 := range allV {
				if dataKey == constant.ConfirmedKey && dataValue != nil {
					confirmed = strconv.FormatFloat(dataValue.(float64), 'f', 0, 64)
				}
				if dataKey == constant.RecoveredKey && dataValue != nil {
					recovered = strconv.FormatFloat(dataValue.(float64), 'f', 0, 64)
				}
				if dataKey == constant.DeathsKey && dataValue != nil {
					deaths = strconv.FormatFloat(dataValue.(float64), 'f', 0, 64)
				}
				if dataKey == constant.CountryKey && dataValue != nil {
					country = dataValue.(string)
				}
				if dataKey == constant.CapitalCityKey && dataValue != nil {
					capital_city = dataValue.(string)
				}
				if dataKey == constant.UpdatedKey && dataValue != nil {
					updated = dataValue.(string)
				}
				if dataKey == constant.ConfirmedKey && dataValue != nil {
					confirmed_id = dataValue.(float64)
				}
				//}
				//append into result slice                                                                                                                  //appending it

			}
			details = StatesDetails{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country, Updated: updated, Confirmed_Id: confirmed_id} //save data into detail variable
			states = append(states, details)
		} //states values

	}
	// sorting the data according to Confirmed cases
	slice.Sort(states, func(i, j int) bool {
		return states[i].Confirmed_Id > states[j].Confirmed_Id
	})
	webPageStates.Execute(w, states) //return  to  web page

}
