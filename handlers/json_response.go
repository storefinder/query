package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	contentType       = "application/json; charset=UTF-8"
	contentTypeHeader = "Content-Type"
)

//JSONWriter interface
type JSONWriter interface {
	Write(w http.ResponseWriter)
}

//JSONResponse defines a JSON response sent to client
type JSONResponse struct {
	status int
	data   interface{}
}

func (jr *JSONResponse) Write(w http.ResponseWriter) {
	log.Printf("HTTP Status : %v", jr.status)
	//write to response
	w.Header().Set(contentTypeHeader, contentType)
	w.WriteHeader(jr.status)
	if err := json.NewEncoder(w).Encode(jr.data); err != nil {
		log.Printf("Error Writing JSON to response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
