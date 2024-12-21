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
		// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusInternalServerError)
		response := struct {
			Error string `json:"error"`
		}{Error: "Internal server error"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var exStruct struct {
		Expression string `json:"expression"`
	}
	err := json.NewDecoder(r.Body).Decode(&exStruct)
	if err != nil {
		raiseError(w, err)
		return
	}
	// expression := exStruct.Expression
	res, err := Calc(exStruct.Expression)
	if err != nil {
		raiseError(w, err)
		return
	}
	response := struct {
		Result string `json:"result"`
	}{Result: strconv.FormatFloat(res, 'f', -1, 64)}
	json.NewEncoder(w).Encode(response)
}

func raiseError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusUnprocessableEntity)
	response := struct {
		Error string `json:"error"`
	}{Error: "Expression is not valid"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/v1/calculate", Calc_svc)
	log.Println("serving on localhost:80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
