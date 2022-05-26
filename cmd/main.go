package main

import (
	"github.com/gorilla/mux"
	"github.com/namle133/json_prac.git/JSON_PRAC/http/decode"
	"github.com/namle133/json_prac.git/JSON_PRAC/http/encode"
	"github.com/namle133/json_prac.git/JSON_PRAC/service"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	b := service.Book{}
	var i service.IBook = &b

	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		item := decode.CreateRequest(r)
		i.Create(item)
		encode.CreateResponse(w, item)
	}).Methods("POST")

	r.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		books := i.ReadAll()
		encode.ReadAllResponse(w, books)
	}).Methods("GET")

	r.HandleFunc("/read/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := decode.ReadOne(r)
		b := i.ReadOne(id)
		encode.ReadOneResponse(w, b)
	}).Methods("GET")

	r.HandleFunc("/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, item := decode.Update(r)
		i.Update(id, item)
		encode.UpdateOneResponse(w)
	}).Methods("PUT")

	r.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := decode.DeleteOne(r)
		idStr := i.Delete(id)
		encode.DeleteOneResponse(w, idStr)
	}).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
