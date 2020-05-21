package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"grafana-json/pkg"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", checkHealth).Methods(http.MethodGet)
	r.HandleFunc("/search", search).Methods(http.MethodPost)
	r.HandleFunc("/query", query).Methods(http.MethodPost)

	go pkg.InitRemainingAMT()
	if err:=http.ListenAndServe(":9999", r); err!= nil{
		log.Fatalln(err)
	}
}

func checkHealth(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func search(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(pkg.InitSearch())
}

func query(w http.ResponseWriter, r * http.Request){
	result := pkg.Query()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&result)
}