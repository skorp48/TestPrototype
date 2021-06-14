package main

import (
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("site/templates/html/main.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func TestWorkHand(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("site/templates/html/TestWork.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./site/templates/"))))
	http.HandleFunc("/", Handler)
	http.HandleFunc("/TestWork", TestWorkHand)
	http.ListenAndServe(":5000", nil)
}
