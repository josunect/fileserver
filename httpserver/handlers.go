package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func ApiHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Print("API Handler")

		res.Header().Set("Content-Type", "application/json")

		body, err := json.Marshal(map[string]interface{}{
			"data": "Hello, world",
		})

		if err != nil {
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(200)
		res.Write(body)
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
