package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/valyala/fastjson"
)

type CountryStatesDetails struct {
	Confirmed     float64 `json:"confirmed"`
	Recovered     float64 `json:"recovered"`
	Deaths        float64 `json:"deaths"`
	Country       string  `json:"country"`
	State_Capital string  `json:"capital_city"`
	Updated       string  `json:"updated"`
}

func getWorldStatesCovidDetails(w http.ResponseWriter, r *http.Request) {
	//func getWorldCovidDetails() {
	countryUrl := r.URL.Query().Get("country")

	url := "https://covid-api.mmediagroup.fr/v1/cases?country=" + countryUrl
	resp1, err := http.Get(url) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\countryStatesWise.html")
	if err != nil {
		log.Fatal(err)
	}
	res := []CountryStatesDetails{} //create Slice of Structure

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
	//Standard encoding/json is good for the majority of use cases,
	//but it may be quite slow comparing to alternative solutions.
	//If you need performance, try using fastjson.
	//It parses arbitrary JSONs without the need for creating structs
	//or maps matching the JSON schema.
	var keyValues []string //create  slice string
	// Visit all the items in the top object
	v.GetObject().Visit(func(k []byte, v *fastjson.Value) { //Visit all the items in the top object
		keyValues = append(keyValues, string(k)) //Append into keyValues

	})
	//iterate the keyValues to get inside  values
	for _, i := range keyValues { //i having country name ==key
		all := msg[i].(map[string]interface{}) //create another one interface to map with inside valuea and keys
		for keyy, value := range all {
			if keyy != "All" && i == countryUrl { // condition should be satisfied
				allV := value.(map[string]interface{}) //create another one
				details := CountryStatesDetails{}      //Array of Struct
				var confirmed float64
				var recovered float64
				var deaths float64
				var country string
				var state_capital string
				var updated string
				country = i //pass the country name
				state_capital = keyy
				//fmt.Println(keyy)
				for k1, v1 := range allV {
					if k1 == "confirmed" && v1 != nil {
						confirmed = v1.(float64)
					}
					if k1 == "recovered" && v1 != nil {
						recovered = v1.(float64)
					}
					if k1 == "deaths" && v1 != nil {
						deaths = v1.(float64)
					}
					// if k1 == "country" && v1 != nil {
					// 	//country = v1.(string)
					// 	country=i
					// }
					// if k1 == "capital_city" && v1 != nil {
					// 	capital_city = v1.(string)
					// }
					if k1 == "updated" && v1 != nil {
						updated = v1.(string)
					}
				}
				details = CountryStatesDetails{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, State_Capital: state_capital, Country: country, Updated: updated} //save data into detail variable
				res = append(res, details)                                                                                                                                   //append into result slice                                                                                                                  //appending it

			}
		} //states values
	}
	fmt.Println(res) //print result of structure details

	var countryNameWithHavingStates []string
	for i := 1; i < len(res); i++ {
		if res[i].Country != res[i-1].Country {
			//fmt.Println(res[i].Country)
			countryNameWithHavingStates = append(countryNameWithHavingStates, res[i].Country)

		}
	}
	// var countryNameWithHavingStates []string
	// for _, i := range res {
	// 	if countryNameWithHavingStates == nil {
	// 		countryNameWithHavingStates = append(countryNameWithHavingStates, i.Country)

	// 	}
	// 	fmt.Print(i.Country)
	// 	for k, _ := range countryNameWithHavingStates {
	// 		if countryNameWithHavingStates[k] != i.Country {
	// 			countryNameWithHavingStates = append(countryNameWithHavingStates, i.Country)
	// 		}

	// 	}
	//}
	//fmt.Println(countryNameWithHavingStates)
	//fmt.Println(len(countryNameWithHavingStates))
	//create csv file
	//createCSVfile(res)
	webPage.Execute(w, res) //return  to  web page

}

func createCSVfile(res []Details) {
	file, _ := os.Create("covidDetailsFile.csv")
	//checkError("Error:", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//define colum headers
	headers := []string{
		"confirmed",
		"recovered",
		"deaths",
		"country",
		"capital_city",
		"updated",
	}

	var ConfirmedString, RecoveredString, DeathsString string

	for key := range res {
		r := make([]string, 0, 1+len(headers))
		ConfirmedString = strconv.Itoa(int(res[key].Confirmed))
		RecoveredString = strconv.Itoa(int(res[key].Recovered))
		DeathsString = strconv.Itoa(int(res[key].Deaths))

		r = append(r,
			ConfirmedString,
			RecoveredString,
			DeathsString,
			res[key].Country,
			//res[key].Capital_City,
			res[key].Updated,
		)

		writer.Write(r)

	}

}

// func handlerMethod() {
// 	log.Println("Server started on: http://localhost:8067")
// 	http.HandleFunc("/", getWorldCovidDetails)
// 	log.Fatal(http.ListenAndServe(":8067", nil))
// }

func main() {
	//handlerMethod()
	getWorldCovidDetails()
}

//[Belgium Brazil Canada Chile China Colombia Denmark France Germany India Italy Japan Mexico Netherlands Pakistan Peru Russia
//Spain Sweden Ukraine United Kingdom US]
