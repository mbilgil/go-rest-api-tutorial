package university

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type UniversityResponseModel []struct {
	Domains       []string    `json:"domains"`
	Country       string      `json:"country"`
	StateProvince interface{} `json:"state-province"`
	WebPages      []string    `json:"web_pages"`
	Name          string      `json:"name"`
	AlphaTwoCode  string      `json:"alpha_two_code"`
}

func GetUniversityByCountry(w http.ResponseWriter, req *http.Request) {
	url := "http://universities.hipolabs.com/search?country={COUNTRY_NAME}"
	new_url := strings.ReplaceAll(url, "{COUNTRY_NAME}", req.FormValue("country"))

	client := &http.Client{}
	req, err := http.NewRequest("GET", new_url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "apikey 6WudO4PJvjsPyAAH8IVN8R:11ZbRiCFeSwryJ8xctO9O5")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject UniversityResponseModel
	json.Unmarshal(bodyBytes, &responseObject)
	json.NewEncoder(w).Encode(responseObject)

}
