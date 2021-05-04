package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type URLStruct struct {
	Url        string
	UrlVaccine string
	UrlIndia   string
}

// type CSVFile struct {
// 	Confirmed    string `json:"confirmed"`
// 	Recovered    string `json:"recovered"`
// 	Deaths       string `json:"deaths"`
// 	Country      string `json:"country"`
// 	Capital_City string `json:"capital_city"`
// 	Updated      string `json:"updated"`
// }

func ReadUrl() URLStruct {
	urlFile, err := ioutil.ReadFile("constant\\file.json")
	if err != nil {
		log.Print(err)
	}

	var urlsite URLStruct
	err = json.Unmarshal(urlFile, &urlsite)
	if err != nil {
		log.Println("error:", err)
	}
	return urlsite

}

// func CreateCSVfile(res []CSVFile) {
// 	file, _ := os.Create("constant\\covidDetailsFile.csv")
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	//define colum headers
// 	headers := []string{"confirmed", "recovered", "deaths", "country", "capital_city", "updated"}

// 	for key := range res {
// 		r := make([]string, 0, 1+len(headers))
// 		r = append(r,
// 			res[key].Confirmed,
// 			res[key].Recovered,
// 			res[key].Deaths,
// 			res[key].Country,
// 			res[key].Capital_City,
// 			res[key].Updated,
// 		)
// 		writer.Write(r)
// 	}
// }
