package main

import (
	"fmt"
	"github.com/fileserver/filemanager"
	"log"
	"net/http"
	"os"
	"strings"
)

var dirName string = "/tmp"
var crtFile string = "localhost.crt"
var keyFile string = "localhost.key"

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(writer, "<h2> Welcome to this file server! </h2> ")
	fmt.Fprintf(writer, "Where to go from here? ")
	fmt.Fprintf(writer, "<ul><li> <a href='/list'>List files </a></li>")
	fmt.Fprintf(writer, "<li> <a href='/directory'>Change directory </a></li>")
	fmt.Fprintf(writer, "</ul>")
}

func directory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(writer, "<h2> Change directory: </h2> ")
	fmt.Fprintf(writer, "<form action='/list' method='get'> ")
	fmt.Fprintf(writer, "<input type='text' name='dir' placeholder='/tmp'>")
	fmt.Fprintf(writer, "<input type='submit' value='Go'>")
	fmt.Fprintf(writer, "</form>")
	fmt.Fprintf(writer, "<a href='/list'>Directory list</a>")
}

// List files from a dir
func listFiles(writer http.ResponseWriter, request *http.Request) {
	// Get parameters
	name, ok := request.URL.Query()["dir"]
	var oldName string
	if !ok {
		log.Print("No parameters")
	}
	oldName = dirName
	if ok && len(name[0]) > 0 {

		dirName = name[0]
	}

	// Send response
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(writer, "<h2> File List ("+dirName+"): </h2> ")
	var fileList, error = filemanager.GetDirFiles(dirName)
	if error != nil {
		if strings.Contains(error.Error(), "not a directory") {

			dat, err := os.ReadFile(dirName)
			if err != nil {
				fmt.Fprintf(writer, "Error reading file.")
			} else {
				fmt.Fprintf(writer, "<pre> "+string(dat)+"</pre>")
			}

			fmt.Fprintf(writer, "<a href='/list'>Go Back </a>")
			dirName = oldName

		} else {
			fmt.Fprintf(writer, "Error getting files from dir. It doesn't look like a correct one?<br>")
			fmt.Fprintf(writer, "<a href='/directory'>Change directory </a>")
		}

	} else {
		fmt.Fprintf(writer, "<ul> ")
		for _, element := range fileList {
			fmt.Fprintf(writer, "<li> <a href='/list?dir="+dirName+"/"+element+"'>"+element+"</a></li>")
		}
		fmt.Fprintf(writer, "</ul> ")
		fmt.Fprintf(writer, "<a href='/list?dir="+oldName+"'>Go Back </a> <br>")
		fmt.Fprintf(writer, "<a href='/directory'>Change directory </a>")
	}

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/list", listFiles)
	http.HandleFunc("/directory", directory)

	err := http.ListenAndServeTLS(":8443", crtFile, keyFile, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}
