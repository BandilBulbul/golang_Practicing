package worldcovid

import (
	"covid19/constant"
	"covid19/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type AllInfo struct {
	All CovidData `json:"Assam"`
}
type CovidData struct {
	Updated string `json:"updated"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	webPageHome, err := template.ParseFiles(constant.HomePage)
	if err != nil {
		log.Fatal(err)
	}

	restUrl := util.ReadUrl().UrlIndia
	response, err := http.Get(restUrl) //restApi
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var updatesTime AllInfo
	json.Unmarshal(bodyBytes, &updatesTime)
	updatedTime := CovidData{Updated: updatesTime.All.Updated}
	updatedTime.Updated = updatedTime.Updated[0:16]

	webPageHome.Execute(w, updatedTime)

}
