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

			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}

			name := req.FormValue("name")
			occupation := req.FormValue("occupation")

			fmt.Fprintf(w, "%s is a %s\n", name, occupation)

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
