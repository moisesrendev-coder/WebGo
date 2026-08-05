package main

import (
	"fmt"
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

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	aboutContent := fmt.Sprintf(htmlContent, "About", "<h1>Hello, Welcome to the about page</h1>")
	_, _ = w.Write([]byte(aboutContent))
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("%s %s", r.Method, r.URL.Path)
	contactContent := fmt.Sprintf(htmlContent, "Contact", "<h1>Hello, Welcome to the contact page</h1>")
	_, _ = w.Write([]byte(contactContent))
}
