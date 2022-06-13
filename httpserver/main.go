package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var dirName string = "/"
var servercert string = os.Getenv("CERTIFICATE_FILE")
var keycert string = os.Getenv("CERTIFICATE_KEY")
var env *string = flag.String("env", "prod", "dev will serve in http")

func main() {

	Router()

	log.Print(" Certificate file: " + servercert)
	log.Print(" Certificate file: " + keycert)

	flag.Parse()

	if *env == "dev" {
		err := http.ListenAndServe(":6080", nil)
		log.Print("Starting server in 6080 - env " + *env)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
	} else {
		err := http.ListenAndServeTLS(":8443", servercert, keycert, nil)
		log.Print("Starting server in 8443 - env " + *env)
		if err != nil {
			fmt.Print(err.Error())
			return
		}
	}

}
