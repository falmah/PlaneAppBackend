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
	a.Get("/operator/pilot/{operator}/all", a.handleRequest(handler.GetPilots))
	a.Get("/operator/ticket/{operator}/all", a.handleRequest(handler.GetTickets))
	a.Get("/operator/plane/{operator}/all", a.handleRequest(handler.GetPlanes))
	a.Get("/operator/plane/{operator}/get/{plane}", a.handleRequest(handler.GetPlane))
	a.Get("/operator/request/{operator}/get/{request}", a.handleRequest(handler.GetRequest))
	a.Get("/operator/request/{operator}/delete/{request}", a.handleRequest(handler.DeleteRequest))
	a.Get("/operator/request/{operator}/all", a.handleRequest(handler.GetRequests))
	a.Post("/operator/plane/{operator}/create", a.handleRequest(handler.CreatePlane))
	a.Post("/operator/plane/{operator}/update/{plane}", a.handleRequest(handler.UpdatePlane))
	a.Post("/operator/request/{operator}/create",  a.handleRequest(handler.CreateRequest))
	a.Post("/operator/request/{operator}/update/{request}", a.handleRequest(handler.UpdateRequest))
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}


/*
// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}
*/
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
