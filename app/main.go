package main

import (
	"html/template"
	"log"
	"net/http"
	"os/user"
	"time"
)

//struct GreetMessage
type GreetMessage struct {
	Name string
	Time string
}

func main() {
	//vars
	user, err := user.Current()

	//error handling
	if err != nil {
		log.Fatalf(err.Error())
	}

	userstring := user.Username

	//generate message

	message := GreetMessage{userstring, time.Now().Format(time.Stamp)}
	template := template.Must(template.ParseFiles("template/template.html"))

	//strip the "/static/" prefix

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//handle function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if username := r.FormValue(userstring); username != "" {
			message.Name = username
		}
		if err := template.ExecuteTemplate(w, "template.html", message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	//listen and serve in port 8080 in http protocol
	http.ListenAndServe(":8080", nil)
}
