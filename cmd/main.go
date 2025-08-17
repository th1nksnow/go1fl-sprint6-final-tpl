package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, `serv `, log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
