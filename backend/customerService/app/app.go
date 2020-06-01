package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"app/handler"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	log.Info("Init HTTP route")
	a.Router = mux.NewRouter()
	a.setRouters()
	log.Info("HTTP route succsesfult initialized")
}

func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/customer/ticket/{customer}/get/{ticket}", a.handleRequest(handler.GetTicket))
	a.Get("/customer/ticket/{customer}/delete/{ticket}", a.handleRequest(handler.DeleteTicket))
	a.Get("/customer/ticket/{customer}/all", a.handleRequest(handler.GetTickets))
	a.Post("/customer/ticket/{customer}/create",a.handleRequest(handler.CreateTicket))
	a.Post("/customer/ticket/{customer}/update/{ticket}",a.handleRequest(handler.UpdateTicket))
	a.Get("/customer/proposal/{customer}/all", a.handleRequest(handler.GetProposals))
	a.Get("/customer/proposal/{customer}/status/{proposal}/{status}", a.handleRequest(handler.ChangeProposalStatus))
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
