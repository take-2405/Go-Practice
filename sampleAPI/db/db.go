package db

import (
	"github.com/git/Go-Practice/sampleAPI/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	db  *gorm.DB
	err error
)

//Init DBの初期化
func Init() {
	
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=gorm dbname=gorm password=gorm sslemode=gorm password=gorm")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.User{})
	
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
