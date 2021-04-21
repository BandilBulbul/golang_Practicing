package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/valyala/fastjson"
)

type Details struct {
	Confirmed    float64 `json:"confirmed"`
	Recovered    float64 `json:"recovered"`
	Deaths       float64 `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	Updated      string  `json:"updated"`
}

func getTesting() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		log.Fatalln(err)
	}
	res := []Details{}
	bodyBytes1, _ := ioutil.ReadAll(resp1.Body)
	defer resp1.Body.Close()
	var msg map[string]interface{}
	json.Unmarshal(bodyBytes1, &msg)

	bodyString1 := string(bodyBytes1)
	var p fastjson.Parser
	v, err := p.Parse(bodyString1)
	if err != nil {
		log.Fatal(err)
	}
	var keyValues []string
	v.GetObject().Visit(func(k []byte, v *fastjson.Value) {
		keyValues = append(keyValues, string(k))

	})
	for _, i := range keyValues {
		all := msg[i].(map[string]interface{})
		for keyy, value := range all {
			if keyy == "All" {
				allV := value.(map[string]interface{})
				details := Details{}
				var confirmed float64
				var recovered float64
				var deaths float64
				var country string
				var capital_city string
				var updated string
				country = i
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
					if k1 == "country" && v1 != nil {
						country = v1.(string)
					}
					if k1 == "capital_city" && v1 != nil {
						capital_city = v1.(string)
					}
					if k1 == "updated" && v1 != nil {
						updated = v1.(string)
					}
				}
				details = Details{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country, Updated: updated}
				res = append(res, details)

			}
		} //states values
	}
	createCSVfile(res)
}

func createCSVfile(res []Details) {
	file, _ := os.Create("covidCsvFile.csv")
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
			res[key].Capital_City,
			res[key].Updated,
		)

		writer.Write(r)

	}

}
func main() {
	getTesting()
}
