package main

import (
	"fmt"
	"gobooks/internal/service"
)

func sum(x, y int) int {
	return x + y
}

func main() {
	book := service.Book{
		ID:     1,
		Title:  "Lord of the Rings",
		Author: "RR Tolkein",
		Genre:  "Fantasy",
	}

	fmt.Println(book.GetFullBook())
}
