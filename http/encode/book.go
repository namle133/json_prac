package encode

import (
	"encoding/json"
	"fmt"
	"github.com/namle133/json_prac.git/JSON_PRAC/domain"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, item domain.Book) {
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadAllResponse(w http.ResponseWriter, books []domain.Book) {
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadOneResponse(w http.ResponseWriter, b domain.Book) {
	err := json.NewEncoder(w).Encode(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateOneResponse(w http.ResponseWriter) {
	fmt.Fprintln(w, "Updated Successfully!")
}

func DeleteOneResponse(w http.ResponseWriter, id string) {
	fmt.Fprintln(w, fmt.Sprintf("Deleted successfully: %s", id))
}
