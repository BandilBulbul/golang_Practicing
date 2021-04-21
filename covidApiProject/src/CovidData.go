package main

import (
	"bufio"
	"encoding/json"
	"os"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fastjson"
)

type Info struct {
	Countries AllInfo `json:"countries"`
}

type AllInfo struct {
	All CountryData `json:"All"`
}

type CountryData struct {
	Confirmed           int       `json:"confirmed"`
	Recovered           int       `json:"recovered"`
	Deaths              int       `json:"deaths"`
	Country             string    `json:"country"`
	Population          int       `json:"population"`
	Sq_Km_Area          int       `json:"sq_km_area"`
	LifeExpectancy      string    `json:"life_expectancy"`
	Elevation_In_Meters int       `json:"elevation_in_meters"`
	Continent           string    `json:"continent"`
	Abbreviation        string    `json:"abbreviation"`
	Location            string    `json:"location"`
	ISO                 int       `json:"iso"`
	Capital_City        string    `json:"capital_city"`
	Lat                 string    `json:"lat"`
	Long                string    `json:"long"`
	Updated             time.Time `json:"updated"`
}

type Details struct {
	Confirmed    int       `json:"confirmed"`
	Recovered    int       `json:"recovered"`
	Deaths       int       `json:"deaths"`
	Country      string    `json:"country"`
	Capital_City string    `json:"capital_city"`
	Updated      time.Time `json:"updated"`
}

func getAll() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes1, _ := ioutil.ReadAll(resp1.Body)
	//fmt.Println(resp1.Body)
	bodyString1 := string(bodyBytes1)
	var p fastjson.Parser
	v, err := p.Parse(bodyString1)
	if err != nil {
		log.Fatal(err)
	}
	var keyValues []string
	// Visit all the items in the top object
	//v.GetObject().Visit(func(k []byte, v *fastjson.Value) {
	//fmt.Printf("key=%s, value=%s\n", k,v)
	//keyValues = append(keyValues, string(k))
	//fmt.Printf("%s", k)
	//fmt.Println(" ")
	//fmt.Println(v)

	//})
	//fmt.Println(bodyString1)
	defer resp1.Body.Close()
	//fmt.Println(keyValues)
	len := len(keyValues)

	for i := 0; i < len; i++ {
		type Info struct {
			Countries AllInfo `json:"KeyValues[i]"`
		}
		fmt.Println(keyValues[i])
		var info Info
		json.Unmarshal(bodyBytes1, &info)
		values := CountryData{Confirmed: info.Countries.All.Confirmed}
		fmt.Println("@@@@@@@@@@@@@@@@@@@")
		fmt.Println(values)

	}

}
func getVAlues() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		log.Fatalln(err)
	}
	//web, _ := template.ParseFiles("html\\webPage.html")

	file, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)

	//f, _ := os.Create("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\src\\covidFile1.txt")
	//defer f.Close()

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
	// Visit all the items in the top object
	v.GetObject().Visit(func(k []byte, v *fastjson.Value) {
		//fmt.Printf("key=%s, value=%s\n", k,v)
		keyValues = append(keyValues, string(k))
		//fmt.Printf("%s", k)
		//fmt.Println(" ")
		//fmt.Println(v)

	})
	var res []string
	//var covidData []string
	//var covidCountriesData map[string]CountryData
	for _, i := range keyValues {
		//fmt.Println(msg[i])//countries values
		all := msg[i].(map[string]interface{})
		for keyy, value := range all {
			if keyy == "All" {
				//fmt.Println(keyy)
				//fmt.Println("#################")
				//fmt.Println(value.(map[string]interface{}))
				allV := value.(map[string]interface{})
				//allV := value.(map[string]CountryData)

				//fmt.Println(allV)
				//fmt.Println(reflect.TypeOf(allV))
				//data, _ := json.Marshal(allV)
				//fmt.Println(string(data))
				//dataString := string(data)
				//fmt.Println(dataString)
				//details := Details{}

				fmt.Println("Country:", i)
				for k1, v1 := range allV {
					if k1 == "confirmed" || k1 == "recovered" || k1 == "deaths" || k1 == "country" || k1 == "capital_city" || k1 == "updated" {
						//fmt.Println(k1, ":", v1)
						//mapstructure.Decode(v1, &details)
						byteData, _ := json.Marshal(v1)
						fmt.Println(string(byteData))
						byteStringData := string(byteData)
						//arrayForm := strings.Split(byteStringData, ",")
						res = append(res, k1+":"+byteStringData)

						//printData := k1
						//details := v1
						//res = append(res, details)
						//fmt.Println(printData, ":", details)
						//web.Execute(w, details)

					}

				}
				fmt.Println(" ")

				//covidCountriesData = append(covidCountriesData, allV)

			} //states values
		}
	}
	len := len(res)
	//for _, data := range res {
	for data := 0; data < len; data++ {
		_, _ = datawriter.WriteString(res[data] + "\n")

	}
	datawriter.Flush()
	file.Close()

	//fmt.Println(res)

}

// func handlerMethod() {
// 	log.Println("Server started on: http://localhost:8090")
// 	http.HandleFunc("/", getVAlues)
// 	log.Fatal(http.ListenAndServe(":8091", nil))
// }

func main() {
	//handlerMethod()
	//getAll()
	getVAlues()
}
