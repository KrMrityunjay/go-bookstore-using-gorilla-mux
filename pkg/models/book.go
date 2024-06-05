/*
	user interacts with the routes and routes sends the

control to the controllers where we have our logic ,and
controllers will have to perform some opeerations with
the databse ,the operations in the databse have to reside
in models i.e. book.go so that means we will have to have
different functions for different controllers that we create
*/
package models

import (
	"github.com/KrMrityunjay/GO-BOOKSTORE/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
	
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db :=db.Where("ID=?" , Id).Find(&getBook)
	return &getBook ,db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?",ID).Delete(book)
	return book
}
