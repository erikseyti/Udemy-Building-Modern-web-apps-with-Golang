package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erikseyti/udemy-go-course/pkg/config"
	"github.com/erikseyti/udemy-go-course/pkg/handlers"
	"github.com/erikseyti/udemy-go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app  config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
