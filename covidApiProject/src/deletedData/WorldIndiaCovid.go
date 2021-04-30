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

	"math"

	"github.com/bradfitz/slice"
	"github.com/valyala/fastjson"
)

type Details struct {
	Confirmed_Id float64 `json:"confirmed_id"`
	Confirmed    string  `json:"confirmed"`
	Recovered    string  `json:"recovered"`
	Deaths       string  `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	Updated      string  `json:"updated"`
}
type StatesDetails struct {
	Confirmed    float64 `json:"confirmed"`
	Recovered    float64 `json:"recovered"`
	Deaths       float64 `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	Updated      string  `json:"updated"`
}
type VaccinatationCountryDetails struct {
	//Administered                float64
	People_Vaccinated string
	//People_Partially_Vaccinated float64
	Country       string
	Population    string
	Population_Id float64
}
type CountryStatesDetails struct {
	Confirmed     float64 `json:"confirmed"`
	Recovered     float64 `json:"recovered"`
	Deaths        float64 `json:"deaths"`
	Country       string  `json:"country"`
	State_Capital string  `json:"capital_city"`
	Updated       string  `json:"updated"`
}
type URLStruct struct {
	Url string
}

func readUrl() string {
	urlFile, err := ioutil.ReadFile("file.json")
	if err != nil {
		log.Print(err)
	}

	var urlsite URLStruct
	err = json.Unmarshal(urlFile, &urlsite)
	if err != nil {
		log.Println("error:", err)
	}
	return urlsite.Url

}
func getWorldCovidDetails(w http.ResponseWriter, r *http.Request) {
	url := readUrl()
	//resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases") //restApi
	resp1, err := http.Get(url) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	//hardcoded
	//webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\webPage.html")
	webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\worldnew.html")
	if err != nil {
		log.Fatal(err)
	}
	res := []Details{}   //create Slice of Structure
	details := Details{} //Array of Struct

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
	//but it may be quite slow comparing to alternative solutions
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
			if keyy == "All" { // condition should be satisfied
				allV := value.(map[string]interface{}) //create another one
				var confirmed_id float64
				var confirmed string
				var recovered string
				var deaths string
				var country string
				var capital_city string
				var updated string
				country = i //pass the country name
				for k1, v1 := range allV {
					if k1 == "confirmed" && v1 != nil {
						//confirmed = v1.(float64)
						// s = strconv.FormatFloat(confirmed, 'f', 0, 64)
						confirmed = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}
					if k1 == "recovered" && v1 != nil {
						//recovered = v1.(float64)
						recovered = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}
					if k1 == "deaths" && v1 != nil {
						//deaths = v1.(float64)
						deaths = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}
					if k1 == "country" && v1 != nil {
						country = v1.(string)
					}
					if k1 == "capital_city" && v1 != nil {
						capital_city = v1.(string)
					}
					if k1 == "updated" && v1 != nil {
						updated = v1.(string)
					}
					if k1 == "confirmed" && v1 != nil {
						confirmed_id = v1.(float64)
						// s = strconv.FormatFloat(confirmed, 'f', 0, 64)
						//confirmed = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}
				}
				details = Details{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country, Updated: updated, Confirmed_Id: confirmed_id} //save data into detail variable
				res = append(res, details)
				//append into result slice                                                                                                                  //appending it

			}

		} //states values
	}
	slice.Sort(res, func(i, j int) bool {
		return res[i].Confirmed_Id > res[j].Confirmed_Id
	})

	//fmt.Println(res) //print result of structure details
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

	//var ConfirmedString, RecoveredString, DeathsString string
	//var ConfirmedString string

	for key := range res {
		r := make([]string, 0, 1+len(headers))
		//ConfirmedString = strconv.Itoa(int(res[key].Confirmed))
		//RecoveredString = strconv.Itoa(int(res[key].Recovered))
		//DeathsString = strconv.Itoa(int(res[key].Deaths))

		r = append(r,
			//ConfirmedString,
			//RecoveredString,
			//DeathsString,
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

func getIndiaStatesdCovidDetails(w http.ResponseWriter, r *http.Request) {
	//func GetIndiaStatesdCovidDetails() {

	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases?country=India") //restApi
	if err != nil {
		log.Fatalln(err)
	}
	webPageStates, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\indianew.html")
	if err != nil {
		log.Fatal(err)
	}
	states := []StatesDetails{} //create Slice of Structure
	details := StatesDetails{}  //Array of Struct

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
			states = append(states, details)
		} //states values

	}
	//fmt.Println(states) //print result of structure details
	//create csv file
	//createCSVfile(res)
	slice.Sort(states, func(i, j int) bool {
		return states[i].Confirmed > states[j].Confirmed
	})
	fmt.Print(states)
	webPageStates.Execute(w, states) //return  to  web page

}

func getWorldVaccinationDetails(w http.ResponseWriter, r *http.Request) {
	//func getWorldVaccinationDetails() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/vaccines")
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\vaccinatedFile.html")
	if err != nil {
		log.Fatal(err)
	}
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
				//var administered float64
				var peopleVaccinated string
				//var peoplePartiallyVaccinated float64
				var country string
				var population_id float64
				var population string
				country = i //pass the country name
				for k1, v1 := range allV {
					// if k1 == "administered" && v1 != nil {
					// 	administered = v1.(float64)
					// }
					if k1 == "people_vaccinated" && v1 != nil {
						//peopleVaccinated = v1.(float64)
						peopleVaccinated = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}
					// if k1 == "people_partially_vaccinated" && v1 != nil {
					// 	peoplePartiallyVaccinated = v1.(float64)
					// }
					if k1 == "country" && v1 != nil {
						country = v1.(string)
					}
					if k1 == "population" && v1 != nil {
						//population = v1.(float64)
						population = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}
					//if k1 == "population" && v1 != nil {
					if k1 == "people_vaccinated" && v1 != nil {
						population_id = v1.(float64)
						//population = strconv.FormatFloat(v1.(float64), 'f', 0, 64)

					}

				}
				details = VaccinatationCountryDetails{People_Vaccinated: peopleVaccinated, Country: country, Population: population, Population_Id: population_id} //save data into detail variable
				res = append(res, details)                                                                                                                         //append into result slice                                                                                                                  //appending it

			}
		} //states values
	}
	//fmt.Println(res) //print result of structure details
	//create csv file
	//createCSVfile(res)
	slice.Sort(res, func(i, j int) bool {
		return res[i].Population_Id > res[j].Population_Id
	})
	webPage.Execute(w, res) //return  to  web page
}

// type CountryNamePrint struct {
// 	Name string
// }

func getWorldStatesCovidDetails(w http.ResponseWriter, r *http.Request) {
	//func getWorldCovidDetails() {
	CountryUrl := r.URL.Query().Get("country")
	fmt.Print(CountryUrl)
	//countryNamePrint := CountryNamePrint{Name: CountryUrl}

	//url := "https://covid-api.mmediagroup.fr/v1/cases?country=" + countryUrl
	//fmt.Print(url)
	url := "https://covid-api.mmediagroup.fr/v1/cases"

	resp1, err := http.Get(url) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\countrystatesnew.html")
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
			//if keyy != "All" && i == countryUrl { // condition should be satisfied
			if i == CountryUrl { // condition should be satisfied
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
						confirmed = math.Floor(v1.(float64))
					}
					if k1 == "recovered" && v1 != nil {
						recovered = math.Round(v1.(float64))
					}
					if k1 == "deaths" && v1 != nil {
						deaths = math.Round(v1.(float64))
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
	//webPage.Execute(w, countryNamePrint)
	webPage.Execute(w, res) //return  to  web page

}

func home(w http.ResponseWriter, r *http.Request) {
	webPageHome, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\homenew.html")
	if err != nil {
		log.Fatal(err)
	}
	webPageHome.Execute(w, "Home")

}
func handlerMethod() {
	log.Println("Server started on: http://localhost:8065")
	http.HandleFunc("/indiaStates/", getIndiaStatesdCovidDetails)
	http.HandleFunc("/world/", getWorldCovidDetails)
	http.HandleFunc("/", home)
	http.HandleFunc("/worldDetails", getWorldCovidDetails)
	http.HandleFunc("/indiaDetails", getIndiaStatesdCovidDetails)
	http.HandleFunc("/vaccinated", getWorldVaccinationDetails)
	http.HandleFunc("/states", getWorldStatesCovidDetails)

	log.Fatal(http.ListenAndServe(":8065", nil))
}

func main() {
	handlerMethod()
}
