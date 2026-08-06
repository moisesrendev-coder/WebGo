package main

import (
	"net/http"
)

var htmlContent = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>%s</title>
</head>
<body>
	<h1>About this simple server</h1>
	<p>This server is built using Go.</p>
	<p>%s</p>
</body>
</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	app.render(w, "index.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	app.render(w, "login.html", nil)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	app.render(w, "register.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.html", nil)
}
