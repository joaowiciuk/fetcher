package model

// Book holds information about a book.
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	ISBN  string `json:"isbn"`
	Likes int    `json:"likes"`
}
