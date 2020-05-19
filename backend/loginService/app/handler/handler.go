package handler

import (
	"encoding/json"
	"net/http"
	_ "github.com/gorilla/mux"
	"driver/loginDriver"
	"model"
	"fmt"
	"bytes"
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Email string
		Password string
	}

	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	json.Unmarshal([]byte(jsonStr), &t)
	fmt.Println(t)
	us := loginDriver.GetUser(t.Email, t.Password)
	fmt.Println(us)

	switch us.Role {
	case "customer":
		getCustomer(w, us)
	case "operator":
		getOperator(w, us)
	case "pilot":
		getPilot(w, us)
	default:
		respondError(w, http.StatusBadRequest, "role not detected")
	}

}

func getCustomer(w http.ResponseWriter, u model.User) {
	customer := loginDriver.GetCustomer(u)
	fmt.Println(customer)
	respondJSON(w, http.StatusOK, customer)
}

func getOperator(w http.ResponseWriter, u model.User) {
	operator := loginDriver.GetOperator(u)
	fmt.Println(operator)
	respondJSON(w, http.StatusOK, operator)
}

func getPilot(w http.ResponseWriter, u model.User) {
	pilot := loginDriver.GetPilot(u)
	fmt.Println(pilot)
	respondJSON(w, http.StatusOK, pilot)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	type Tmp struct {
		Id			string
		Name 		string
		Surname 	string
		Phone 		string
		Email 		string
		Role		string
		Password 	string
	}
	buf := new(bytes.Buffer)
    buf.ReadFrom(r.Body)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	var t Tmp
	json.Unmarshal([]byte(jsonStr), &t)
	fmt.Println(t)

	us := model.User{}
	us.Name = t.Name
	us.Surname = t.Surname
	us.Phone = t.Phone
	us.Email = t.Email
	us.Role = t.Role
	us.Password = t.Password
	fmt.Println(us)

	switch us.Role {
	case "customer":
		registerCustomer(w, us)
	case "operator":
		type data struct {
			Company_name string
			City string
		}
		var d data
		json.Unmarshal([]byte(jsonStr), &d)
		fmt.Println(d)
		registerOperator(w, us, d.City, d.Company_name)
	case "pilot":
		type data struct {
			City string
		}
		var d data
		json.Unmarshal([]byte(jsonStr), &d)
		fmt.Println(d)
		registerPilot(w, us, d.City)
	default:
		respondError(w, http.StatusBadRequest, "role not detected")
	}

}

func registerCustomer(w http.ResponseWriter, u model.User) {
	customer := model.Customer{}
	customer.User = u
	loginDriver.RegisterCustomer(&customer)
	respondJSON(w, http.StatusCreated, customer)
}

func registerOperator(w http.ResponseWriter, u model.User, 
						city_name string, company_name string) {
	operator := model.Operator{}
	operator.Company_name = company_name
	operator.User = u
	operator.City = loginDriver.GetCity(city_name)
	operator.City_id = operator.City.Id
	loginDriver.RegisterOperator(&operator)
	respondJSON(w, http.StatusCreated, operator)
}

func registerPilot(w http.ResponseWriter, u model.User, city_name string) {
	pilot := model.Pilot{}
	pilot.Busy = false
	pilot.User = u
	pilot.City = loginDriver.GetCity(city_name)
	pilot.Current_location = pilot.City.Id
	loginDriver.RegisterPilot(&pilot)
	respondJSON(w, http.StatusCreated, pilot)
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