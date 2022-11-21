package db

import (
	"github.com/MaiMouri/go-gin-gorm-user-app/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Initialize() {
	// 宣言済みのグローバル変数dbをshort variable declaration(:=)で初期化しようとすると
	// ローカル変数dbを初期化することになるので注意する

	// DBのコネクションを接続する
	db, _ = gorm.Open("mysql", "tester:secret@tcp(db:3306)/test")
	//if err != nil {
	//  panic("failed to connect database\n")
	//}

	// ログを有効にする
	db.LogMode(true)

	// Task構造体(Model)を元にマイグレーションを実行する
	db.AutoMigrate(&models.Task{})
}

func Get() *gorm.DB {
	return db
}

// DBのコネクションを切断する
func Close() {
	db.Close()
}
