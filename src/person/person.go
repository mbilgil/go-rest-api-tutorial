package person

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Router    *mux.Router
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

// func (a *Person) PersonEndpoints() {
// 	a.Router = mux.NewRouter()
// 	people = append(people, Person{ID: "1", FirstName: "Mehmet", LastName: "Bilgil", Address: &Address{City: "İstanbul", State: "Türkiye"}})
// 	people = append(people, Person{ID: "2", FirstName: "Hikmet", LastName: "Bilgil", Address: &Address{City: "Hatay", State: "Türkiye"}})
// 	a.Router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
// 	a.Router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
// 	a.Router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
// 	a.Router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
// }
