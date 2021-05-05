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
