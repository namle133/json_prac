package service

import "github.com/namle133/json_prac.git/JSON_PRAC/domain"

type IBook interface {
	Create(item domain.Book)
	ReadAll() []domain.Book
	ReadOne(id string) domain.Book
	Update(id string, item domain.Book)
	Delete(id string) string
}
