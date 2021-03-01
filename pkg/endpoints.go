package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AddIn struct {
	A int `json:"a"`
	B int `json:"b"`
}

type DaysIn struct {
	Days int `json:"days"`
}

// Add endpoint returns the addition of 2 input numbers
func Add(w http.ResponseWriter, r *http.Request) {
	in := &AddIn{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(strconv.Itoa(in.A + in.B)))
	if err != nil {
		log.Fatal("Error writing or returning the response", err)
	}
}

// Hello endpoint returns the Hello string
func Hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello!!"))
	if err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}

// Time endpoint returns the Hello string
func Time(w http.ResponseWriter, r *http.Request) {
	in := &DaysIn{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(time.Now().AddDate(0, 0, in.Days).String()))
	if err != nil {
		log.Fatal("Error writing or returning the response", err)
	}
}
