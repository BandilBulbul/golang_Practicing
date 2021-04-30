package worldcovid

import (
	"covid19/constant"
	"covid19/util"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/bradfitz/slice"
	"github.com/valyala/fastjson"
)

type WorldDetails struct {
	Confirmed_Id float64 `json:"confirmed_id"`
	Confirmed    string  `json:"confirmed"`
	Recovered    string  `json:"recovered"`
	Deaths       string  `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	Updated      string  `json:"updated"`
}

func GetWorldCovidDetails(w http.ResponseWriter, r *http.Request) {
	restUrl := util.ReadUrl()
	response, err := http.Get(restUrl.Url) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	//hardcoded
	webPage, err := template.ParseFiles(constant.WorldTemplate)
	if err != nil {
		log.Fatal(err)
	}
	CountryDetails := []WorldDetails{} //create Slice of Structure
	//var CountryDetails []WorldDetails //create Slice of Structure

	worldDetails := WorldDetails{} //Array of Struct

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var countryKeysValues map[string]interface{}  //create map with string as key and interface for values
	json.Unmarshal(bodyBytes, &countryKeysValues) //map the values into msg from json

	bodyString := string(bodyBytes) //Convert into String

	//to get the all countries key values
	var json_iterate fastjson.Parser //using package for iterate get the key and values
	//May parse array containing values with distinct types (aka non-homogenous types).
	values, err := json_iterate.Parse(bodyString)
	if err != nil {
		log.Fatal(err)
	}
	//Standard encoding/json is good for the majority of use cases,
	//but it may be quite slow comparing to alternative solutions
	//If you need performance, try using fastjson.
	//It parses arbitrary JSONs without the need for creating structs or maps matching the JSON schema.
	var keyValues []string //create  slice string
	// Visit all the items in the top object
	values.GetObject().Visit(func(keys []byte, values *fastjson.Value) { //Visit all the items in the top object
		keyValues = append(keyValues, string(keys)) //Append into keyValues

	})
	//iterate the keyValues to get inside  values
	for _, i := range keyValues { //i having country name ==key
		all := countryKeysValues[i].(map[string]interface{}) //create another one interface to map with inside valuea and keys

		for country_key, country_value := range all {
			if country_key == constant.ALLKey { // condition should be satisfied
				all_countryValues := country_value.(map[string]interface{}) //create another one
				var confirmed_id float64
				var confirmed, recovered, deaths, country, capital_city, updated string
				country = i //pass the country name

				for dataKey, dataValue := range all_countryValues {

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
				}
				worldDetails = WorldDetails{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country, Updated: updated, Confirmed_Id: confirmed_id} //save data into detail variable
				CountryDetails = append(CountryDetails, worldDetails)                                                                                                                               //append into result slice                                                                                                                  //appending it
			}

		}
	}
	slice.Sort(CountryDetails, func(i, j int) bool { //sort the slice
		return CountryDetails[i].Confirmed_Id > CountryDetails[j].Confirmed_Id
	})
	CreateCSVfile(CountryDetails)      //creating csv file
	webPage.Execute(w, CountryDetails) //return  to  web page
}

func CreateCSVfile(res []WorldDetails) { // we can with different entities
	file, _ := os.Create("constant\\covidDetailsFile.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//define colum headers
	headers := []string{"confirmed", "recovered", "deaths", "country", "capital_city", "updated"}

	for key := range res {
		r := make([]string, 0, 1+len(headers))
		r = append(r,
			res[key].Confirmed,
			res[key].Recovered,
			res[key].Deaths,
			res[key].Country,
			res[key].Capital_City,
			res[key].Updated,
		)
		writer.Write(r)
	}
}
