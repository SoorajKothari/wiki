package router

import (
	"encoding/json"
	"log"
	"net/http"
	"wiki/pkg/app"
	"wiki/pkg/model"
)

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := app.GetAllBooks(globalConfig)
	if err != nil {
		log.Println("GET Book failed to process", err.Error())
		writer.Write([]byte("Failed to get books"))
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(writer).Encode(books)
	writer.WriteHeader(http.StatusOK)
}

func AddBook(writer http.ResponseWriter, request *http.Request) {
	newBook := model.Book{}
	json.NewDecoder(request.Body).Decode(newBook)
	err := app.AddBooks(globalConfig, &newBook)

	if err != nil {
		log.Println("POST BOOK failed to process", err.Error())
		writer.Write([]byte("Failed to add book"))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Write([]byte("Success!"))
	writer.WriteHeader(http.StatusOK)
}
