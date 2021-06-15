package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Тело запроса с сервера
type requsetBody struct {
	Action    string `json:"action"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Id        string `json:"id"`
	HiddenId  string `json:"hiddenId"`
}

// Структура для обработки JSON
type QConfig struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Image       string `json:"Image"`
	Link        string `json:"Link"`
}

// Возварщает главную страницу
func Handler(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("site/templates/html/main.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, nil)
}

// Возварщает страницу с вопросами
func TestWorkHand(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("site/templates/html/TestWork.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, nil)
}

// Получает списко направлений из  JSON
func SendObj(w http.ResponseWriter, r *http.Request) {
	rBody := requsetBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rBody)
	if err != nil {
		fmt.Println(err)
	}

	in, err := os.Open("site/Handlers/Config/MainConfig/QConfig.json")
	if err != nil {
		fmt.Println(err)
	}
	defer in.Close()
	decodeJson := json.NewDecoder(in)
	var obj []QConfig
	err = decodeJson.Decode(&obj)
	if err != nil {
		fmt.Println(err)
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./site/templates/"))))
	http.HandleFunc("/", Handler)
	http.HandleFunc("/fetchobj", SendObj)
	http.HandleFunc("/TestWork", TestWorkHand)
	http.ListenAndServe(":5000", nil)
}
