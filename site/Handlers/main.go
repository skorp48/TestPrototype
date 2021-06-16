package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var Link *string
var ResThem []Result

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
	Question     string     `json:"Question"`
	QuestionType int        `json:"QuestionType"`
	Image        string     `json:"Image"`
	Answers      []string   `json:"Answers"`
	Weights      [][]string `json:"Weights"`
}

type QuestionsList struct {
	Name        string     `json:"Name"`
	Description string     `json:"Description"`
	Image       string     `json:"Image"`
	Object      []string   `json:"object"`
	Questions   []Question `json:"Questions"`
}

type ResultThemes struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type Result struct {
	Type   string         `json:"type"`
	Name   string         `json:"Name"`
	Themes []ResultThemes `json:"themes"`
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
		return nil, errors.New("not found")
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

	var Questions QuestionsList
	decodeJSON("site/Handlers/Config/Questions/IVTQuestions.json", &Questions)

	ObjectMap := make(map[string]int)
	for _, v := range Questions.Object {

		ObjectMap[v] = 0
	}
	fmt.Printf("ObjectMap: %v\n", ObjectMap)

	var AnswerID []string
	var Answer []string

	r.ParseForm()
	fmt.Printf("r.Form: %v\n", r.Form)
	for i, v := range r.Form {
		fmt.Printf("v: %v\n", v)
		fmt.Printf("i: %v\n", i)
		words := strings.Split(i, "_")
		if len(words) == 2 {
			AnswerID = append(AnswerID, words[0])
			Answer = append(Answer, words[1])
		} else {
			AnswerID = append(AnswerID, words[0])
			Answer = append(Answer, v[0])
		}
	}

	fmt.Printf("AnswerID: %v\n", AnswerID)
	fmt.Printf("Answer: %v\n", Answer)

	for i, v := range AnswerID {
		int_v, _ := strconv.Atoi(v)
		int_answer, _ := strconv.Atoi(Answer[i])
		for _, val := range Questions.Questions[int_v].Weights[int_answer] {
			ObjectMap[val] += 1
		}
	}
	fmt.Printf("ObjectMap: %v\n", ObjectMap)

	max := 0
	for Key, v := range ObjectMap {
		if v > max && Key != "none" {
			max = v
		}
	}

	var themes []string
	for Key, v := range ObjectMap {
		if v == max {
			themes = append(themes, Key)
		}
	}

	fmt.Println(themes)

	var ResThemJS []Result
	decodeJSON("site/Handlers/Config/MainConfig/IVTThemes.json", &ResThemJS)
	for _, v := range ResThemJS {
		for _, val := range themes {
			if v.Type == val {
				ResThem = append(ResThem, v)
			}
		}
	}

	fmt.Println(ResThem)

	temp, err := template.ParseFiles("site/templates/html/TestResult.html")
	if err != nil {
		fmt.Println(err)
	}
	temp.Execute(w, nil)
}

func fetchResult(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(ResThem)
	ResThem = nil
}

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./site/templates/"))))
	http.HandleFunc("/", Handler)
	http.HandleFunc("/fetchobj", SendObj)
	http.HandleFunc("/fetchQuestions", fetchQuestions)
	http.HandleFunc("/Questions", TestWorkHand)
	http.HandleFunc("/testcalc", testCalc)
	http.HandleFunc("/fetchres", fetchResult)
	http.ListenAndServe(":5000", nil)
}
