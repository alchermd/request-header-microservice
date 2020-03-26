package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Header struct {
	IpAddress string `json:"ipaddress"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Serving " + r.URL.Path)

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, nil)
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Serving " + r.URL.Path)

	h := &Header{
		IpAddress: r.RemoteAddr,
	}

	j, err := json.Marshal(h)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(j))
}

func main() {
	const API_URL = "/api"
	const HEADER_URL = "/whoami"

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT is not set")
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc(API_URL+HEADER_URL, headerHandler)

	log.Print("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
