package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("INFO: starting app")

	router := http.NewServeMux()
	router.HandleFunc("/", Home)

	// router.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
	// 	bytesCount, err := w.Write([]byte("hello api"))
	// 	if err != nil {
	// 		log.Println("ERR: could not write")
	// 	}
	// 	log.Println("number of bytes", bytesCount)
	// })

	http.ListenAndServe(":8080", router)

}

func Home(w http.ResponseWriter,r *http.Request) {
	respond(w, "hello from Home")
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