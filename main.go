package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type UniversityResponseModel []struct {
	Domains       []string    `json:"domains"`
	Country       string      `json:"country"`
	StateProvince interface{} `json:"state-province"`
	WebPages      []string    `json:"web_pages"`
	Name          string      `json:"name"`
	AlphaTwoCode  string      `json:"alpha_two_code"`
}

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func QueryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "query: "+req.FormValue("name"))
	io.WriteString(res, "\nphone: "+req.FormValue("phone"))
	println("Enter this in your browser:  http://localhost:9000/example?name=mehmet&phone=533-533")
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

func main() {
	router := mux.NewRouter().StrictSlash(true)
	people = append(people, Person{ID: "1", FirstName: "Mehmet", LastName: "Bilgil", Address: &Address{City: "İstanbul", State: "Türkiye"}})
	people = append(people, Person{ID: "2", FirstName: "Hikmet", LastName: "Bilgil", Address: &Address{City: "Hatay", State: "Türkiye"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/example", QueryParamDisplayHandler).Methods("GET")
	router.HandleFunc("/university", GetUniversityByCountry).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}
