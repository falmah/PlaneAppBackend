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
	a.Post("/pilot/license/{license}/create/img/{oid}", a.handleRequest(handler.WriteImage))
	a.Get("/pilot/license/{license}/get/img/{oid}", a.handleRequest(handler.GetImage))
	a.Post("/pilot/license/{pilot}/create", a.handleRequest(handler.CreateLicense))
	a.Post("/pilot/visa/{visa}/create/img/{oid}", a.handleRequest(handler.WriteVisaImage))
	a.Get("/pilot/visa/{visa}/get/img/{oid}", a.handleRequest(handler.GetVisaImage))
	a.Post("/pilot/visa/{pilot}/create", a.handleRequest(handler.CreateVisa))

	//a.Get("/pilot/license/{pilot}/delete/{license}", a.handleRequest(handler.DeleteLicense))
	a.Get("/pilot/license/{pilot}/all", a.handleRequest(handler.GetLicenses))
	a.Get("/pilot/visa/{pilot}/all", a.handleRequest(handler.GetVisas))
	a.Get("/pilot/request/{pilot}/all", a.handleRequest(handler.GetRequests))
	a.Get("/pilot/request/{pilot}/status/{request}/{status}", a.handleRequest(handler.ChangeRequestStatus))
	//a.Post("/pilot/license/create/img", a.handleRequest(handler.ImageTest))
}
///pilot/license/{operator}/create/img
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
