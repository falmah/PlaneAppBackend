package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/hex"
	"driver/operatorDriver"
	"driver/pilotDriver"
	"model"
	"fmt"
	"os"
	"io"
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
	pilotDriver.WriteImage(oid, data)

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
