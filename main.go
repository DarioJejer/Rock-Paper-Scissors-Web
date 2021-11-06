package main

import (
	"app/rps"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func playRound(w http.ResponseWriter, r *http.Request) {
	winner, computerChoise, roundResult := rps.PlayRound(1)
	fmt.Println(winner, computerChoise, roundResult)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
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
	http.HandleFunc("/playRound", playRound)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
