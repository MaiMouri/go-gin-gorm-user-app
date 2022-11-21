package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Status string
}

// func GetAll() (datas []User) {
// 	result := Db.Find(&datas)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

// func GetOne(id int) (data User) {
// 	result := Db.First(&data, id)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return
// }

// func (b *User) Create() {
// 	result := Db.Create(b)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// }

// func (b *User) Update() {
// 	result := Db.Save(b)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// }

// func (b *User) Delete() {
// 	result := Db.Delete(b)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// }
