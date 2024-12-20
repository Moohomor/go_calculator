package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Calc_svc(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var exStruct struct {
		Expression string `json:"expression"`
	}
	err := json.NewDecoder(r.Body).Decode(&exStruct)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// expression := exStruct.Expression
	res, err := Calc(exStruct.Expression)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := struct {
		Result string `json:"result"`
	}{Result: strconv.FormatFloat(res, 'f', -1, 64)}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/v1/calculate", Calc_svc)
	log.Println("serving")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
