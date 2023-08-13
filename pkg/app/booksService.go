package app

import (
	"log"
	"wiki/pkg/config"
	. "wiki/pkg/model"
)

var booksOrm = NewBooksOrm()

func GetAllBooks(config *config.Config) ([]*Book, error) {
	books, err := booksOrm.GetBooks(config.DB)
	if err != nil {
		log.Println("Error while getting books from DB", err)
		return nil, err
	}
	return books, nil
}

func AddBooks(config *config.Config, book *Book) error {
	err := booksOrm.AddBook(config.DB, book)
	if err != nil {
		log.Println("Error while adding book", err)
		return err
	}
	return nil
}
