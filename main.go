package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepository
	mux         *http.ServeMux
	templateDir string
	tp          *TemplateRenderer
}

func main() {
	db, err := connectToDatabase("users_database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog:    log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:     log.New(os.Stdout, "INFO\t", log.Ltime|log.LstdFlags),
		userRepo:    NewSQLUserRepository(db),
		templateDir: "templates",
	}

	app.tp = NewTemplateRenderer(app.templateDir, true)

	log.Println("Server starting on :8080")
	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
