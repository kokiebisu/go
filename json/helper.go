package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func unmarshal(w http.ResponseWriter, req *http.Request) {
	var u user
	err := json.Unmarshal([]byte(jsonuser), &u)
	if err != nil {
		http.Error(w, "Cannot unmarshal", http.StatusInternalServerError)
	}
	fmt.Println("unmarshaled user", u)
	fmt.Println("unmarshalled lastname", u.LastName)
}

func decode(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder()
}