package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	
)

type Currencylayer struct {
	Success   bool   `json:"success"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		Usdeur float64 `json:"USDEUR"`
		Usdinr float64 `json:"USDINR"`
	} `json:"quotes"`
}


func main() {
	fmt.Println("Starting Application....\n")

	response, err := http.Get("http://apilayer.net/api/live?access_key=f9f65847ef1f170f5566767861fdf62a&currencies=EUR,INR&source=USD")
	if err != nil{
		log.Println(err)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))

	var record Currencylayer
	json.Unmarshal(responseData, &record)

	fmt.Println(record)
}