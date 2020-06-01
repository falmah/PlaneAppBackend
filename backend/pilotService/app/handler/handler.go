package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/hex"
	"driver/pilotDriver"
	"model"
	"fmt"
	"os"
	"io"
	"bytes"
	"strconv"
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

func ImageTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("there")
	var buf bytes.Buffer
	r.ParseMultipartForm(32 << 20) // limit your max input length!
	
	file, _, err := r.FormFile("fileupload")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	oid := pilotDriver.CreateOid()

	fmt.Println("oid created", oid)
	io.Copy(&buf, file)

	size := buf.Len()
	data := hex.EncodeToString(buf.Next(size))
	pilotDriver.WriteImage(int(oid), data)

	new_buf := pilotDriver.ReadImage(int(oid), int(size))
	fmt.Println(new_buf)
	// copy example
	f, err := os.OpenFile("./testfile.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, bytes.NewBuffer(new_buf))
}

func GetLicenses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pilot := vars["pilot"]

	req := pilotDriver.GetLicenses(pilot)
	respondJSON(w, http.StatusOK, req)
}

func CreateLicense(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pilot := vars["pilot"]

	bufs := new(bytes.Buffer)
    bufs.ReadFrom(r.Body)
	jsonStr := bufs.String()
	fmt.Println(jsonStr)

	var req model.License
	json.Unmarshal([]byte(jsonStr), &req)
	req.Pilot_id = pilot
	req.Is_active = true
	fmt.Println(req)

	pilotDriver.CreateLicense(pilot, &req)

	respondJSON(w, http.StatusCreated, req)
}

func WriteImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	license := vars["license"]
	oid, _ := strconv.Atoi(vars["oid"])

	var buf bytes.Buffer
	r.ParseMultipartForm(32 << 20) // limit your max input length!

	file, _, err := r.FormFile("fileupload")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	io.Copy(&buf, file)

	size := buf.Len()
	data := hex.EncodeToString(buf.Next(size))

	pilotDriver.WriteImage(oid, data)

	pilotDriver.UpdateImageSize(license, size)

	new_buf := pilotDriver.ReadImage(oid, size)
	fmt.Println(new_buf)
	// copy example
	f, err := os.OpenFile("./testfile.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, bytes.NewBuffer(new_buf))

}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	license := vars["license"]
	//oid, _ := strconv.Atoi(vars["oid"])

	lis := pilotDriver.GetLicense(license)

	new_buf := pilotDriver.ReadImage(int(lis.Image), int(lis.Image_size))
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, bytes.NewBuffer(new_buf))

}

func GetRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pilot := vars["pilot"]

	req := pilotDriver.GetRequests(pilot)
	respondJSON(w, http.StatusOK, req)
}

func ChangeRequestStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pilot := vars["pilot"]
	status := vars["status"]
	request := vars["request"]

	pilotDriver.ChangeRequestStatus(pilot, status, request)

	respondJSON(w, http.StatusOK, nil)
}
