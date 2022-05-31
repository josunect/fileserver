package filemanager

import (
	"io/ioutil"
	"log"
)

func GetDirFiles(directory string) ([]string, error) {
	var files, err = ioutil.ReadDir(directory)
	fileList := make([]string, len(files))

	if err != nil {
		log.Print(err)
	}

	for index := 0; index < len(files); index++ {
		fileList[index] = files[index].Name()
	}

	return fileList, err
}
