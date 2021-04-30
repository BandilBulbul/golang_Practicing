package createfile

type CSVFile struct {
	Confirmed    string `json:"confirmed"`
	Recovered    string `json:"recovered"`
	Deaths       string `json:"deaths"`
	Country      string `json:"country"`
	Capital_City string `json:"capital_city"`
	Updated      string `json:"updated"`
}

// func CreateCSVfile(res []CSVFile) { // we can with different entities
// 	file, _ := os.Create("constant\\covidDetailsFile.csv")
// 	//checkError("Error:", err)
// 	defer file.Close()
// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	//define colum headers
// 	headers := []string{"confirmed", "recovered", "deaths", "country", "capital_city", "updated"}

// 	//var ConfirmedString, RecoveredString, DeathsString string
// 	//var ConfirmedString string
// 	//r := make([]string, 0, 1+len(headers))
// 	//r = append(r, "confirmed", "recovered", "deaths", "country", "capital_city", "updated")

// 	for key := range res {
// 		r := make([]string, 0, 1+len(headers))
// 		//ConfirmedString = strconv.Itoa(int(res[key].Confirmed))
// 		//RecoveredString = strconv.Itoa(int(res[key].Recovered))
// 		//DeathsString = strconv.Itoa(int(res[key].Deaths))

// 		r = append(r,
// 			//ConfirmedString,
// 			//RecoveredString,
// 			//DeathsString,
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
