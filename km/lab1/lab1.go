package main

import (
	"encoding/json"
	"github.com/pityara/labs/km/lab1/lib"
	"github.com/pityara/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ownvector", ownVectorHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error while trying to start http server");
	}
}

func ownVectorHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var in lib.InStruct
	err := decoder.Decode(&in)
	if err != nil {
		panic(err)
	}

	outJson := lib.GetOwnVector(in);

	if err != nil {
		log.Fatal("Error occurs while trying make ANR");
	}

	utils.HttpJSONRequest(w, outJson)
}