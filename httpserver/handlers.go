package main

import (
	"encoding/json"
	"fmt"
	"github.com/fileserver/httpserver/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func ApiHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Print("API Handler")

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		body, err := json.Marshal(map[string]interface{}{
			"GET /api/":            "Info",
			"GET /api/todo":        "Get all todos",
			"GET /api/todo?text=":  "Get todo by text",
			"POST /api/todo":       "Add todo",
			"POST /api/todo/:id":   "Update todo",
			"DELETE /api/todo/:id": "Delete todo",
		})

		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(body)
	}
}

func TodoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Print("API Handler")

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch req.Method {
		case "GET":
			text, ok := req.URL.Query()["text"]
			if ok {
				for _, s := range model.Todos {
					if strings.Contains(s.Text, text[0]) {
						ReturnJson(s, w, 200)
						return
					}

					w.WriteHeader(404)
				}
			} else {
				ReturnJson(model.Todos, w, 200)
			}
		case "POST":

			var newTodo model.TodoItem
			err := json.NewDecoder(req.Body).Decode(&newTodo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			//log.Printf("Post from website! r.PostFrom = %v\n", newTodo)

			newTodo.ID = GetId()

			model.Todos = append(model.Todos, newTodo)
			ReturnJson(newTodo, w, 200)

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

	}
}

func TodoIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := strings.TrimPrefix(req.URL.Path, "/api/todo/")
		log.Print("API Id Handler " + id)
		intId, err := strconv.ParseInt(id, 0, 8)

		if err != nil {
			ReturnJson("Error with ID", w, 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch req.Method {
		case "GET":

			for _, s := range model.Todos {
				if int64(s.ID) == intId {
					ReturnJson(s, w, 200)
					return
				}
			}

			w.WriteHeader(404)

		case "PUT":

			var newTodo model.TodoItem
			err := json.NewDecoder(req.Body).Decode(&newTodo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			found := false
			for i, v := range model.Todos {
				if v.ID == newTodo.ID {
					found = true
					model.Todos[i] = newTodo
				}
			}
			if found == false {
				ReturnJson("Error, ID not found", w, 400)
				return
			} else {
				ReturnJson(newTodo, w, 200)
			}
		case "DELETE":

			for index, s := range model.Todos {
				if int64(s.ID) == intId {
					model.Todos = append(model.Todos[:index], model.Todos[index+1:]...)
					ReturnJson(s, w, 200)
					return
				}
			}

			w.WriteHeader(404)

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}

	}
}

func StaticHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("Static Handler")

		path, _ := filepath.Abs("./index.html")
		b, err := ioutil.ReadFile(path)
		if err != nil {
			http.Error(w, r.RequestURI, http.StatusExpectationFailed)
			return
		}

		html := string(b)
		w.Header().Set("content-type", "text/html")
		_, err = io.WriteString(w, html)
		if err != nil {
			log.Print("HTTP I/O error [%v]", err.Error())
		}
	}
}

// Marshall data and return
// Return 500 if error during Marshaling
func ReturnJson(data any, w http.ResponseWriter, code int) {
	body, err := json.Marshal(data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(body)
}

func GetId() int {
	var max int = 0
	for _, v := range model.Todos {
		if v.ID > max {
			max = v.ID
		}
	}
	max++
	return max
}
