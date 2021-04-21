package main

import (
	"encoding/json"
	"fmt"
	"text/template"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/valyala/fastjson"
)

type StatesDetails struct {
	Confirmed    float64 `json:"confirmed"`
	Recovered    float64 `json:"recovered"`
	Deaths       float64 `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	Updated      string  `json:"updated"`
}

func GetIndiaStatesdCovidDetails(w http.ResponseWriter, r *http.Request) {
	//func GetIndiaStatesdCovidDetails() {

	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India") //restApi
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\statesCovidDetails.html")
	if err != nil {
		log.Fatal(err)
	}
	res := []StatesDetails{}   //create Slice of Structure
	details := StatesDetails{} //Array of Struct

	bodyBytes1, _ := ioutil.ReadAll(resp1.Body)
	defer resp1.Body.Close()

	var statesCovidDetails map[string]interface{}   //create map with string as key and interface for values
	json.Unmarshal(bodyBytes1, &statesCovidDetails) //map the values into msg from json

	bodyString1 := string(bodyBytes1) //Convert into String

	//to get the all countries key values
	var p fastjson.Parser //using package for iterate get the key and values
	//May parse array containing values with distinct types (aka non-homogenous types).
	StatesCovidValues, err := p.Parse(bodyString1)
	if err != nil {
		log.Fatal(err)
	}
	//Standard encoding/json is good for the majority of use cases,
	//but it may be quite slow comparing to alternative solutions.
	//If you need performance, try using fastjson.
	//It parses arbitrary JSONs without the need for creating structs
	//or maps matching the JSON schema.
	var keyValues []string //create  slice string
	// Visit all the items in the top object
	StatesCovidValues.GetObject().Visit(func(key []byte, values *fastjson.Value) { //Visit all the items in the top object
		keyValues = append(keyValues, string(key)) //Append into keyValues

	})
	//iterate the keyValues to get inside  values
	for _, i := range keyValues { //i having country name ==key
		var confirmed float64
		var recovered float64
		var deaths float64
		var country string
		var capital_city string
		var updated string
		all := statesCovidDetails[i].(map[string]interface{}) //create another one interface to map with inside valuea and keys
		if i != "All" {                                       // condition should be satisfied
			for key, value := range all {

				//allV := value.(map[string]interface{}) //create another one
				country = i //pass the country name
				//for k1, v1 := range allV {
				if key == "confirmed" && value != nil {
					confirmed = value.(float64)
				}
				if key == "recovered" && value != nil {
					recovered = value.(float64)
				}
				if key == "deaths" && value != nil {
					deaths = value.(float64)
				}
				if key == "country" && value != nil {
					country = value.(string)
				}
				if key == "capital_city" && value != nil {
					capital_city = value.(string)
				}
				if key == "updated" && value != nil {
					updated = value.(string)
				}
				//}
				//append into result slice                                                                                                                  //appending it

			}
			details = StatesDetails{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country, Updated: updated} //save data into detail variable
			res = append(res, details)
		} //states values

	}
	fmt.Println(res) //print result of structure details
	//create csv file
	//createCSVfile(res)
	webPage.Execute(w, res) //return  to  web page

}

// func createCSVfile(res []Details) {
// 	file, _ := os.Create("covidDetailsFile.csv")
// 	//checkError("Error:", err)
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	//define colum headers
// 	headers := []string{
// 		"confirmed",
// 		"recovered",
// 		"deaths",
// 		"country",
// 		"capital_city",
// 		"updated",
// 	}

// 	var ConfirmedString, RecoveredString, DeathsString string

// 	for key := range res {
// 		r := make([]string, 0, 1+len(headers))
// 		ConfirmedString = strconv.Itoa(int(res[key].Confirmed))
// 		RecoveredString = strconv.Itoa(int(res[key].Recovered))
// 		DeathsString = strconv.Itoa(int(res[key].Deaths))

// 		r = append(r,
// 			ConfirmedString,
// 			RecoveredString,
// 			DeathsString,
// 			res[key].Country,
// 			res[key].Capital_City,
// 			res[key].Updated,
// 		)

// 		writer.Write(r)

// 	}

// }

func handlerMethod() {
	log.Println("Server started on: http://localhost:8066")
	http.HandleFunc("/", GetIndiaStatesdCovidDetails)
	log.Fatal(http.ListenAndServe(":8066", nil))
}

func main() {
	handlerMethod()
	//GetIndiaStatesdCovidDetails()
}
