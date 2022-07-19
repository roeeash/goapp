package main

import (
	"log"
	"net/http"
	"os/user"
	"text/template"
	"time"
)

type GreetMessage struct {
	Name string
	Time string
}

func main() {
	//vars
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	userstring := user.Username

	message := GreetMessage{userstring, time.Now().Format(time.Stamp)}
	template := template.Must(
		template.ParseFiles("html/homepage.html"))

	//http handle settings

	http.Handle("/html/", http.StripPrefix("/html/",
		http.FileServer(http.Dir("html"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if username := r.FormValue(userstring); username != "" {
			message.Name = username
		}

		if err := template.ExecuteTemplate(w, "homepage.html", message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	http.ListenAndServe(":8080", nil)
}
