package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var dirName string = "/"
var servercert string = os.Getenv("CERTIFICATE_FILE")
var keycert string = os.Getenv("CERTIFICATE_KEY")

func main() {

	Router()

	log.Print("Certificate file: " + servercert)
	log.Print("Certificate file: " + keycert)

	//err := http.ListenAndServe(":6080", nil)
	err := http.ListenAndServeTLS(":8443", servercert, keycert, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}
