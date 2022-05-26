package service

import "github.com/namle133/json_prac.git/JSON_PRAC/domain"

type Book struct {
}

type AllBooks []domain.Book

var books = AllBooks{}

func (b *Book) Create(item domain.Book) {
	books = append(books, item)
}

func (b *Book) ReadAll() []domain.Book {
	return books
}

func (b *Book) ReadOne(id string) domain.Book {
	for i, singleID := range books {
		if singleID.ID == id {
			return books[i]
		}
	}
	return domain.Book{
		ID:    "",
		Title: "",
	}
}

func (b *Book) Update(id string, item domain.Book) {
	for i, singleID := range books {
		if singleID.ID == id {
			books[i].Title = item.Title
			return
		}
	}
}

func (b *Book) Delete(id string) string {
	for i, singleID := range books {
		if singleID.ID == id {
			books = append(books[:i], books[i+1:]...)
			return id
		}
	}
	return ""
}
