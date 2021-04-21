package main

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fastjson"
)


type Details struct {
	Confirmed    int       `json:"confirmed"`
	Recovered    int       `json:"recovered"`
	Deaths       int       `json:"deaths"`
	Country      string    `json:"country"`
	Capital_City string    `json:"capital_city"`
	Updated      time.Time `json:"updated"`
}
func getVAlues() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp1.Body.Close()

	bodyBytes1, _ := ioutil.ReadAll(resp1.Body)
	bodyString1 := string(bodyBytes1)



	var msg map[string]interface{}
	json.Unmarshal(bodyBytes1, &msg)


	var p fastjson.Parser
	v, err := p.Parse(bodyString1)
	if err != nil {
		log.Fatal(err)
	}


	var keyValues []string
	// Visit all the items in the top object
	v.GetObject().Visit(func(k []byte, v *fastjson.Value) {
		//fmt.Printf("key=%s, value=%s\n", k,v)
		keyValues = append(keyValues, string(k))

	})
	res := []Details{}

	for _, i := range keyValues {
		//fmt.Println(msg[i])//countries values
		all := msg[i].(map[string]interface{})
		for keyy, value := range all {
			if keyy == "All" {
				allV := value.(map[string]interface{})

				fmt.Println("Country:", i)
				for k1, v1 := range allV {
					if k1 == "confirmed" || k1 == "recovered" || k1 == "deaths" || k1 == "country" || k1 == "capital_city" || k1 == "updated" {
						printData := k1
						details := v1
						fmt.Println(printData, ":", details)
					}
				}
				fmt.Println(" ")
			}
		}
	}
	fmt.Println(res)

}

func main(){
	getVAlues()
}