package app

import (
	"gorm.io/gorm"
	. "wiki/pkg/model"
)

type IBookORM interface {
	GetBooks(db *gorm.DB) ([]*Book, error)
	AddBook(db *gorm.DB, book *Book) error
}

type BooksORM struct {
}

func (service *BooksORM) GetBooks(db *gorm.DB) ([]*Book, error) {
	books := make([]*Book, 0)
	err := db.Model(&Book{}).Find(books).Error
	if err != nil {
		return nil, err
	}
	return books, err
}

func (service *BooksORM) AddBook(db *gorm.DB, book *Book) error {
	err := db.Session(&gorm.Session{FullSaveAssociations: true, CreateBatchSize: 100}).Model(Book{}).Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func NewBooksOrm() IBookORM {
	return &BooksORM{}
}
