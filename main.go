package main

import (
	"app/rps"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func playRound(rw http.ResponseWriter, r *http.Request) {
	playerValue, _ := strconv.Atoi(r.URL.Query().Get("playerValue"))
	result := rps.PlayRound(playerValue)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
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
