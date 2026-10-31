package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stderr, "server: ", log.LstdFlags)
	srv := server.New(logger)
	logger.Println("Сервер запускается на %s", srv.HTTPServer.Addr)
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}
