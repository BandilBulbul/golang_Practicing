package worldstatescovid

import (
	"encoding/json"
	"math"
	"strconv"

	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"covid19/constant"
	"covid19/util"

	"github.com/bradfitz/slice"
	"github.com/valyala/fastjson"
)

type CountryStatesDetails struct {
	Confirmed_Id  float64 `json:"confirmed_id"`
	Confirmed     string  `json:"confirmed"`
	Recovered     string  `json:"recovered"`
	Deaths        string  `json:"deaths"`
	Country       string  `json:"country"`
	State_Capital string  `json:"capital_city"`
	Updated       string  `json:"updated"`
}

func GetWorldStatesCovidDetails(w http.ResponseWriter, r *http.Request) {
	CountryUrl := r.URL.Query().Get("country")
	restUrl := util.ReadUrl()
	response, err := http.Get(restUrl.Url) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles(constant.StatesTemplate)
	if err != nil {
		log.Fatal(err)
	}
	statesDetails := []CountryStatesDetails{} //create Slice of Structure

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
	for _, countryName := range keyValues { //i having country name ==key
		all := msg[countryName].(map[string]interface{}) //create another one interface to map with inside valuea and keys
		for country_key, country_value := range all {
			//if keyy != "All" && i == countryUrl { // condition should be satisfied
			if countryName == CountryUrl { // condition should be satisfied
				allValues := country_value.(map[string]interface{}) //create another one
				details := CountryStatesDetails{}                   //Array of Struct
				var confirmed_id float64
				var confirmed string
				var recovered string
				var deaths string
				var country string
				var state_capital string
				var updated string
				country = countryName //pass the country name
				state_capital = country_key
				for dataKey, dataValue := range allValues {
					if dataKey == constant.ConfirmedKey && dataValue != nil {
						confirmed = strconv.FormatFloat(dataValue.(float64), 'f', 0, 64)
					}
					if dataKey == constant.RecoveredKey && dataValue != nil {
						recovered = strconv.FormatFloat(dataValue.(float64), 'f', 0, 64)
					}
					if dataKey == constant.DeathsKey && dataValue != nil {
						deaths = strconv.FormatFloat(dataValue.(float64), 'f', 0, 64)
					}
					if dataKey == constant.UpdatedKey && dataValue != nil {
						updated = dataValue.(string)
					}
					if dataKey == constant.ConfirmedKey && dataValue != nil {
						confirmed_id = math.Floor(dataValue.(float64))
					}
				}
				details = CountryStatesDetails{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, State_Capital: state_capital, Country: country, Updated: updated, Confirmed_Id: confirmed_id} //save data into detail variable
				statesDetails = append(statesDetails, details)                                                                                                                                           //append into result slice                                                                                                                  //appending it

			}
		} //states values
	}

	slice.Sort(statesDetails, func(i, j int) bool { //sort the slice
		return statesDetails[i].Confirmed_Id > statesDetails[j].Confirmed_Id
	})
	webPage.Execute(w, statesDetails) //return  to  web page

}
