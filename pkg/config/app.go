package config

import (
	"github.com/jinzhu/gorm"
	 _"github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Kumar@123@tcp(127.0.0.1:3306)/Book?charset=utf8mb4&parseTime=True&loc=Local") 
	if err != nil {
		panic(err)
	}
	db = d

}
func GetDB() *gorm.DB {
	return db
}


/* gorm for intrecting to our database
gorilla mux for routing
*/