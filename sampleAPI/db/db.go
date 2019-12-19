package db

import (
	"model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

//Init DBの初期化
func Init() {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslemode=password")
	if err != nil {
		panic(err)
	}
	autoMigration()
}

//GetDB データベースを返す
func GetDB() *gorm.DB {
	return db
}

//Close DBとの接続を切る
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}

}

func autoMigration() {
	db.AutoMigrate(&model.User{})
}
