package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var Link *string

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

type Question struct {
	Question     string   `json:"Question"`
	QuestionType int      `json:"QuestionType"`
	Image        string   `json:"Image"`
	Answers      []string `json:"Answers"`
	Weights      []int    `json:"Weights"`
}

type QuestionsList struct {
	Name        string     `json:"Name"`
	Description string     `json:"Description"`
	Image       string     `json:"Image"`
	Questions   []Question `json:"Questions"`
}

func decodeJSON(filename string, v interface{}) {
	in, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer in.Close()
	decodeJson := json.NewDecoder(in)
	err = decodeJson.Decode(&v)
	if err != nil {

		fmt.Println(err)
	}
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) (*string, error) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.NotFound(w, r)
		fmt.Println("123")
		return nil, errors.New("Not Found")
	}
	return &q, nil
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
	Link, err = showSnippet(w, r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*Link)
	temp.Execute(w, nil)
}

func fetchQuestions(w http.ResponseWriter, r *http.Request) {
	var obj QuestionsList

	decodeJSON("site/Handlers/Config/Questions/"+*Link+".json", &obj)
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}

// Получает списко направлений из  JSON
func SendObj(w http.ResponseWriter, r *http.Request) {
	rBody := requsetBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rBody)
	if err != nil {
		fmt.Println(err)
	}
	var obj []QConfig
	decodeJSON("site/Handlers/Config/MainConfig/QConfig.json", &obj)
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}

func testCalc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	temp, err := template.ParseFiles("site/templates/html/TestResult.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, nil)
}

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./site/templates/"))))
	http.HandleFunc("/", Handler)
	http.HandleFunc("/fetchobj", SendObj)
	http.HandleFunc("/fetchQuestions", fetchQuestions)
	http.HandleFunc("/Questions", TestWorkHand)
	http.HandleFunc("/testcalc", testCalc)
	http.ListenAndServe(":5000", nil)
}
