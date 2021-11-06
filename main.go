package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var todos []string
var file *os.File

func GetToDos(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		json, err := json.Marshal(&todos)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(json)
	}
}
func PostToDo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if (*r).Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("cannot read the request body %v", err)
			return
		}
		req_j := &struct {
			Todo string `json:"todo"`
		}{}
		err = json.Unmarshal(body, req_j)
		if err != nil {
			log.Printf("error parsing request %v", err)
			return
		}
		str := ""
		if len(todos) == 0 {
			str = fmt.Sprintf("\"%v\"", (*req_j).Todo)
		} else {
			str = fmt.Sprintf(",\"%v\"", (*req_j).Todo)
		}
		todos = append(todos, req_j.Todo)
		_, err_wr := file.WriteString(str)
		if err_wr != nil && file != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("error writing to file %v %v", err_wr, file)
			return
		}
		rw.WriteHeader(201)
	}
}

func main() {
	var err error
	file, err = os.OpenFile("todos.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	file_content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read from file %v", err)
		return
	}
	if string(file_content) == "" {
		file_content = append(file_content, '[')
		_, err = file.WriteString("[")
		if err != nil {
			log.Fatalf("cannot write to file %v", err)
		}
	}
	file_content = append(file_content, ']')
	todos = make([]string, 5)
	err = json.Unmarshal(file_content, &todos)
	if err != nil {
		log.Fatalf("unable to read file2 %v", err)
		return
	}
	http.HandleFunc("/api/todos", GetToDos)
	http.HandleFunc("/api/todo", PostToDo)
	log.Fatalf("%v", http.ListenAndServe(":4000", nil))
}
