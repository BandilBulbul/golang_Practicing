package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func createfile() {
	file, _ := os.Create("export.csv")
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
	}

	var idString string

	for key := range res {
		r := make([]string, 0, 1+len(headers))
		ConfirmedString = strconv.Itoa(m[key].Confirmed)
		RecoveredString = strconv.Itoa(m[key].Recovered)
		DeathsString = strconv.Itoa(m[key].Deaths)

		r = append(r,
			ConfirmedString,
			RecoveredString,
			DeathsString,
			m[key].Country,
			m[key].Capital_City,
		)

		writer.Write(r)

	}

}
