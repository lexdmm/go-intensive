package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, genre TEXT)"); err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)
	bookHandlers := web.NewBookHandlers(bookService)

	router := http.NewServeMux()
	router.HandleFunc("GET /books", bookHandlers.GetBooks)
	router.HandleFunc("POST /book", bookHandlers.CreateBook)
	router.HandleFunc("GET /book/{id}", bookHandlers.GetBook)
	router.HandleFunc("PUT /book/{id}", bookHandlers.UpdateBook)
	router.HandleFunc("DELETE /book/{id}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8083", router)
}
