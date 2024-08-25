package service

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

// getFullBook returns a string representation of a Book.
//
// It takes no parameters besides the Book instance itself.
// Returns a string in the format "ID Title Author Genre".
func (b Book) GetFullBook() string {
	return b.Title + " by " + b.Author + " Genre " + b.Genre
}
