package decode

import (
	"encoding/json"
	"fmt"
	"github.com/namle133/json_prac.git/JSON_PRAC/domain"
	"net/http"
)

func CreateRequest(r *http.Request) domain.Book {
	var b domain.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println(err)
		return domain.Book{ID: "", Title: ""}
	}
	return b
}

func ReadOne(r *http.Request) string {
	id := r.URL.Path[6:]
	return id
}

func Update(r *http.Request) (string, domain.Book) {
	id := r.URL.Path[8:]
	var b domain.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Println(err)
		return "", domain.Book{ID: "", Title: ""}
	}
	return id, b
}

func DeleteOne(r *http.Request) string {
	id := r.URL.Path[8:]
	return id
}
