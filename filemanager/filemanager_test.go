package filemanager

import (
	"testing"
)

func TestGetDirFiles(t *testing.T) {
	dir := "/tmp"
	files, err := GetDirFiles(dir)

	if len(files) == 0 || err != nil {
		t.Fatal("Error reading files")
	}
}

func TestGetDirFilesFail(t *testing.T) {
	dir := "/blablabla"
	files, err := GetDirFiles(dir)

	if len(files) > 0 || err == nil {
		t.Fatal("Error reading files")
	}
}
