package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicPath))))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/login", app.login)
	mux.HandleFunc("/register", app.register)
	mux.HandleFunc("/contact", app.contact)
	mux.HandleFunc("/about", app.about)

	return mux
}
