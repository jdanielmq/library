package book

import (
	"fmt"
)

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

// si se quiere dejar privado los atributos, se deben cambiar a minisculas
type Book struct {
	Title  string //title
	Author string //author
	Pages  int    //pages
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}

}

func (b *Book) SetTitle(title string) {
	b.Title = title
}
func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s \nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s \nAuthor: %s\nPages: %d\nEditorial: %s \nLevel: %s\n",
		b.Title, b.Author, b.Pages, b.editorial, b.level)
}
