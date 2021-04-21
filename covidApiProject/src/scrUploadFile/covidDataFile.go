package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"

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
	confirmed    int       `json:"confirmed"`
	recovered    int       `json:"recovered"`
	deaths       int       `json:"deaths"`
	country      string    `json:"country"`
	capital_city string    `json:"capital_city"`
	updated      time.Time `json:"updated"`
}

type Details struct {
	Confirmed    float64 `json:"confirmed"`
	Recovered    float64 `json:"recovered"`
	Deaths       float64 `json:"deaths"`
	Country      string  `json:"country"`
	Capital_City string  `json:"capital_city"`
	//Updated      string  `json:"updated"`
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
	v.GetObject().Visit(func(k []byte, v *fastjson.Value) {
		//fmt.Printf("key=%s, value=%s\n", k,v)
		keyValues = append(keyValues, string(k))
		//fmt.Printf("%s", k)
		//fmt.Println(" ")
		//fmt.Println(v)

	})
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
		//values := CountryData{Confirmed: info.Countries.All.Confirmed}
		fmt.Println("@@@@@@@@@@@@@@@@@@@")
		//fmt.Println(values)

	}

}
func getValues() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		log.Fatalln(err)
	}
	res := []Details{}

	//web, _ := template.ParseFiles("html\\webPage.html")

	//file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//log.Fatalf("failed creating file: %s", err)
	//}
	//datawriter := bufio.NewWriter(file)

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
	//var res []string
	//var res []Details
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
				details := Details{}
				var confirmed float64
				var recovered float64
				var deaths float64
				var country string
				var capital_city string
				//var updated string
				//fmt.Println("Country:", i)
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
						//country = i
					}
					if k1 == "capital_city" && v1 != nil {
						capital_city = v1.(string)
					}
					// if k1 == "updated" && v1 != nil {
					// 	updated = v1.(string)
					// }

					//fmt.Println(k1, ":", v1)
					//mapstructure.Decode(v1, &details)
					//details = v1
					//byteData, _ := json.Marshal(v1)
					//fmt.Println(string(byteData))
					//byteStringData := string(byteData)
					//arrayForm := strings.Split(byteStringData, ",")
					//	res = append(res, k1+":"+byteStringData)
					//fmt.Println(res)

					//printData := k1
					//details := v1
					//res = append(res, details)
					//fmt.Println(printData, ":", details)
					//web.Execute(w, details)
					//country = i

				}
				details = Details{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country}
				res = append(res, details)

			}
			//fmt.Println(" ")

			//covidCountriesData = append(covidCountriesData, allV)

		} //states values
	}
	fmt.Println(res)
}
func getTesting() {
	resp1, err := http.Get("https://covid-api.mmediagroup.fr/v1/cases")
	if err != nil {
		log.Fatalln(err)
	}

	// file, err := os.OpenFile("test.xlsx", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }
	// datawriter := bufio.NewWriter(file)
	// webPage, err := template.ParseFiles("C:\\Users\\SRS\\gitProject16april\\golang_Practicing\\covidApiProject\\html\\webPage.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	res := []Details{}

	//web, _ := template.ParseFiles("html\\webPage.html")

	//file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//log.Fatalf("failed creating file: %s", err)
	//}
	//datawriter := bufio.NewWriter(file)

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
	//var res []string
	//var res []Details
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
				details := Details{}
				var confirmed float64
				var recovered float64
				var deaths float64
				var country string
				var capital_city string
				//var updated string
				//fmt.Println("Country:", i)
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
					// if k1 == "updated" && v1 != nil {
					// 	updated = v1.(string)
					// }

					//fmt.Println(k1, ":", v1)
					//mapstructure.Decode(v1, &details)
					//details = v1
					//byteData, _ := json.Marshal(v1)
					//fmt.Println(string(byteData))
					//byteStringData := string(byteData)
					//arrayForm := strings.Split(byteStringData, ",")
					//	res = append(res, k1+":"+byteStringData)
					//fmt.Println(res)

					//printData := k1
					//details := v1
					//res = append(res, details)
					//fmt.Println(printData, ":", details)
					//web.Execute(w, details)

				}
				details = Details{Confirmed: confirmed, Recovered: recovered, Deaths: deaths, Capital_City: capital_city, Country: country}
				res = append(res, details)

			}
			//fmt.Println(" ")

			//covidCountriesData = append(covidCountriesData, allV)

		} //states values
	}
	//len := len(res)
	// for _, data := range res {
	// 	//for data := 0; data < len; data++ {
	// 	_, _ = datawriter.WriteString(data + "\n")

	// }
	// datawriter.Flush()
	// file.Close()
	//writeOrders("ordersReport.csv", res)
	//fileCreated("ordersReport.csv", res)
	//text(res)
	createCSVfile(res)

	fmt.Println(res)
	//webPage.Execute(w, res)

}

// func writeOrders(name string, res []Details) {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
// 	}

// 	defer func() {
// 		e := f.Close()
// 		if e != nil {
// 			log.Fatalf("Cannot close '%s': %s\n", name, e.Error())
// 		}
// 	}()

// 	w := csv.NewWriter(f)

//len := len(res)
// for _, data := range res {
// 	for _, rowCell := range data {
// 		//_, _ = datawriter.WriteString(data + "\n")
// 		err = w.WriteAll(rowCell)
// 	}
// }

//}
func fileCreated(name string, res []Details) {
	csvfile, err := os.Create(name)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvfile)
	//	a := 0
	fmt.Println(res[1], res[2])

	// for k, i := range res {

	// 	_ = csvwriter.WriteAll(i)
	// }
	// var result string
	// for k, _ := range res {
	// 	result + k = strings.Join(res[k], " ")
	// }

	for k, i := range res {
		fmt.Println(k, i)
		for k1, i1 := range i.Capital_City {
			fmt.Println(k1, i1)
		}
	}

	csvwriter.Flush()

	csvfile.Close()
}
func text(res []Details) {
	file, err := os.OpenFile("covidREport.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)
	//len := len(res)
	//for _, data := range res {
	// for data := 0; data < len; data++ {
	// 	_, _ = datawriter.Write(res[data])

	// }
	datawriter.Flush()
	file.Close()
}

func createCSVfile(res []Details) {
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
		)

		writer.Write(r)

	}

}

// len := len(res)
// //for _, data := range res {
// for data := 0; data < len; data++ {
// 	_, _ = datawriter.WriteString(res[data] + "\n")

// }
// datawriter.Flush()
// file.Close()

//fmt.Println(res)
// func handlerMethod() {
// 	log.Println("Server started on: http://localhost:8081")
// 	http.HandleFunc("/", getTesting)
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

func main() {
	//handlerMethod()
	//getAll()
	//getValues()
	getTesting()

}
