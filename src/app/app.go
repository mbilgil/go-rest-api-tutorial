package app

import (
	"go-rest-api-tutorial/src/example"
	"go-rest-api-tutorial/src/key"
	"go-rest-api-tutorial/src/person"
	"go-rest-api-tutorial/src/university"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the person
	a.Get("/people/{id}", a.handleRequest(person.GetPersonEndpoint))
	a.Post("/people", a.handleRequest(person.CreatePersonEndpoint))
	a.Get("/people", a.handleRequest(person.GetPeopleEndpoint))
	a.Delete("/people/{id}", a.handleRequest(person.DeletePersonEndpoint))

	// Routing for handling university
	a.Get("/university", a.handleRequest(university.GetUniversityByCountry))

	// Routing for handling university
	a.Get("/example", a.handleRequest(example.QueryParamDisplayHandler))

	// Routing for handling jwt
	a.Get("/token", a.handleRequest(key.GetToken))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
