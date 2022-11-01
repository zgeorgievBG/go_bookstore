package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zgeorgievBG/go_bookstore/pkg/config"
)

var db * gorm.DB

type Book struct{
	gorm.Model
	// ID 		int 		`json:"id"`
	Title 	string 		`json:"title"`
	Author 	string 		`json:"author"`
	Price 	int 		`json:"price"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db

}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book

}
