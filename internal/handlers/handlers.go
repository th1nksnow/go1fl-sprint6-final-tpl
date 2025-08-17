package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func HandleRoot(res http.ResponseWriter, req *http.Request) {

	path := "./index.html"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(res, "index.html is not found", http.StatusInternalServerError)
		return
	}
	http.ServeFile(res, req, path)
}

func HandleUpload(res http.ResponseWriter, req *http.Request) {

	file, _, err := req.FormFile("myFile")
	if err != nil {
		log.Print(err)
		http.Error(res, "failed to recieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// log.Print(handler.Filename)
	// service logic

	root, err := os.OpenRoot(".")
	if err != nil {
		log.Print(err)
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	dst, err := root.Create(time.Now().UTC().String())
	if err != nil {
		log.Print(err)
		http.Error(res, "failed to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		log.Print(err)
		http.Error(res, "failed to write into file", http.StatusInternalServerError)
		return
	}
}
