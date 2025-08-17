package handlers

import (
	"net/http"
	"os"
)

func HandleRoot(res http.ResponseWriter, req *http.Request) {

	path := "../index.html"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(res, "index.html is not found", http.StatusInternalServerError)
		return
	}
	http.ServeFile(res, req, path)
}
