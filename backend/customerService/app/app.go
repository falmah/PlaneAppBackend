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
	//a.Post("/register", a.handleRequest(handler.RegisterUser))
	//a.Post("/login", a.handleRequest(handler.GetUser))
	/*a.Post("/projects", a.handleRequest(handler.CreateProject))
	a.Get("/projects/{title}", a.handleRequest(handler.GetProject))
	a.Put("/projects/{title}", a.handleRequest(handler.UpdateProject))
	a.Delete("/projects/{title}", a.handleRequest(handler.DeleteProject))
	a.Put("/projects/{title}/archive", a.handleRequest(handler.ArchiveProject))
	a.Delete("/projects/{title}/archive", a.handleRequest(handler.RestoreProject))

	// Routing for handling the tasks
	a.Get("/projects/{title}/tasks", a.handleRequest(handler.GetAllTasks))
	a.Post("/projects/{title}/tasks", a.handleRequest(handler.CreateTask))
	a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
	a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
	a.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.UndoTask))*/
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
