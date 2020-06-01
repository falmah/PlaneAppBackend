package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"driver/customerDriver"
	"model"
	"fmt"
	"bytes"
	"time"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ticket := vars["ticket"]
	customer := vars["customer"]

	tic := customerDriver.GetTicket(customer, ticket)
	fmt.Println(tic)

	respondJSON(w, http.StatusOK, tic)
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer := vars["customer"]

	tic := customerDriver.GetTickets(customer)
	fmt.Println(tic)

	respondJSON(w, http.StatusOK, tic)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Status 			string
		Cargo_type		string
		Title			string
		Date_from		string
		Date_to			string
		Dest_from		string
		Dest_to			string
		Price			uint
		Ticket_comment	string
	}

	vars := mux.Vars(r)
	customer := vars["customer"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var c model.Ticket
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &c)

	c.Customer_id = customer
	c.Date_from, _ = time.Parse("2006-01-02", t.Date_from)
	c.Date_to, _ = time.Parse("2006-01-02", t.Date_to)
	c.Dest_from = customerDriver.GetAirportId(t.Dest_from)
	c.Dest_to = customerDriver.GetAirportId(t.Dest_to)
	c.Price = t.Price
	c.Ticket_comment = t.Ticket_comment 

	customerDriver.CreateTicket(&c, customer)
	respondJSON(w, http.StatusCreated, c)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Status 			string
		Cargo_type		string
		Title			string
		Date_from		string
		Date_to			string
		Dest_from		string
		Dest_to			string
		Price			uint
		Ticket_comment	string
	}

	vars := mux.Vars(r)
	customer := vars["customer"]
	ticket := vars["ticket"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var c model.Ticket
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &c)

	c.Id = ticket
	c.Customer_id = customer
	c.Date_from, _ = time.Parse("2006-01-02", t.Date_from)
	c.Date_to, _ = time.Parse("2006-01-02", t.Date_to)
	c.Dest_from = customerDriver.GetAirportId(t.Dest_from)
	c.Dest_to = customerDriver.GetAirportId(t.Dest_to)
	c.Price = t.Price
	c.Ticket_comment = t.Ticket_comment

	fmt.Println(c);

	customerDriver.UpdateTicket(customer, &c)
	respondJSON(w, http.StatusOK, c)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer := vars["customer"]
	ticket := vars["ticket"]
	customerDriver.DeleteTicket(customer, ticket)
	respondJSON(w, http.StatusOK, nil)
}

func GetProposals(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer := vars["customer"]

	fmt.Println(customer)
	tic := customerDriver.GetProposals(customer)
	fmt.Println(tic)

	respondJSON(w, http.StatusOK, tic)
}

func ChangeProposalStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer := vars["customer"]
	status := vars["status"]
	proposal := vars["proposal"]

	customerDriver.ChangeProposalStatus(customer, status, proposal)

	respondJSON(w, http.StatusOK, nil)
}
/*
func CreateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	project := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func GetProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func UpdateProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func DeleteProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	if err := db.Delete(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func ArchiveProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	project.Archive()
	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func RestoreProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getProjectOr404(db, title, w, r)
	if project == nil {
		return
	}
	project.Restore()
	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// getProjectOr404 gets a project instance if exists, or respond the 404 error otherwise
func getProjectOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Project {
	project := model.Project{}
	if err := db.First(&project, model.Project{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &project
}
*/