package main

import (
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
		Handler:           app.routes(),
	}

	return srv.ListenAndServe()
}
