package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Simon-Busch/go_simple_api/internal/handlers"
	log "github.com/sirupsen/logrus" // Alias the package name
)


func main () {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Hanlder(r)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to start server")
	}
}
