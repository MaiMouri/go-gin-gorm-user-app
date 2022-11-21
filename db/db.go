package db

import (
	"github.com/hanadaUG/go-gin-gorm-todo-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

func Initialize() {
	// 宣言済みのグローバル変数dbをshort variable declaration(:=)で初期化しようとすると
	// ローカル変数dbを初期化することになるので注意する

	// DBのコネクションを接続する
	db, _ = gorm.Open("mysql", "task.db")
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
