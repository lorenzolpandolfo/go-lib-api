package main

import (
	"fmt"
	"go-lib-api/internal/controller"
	"go-lib-api/internal/infrastructure/db"
	"go-lib-api/internal/repository"
	"go-lib-api/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func connect() *sqlx.DB {
	database, err := db.NewConnection()

	if err != nil {
		log.Fatal("Could not connect to DB: &v", err)
	}

	return database
}

func initDatabase(database *sqlx.DB) {
		schema := `CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		author TEXT NOT NULL,
		title TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	database.MustExec(schema)
}

func main() {
	database := connect()
	defer database.Close()

	initDatabase(database)

	repo := repository.NewBookRepository(database)
	svc := service.NewBookService(repo)
	ctrl := controller.NewBookController(svc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/books", ctrl.Create)
	r.Get("/books/{id}", ctrl.Get)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)

}