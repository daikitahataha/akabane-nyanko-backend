package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	// TODO DB設定は環境変数に置き換える必要あり
	db, err = gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres dbname=akabane_nyanko password=password sslmode=disable")
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
