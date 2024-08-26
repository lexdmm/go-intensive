package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
)

type BookHandlers struct {
	bookService *service.BookService
}

func NewBookHandlers(bookService *service.BookService) *BookHandlers {
	return &BookHandlers{bookService: bookService}
}

/**
 * @api {get} /books Get all books
 * @apiName GetBooks
 * @apiGroup Books
 */
func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetFullBook()
	if err != nil {
		http.Error(w, "Failed to get books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/**
 * @api {post} /book Create book
 * @apiName CreateBook
 * @apiGroup Books
 */
func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.bookService.CreateBook(&book)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

/**
 * @api {get} /book/:id Get book by id
 * @apiName GetBook
 * @apiGroup Books
 */
func (h *BookHandlers) GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	id, err := strconv.Atoi(bookId) // convert string to int
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.bookService.GetBook(id)
	if err != nil {
		http.Error(w, "Failed to get book by id", http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

/**
 * @api {put} /book/:id Update book
 * @apiName UpdateBook
 * @apiGroup Books
 */
func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	id, err := strconv.Atoi(bookId) // convert string to int
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	var book service.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = id

	if err = h.bookService.UpdateBook(&book); err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

/**
 * @api {delete} /book/:id Delete book
 * @apiName DeleteBook
 * @apiGroup Books
 */
func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.PathValue("id")
	id, err := strconv.Atoi(bookId) // convert string to int
	if err != nil {
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	if err = h.bookService.DeleteBook(id); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
