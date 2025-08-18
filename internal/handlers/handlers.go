package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
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

	data, err := io.ReadAll(file)
	if err != nil {
		log.Print(err)
		http.Error(res, "failed to read from file", http.StatusBadRequest)
		return
	}

	if string(data) == "" {
		http.Error(res, "recieved file is empty", http.StatusPaymentRequired)
		return
	}

	convertedString := service.Convert(string(data))

	err = os.WriteFile(time.Now().UTC().String(), []byte(convertedString), 0755)
	if err != nil {
		log.Print(err)
http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}
	res.Write([]byte(convertedString))
}
