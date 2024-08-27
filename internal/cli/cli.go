package cli

import (
	"fmt"
	"gobooks/internal/service"
	"os"
	"strconv"
	"time"
)

type BookCLI struct {
	bookService *service.BookService
}

func NewCLI(bookService *service.BookService) *BookCLI {
	return &BookCLI{bookService: bookService}
}

func (cli *BookCLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gobooks <command> [arguments]")
		return
	}

	command := os.Args[1]
	switch command {
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks search <book title>")
			return
		}
		bookName := os.Args[2]
		cli.SearchBooks(bookName)

	case "simulate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gobooks simulate <book_id> <book_id> <book_id> ...")
			return
		}
		bookIDs := os.Args[2:]
		cli.simulateReadingBooks(bookIDs)
	}
}

func (cli *BookCLI) SearchBooks(bookName string) {
	books, err := cli.bookService.SearchBooksByName(bookName)
	if err != nil {
		fmt.Println("Error searching books:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books found")
		return
	}

	fmt.Printf("%d books found\n", len(books))
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Genre: %s\n", book.ID, book.Title, book.Author, book.Genre)
	}
}

func (cli *BookCLI) simulateReadingBooks(bookIDs []string) {
	var bookIDsInt []int

	for _, bookID := range bookIDs {
		id, err := strconv.Atoi(bookID)
		if err != nil {
			fmt.Println("Invalid book ID:", bookID)
			continue
		}
		bookIDsInt = append(bookIDsInt, id)
	}
	responses := cli.bookService.SimulateMultipleReadings(bookIDsInt, 5*time.Second)
	for _, response := range responses {
		fmt.Println(response)
	}
}
