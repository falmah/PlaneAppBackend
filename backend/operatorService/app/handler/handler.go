package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"driver/operatorDriver"
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
/*
func GetAllProjects(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	projects := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}*/

func GetTickets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator := vars["operator"]

	tic := operatorDriver.GetTickets(operator)
	fmt.Println(tic)

	respondJSON(w, http.StatusOK, tic)
}

func CreatePlane(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Name 				string
		Registration_prefix string
		Registration_id     string
		Plane_type     		string
		Current_location	string	
	}
	vars := mux.Vars(r)
	operator := vars["operator"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var p model.Plane
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &p)
	p.Current_location = t.Current_location
	fmt.Println(p)

	operatorDriver.CreatePlane(operator, &p)

	respondJSON(w, http.StatusCreated, p)
}

func GetPlanes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator := vars["operator"]
	pl := operatorDriver.GetPlanes(operator)
	fmt.Println(pl)
	respondJSON(w, http.StatusOK, pl)
}

func UpdatePlane(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Name 				string
		Registration_prefix string
		Registration_id     string
		Plane_type     		string
		Current_location	string
	}
	vars := mux.Vars(r)
	plane := vars["plane"]
	operator := vars["operator"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var p model.Plane
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &p)
	p.Id = plane
	p.Current_location = t.Current_location
	fmt.Println(p)

	operatorDriver.UpdatePlane(operator, &p)
	respondJSON(w, http.StatusOK, p)
}

func GetPlane(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plane := vars["plane"]
	operator := vars["operator"]

	pl := operatorDriver.GetPlane(operator, plane)
	fmt.Println(pl)

	respondJSON(w, http.StatusOK, pl)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Status 				string
		Pilot_id 			string
		Required_license	string
		Price				uint
		Required_visa       string
		Deadline			string
		Request_comment		string
		Ticket_id			string
		Plane_id            string
	}
	vars := mux.Vars(r)
	operator := vars["operator"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var req model.Request
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &req)

	req.Operator_id = operator
	req.Deadline, _ = time.Parse("2006-01-02", t.Deadline)
	req.Pilot_id = t.Pilot_id
	req.Ticket_id = t.Ticket_id
	req.Plane_id = t.Plane_id
	req.Request_comment = t.Request_comment

	//fmt.Println(req)

	operatorDriver.CreateRequest(operator, &req)

	respondJSON(w, http.StatusCreated, req)
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	request := vars["request"]
	operator := vars["operator"]

	pl := operatorDriver.GetRequest(operator, request)
	fmt.Println(pl)

	respondJSON(w, http.StatusOK, pl)
}

func UpdateRequest(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Status 				string
		Pilot_id 			string
		Required_license	string
		Price				uint
		Required_visa       string
		Deadline			string
		Request_comment		string
		Ticket_id			string
		Plane_id            string
	}
	vars := mux.Vars(r)
	operator := vars["operator"]
	request := vars["request"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var req model.Request
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &req)
	
	req.Id = request
	req.Operator_id = operator
	req.Deadline, _ = time.Parse("2006-01-02", t.Deadline)
	req.Pilot_id = t.Pilot_id
	req.Ticket_id = t.Ticket_id
	req.Plane_id = t.Plane_id
	req.Request_comment = t.Request_comment

	operatorDriver.UpdateRequest(operator, &req)

	respondJSON(w, http.StatusCreated, req)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator := vars["operator"]
	request := vars["request"]
	operatorDriver.DeleteRequest(operator, request)
	respondJSON(w, http.StatusOK, nil)
}

func GetRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator := vars["operator"]

	tic := operatorDriver.GetRequests(operator)
	fmt.Println(tic)

	respondJSON(w, http.StatusOK, tic)
}

func GetPilots(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator := vars["operator"]

	pil := operatorDriver.GePilots(operator)
	fmt.Println(pil)

	respondJSON(w, http.StatusOK, pil)
}
/*
func GetPlanes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	operator := vars["operator"]
	pl := operatorDriver.GetPlanes(operator)
	fmt.Println(pl)
}

func UpdatePlane(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Name 				string
		Registration_prefix string
		Registration_id     string
		Plane_type     		string
		Current_location	string
	}
	vars := mux.Vars(r)
	plane := vars["plane"]
	operator := vars["operator"]

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	var p model.Plane
	json.Unmarshal([]byte(jsonStr), &t)
	json.Unmarshal([]byte(jsonStr), &p)
	p.Id = plane
	p.Current_location = t.Current_location
	fmt.Println(p)

	operatorDriver.UpdatePlane(operator, &p)
	respondJSON(w, http.StatusOK, p)
}

*/
/*
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
	c.Dest_from = t.Dest_from
	c.Dest_to = t.Dest_to
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
	c.Dest_from = t.Dest_from
	c.Dest_to = t.Dest_to
	c.Price = t.Price
	c.Ticket_comment = t.Ticket_comment

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