package main

import (
	"encoding/json"
	"github.com/pityara/labs/km/lab1/ownvector/matrix"
	"github.com/pityara/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error while trying to start http server");
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var in matrix.Matrix
	err := decoder.Decode(&in)
	if err != nil {
		panic(err)
	}

	outJson := matrix.GetOwnRelativeVector(in)
	if err != nil {
		log.Fatal("Error occurs while trying make own matrix vector");
	}

	utils.HttpJSONRequest(w, outJson)
}
