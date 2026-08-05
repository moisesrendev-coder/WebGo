package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/contact", app.contact)
	mux.HandleFunc("/about", app.about)

	return mux
}
