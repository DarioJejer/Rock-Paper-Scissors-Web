package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
