package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter,r *http.Request) {
	repsond(w, "hello api")
} 

func respond(w http.ResponseWriter,data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("ERR: could not marshal data", err.Error())
	}

	bytesCount, err := w.Write(bytes)
	if err != nil {
		log.Println("ERR: could not write")
	}
	log.Println("number of bytes", bytesCount)

}