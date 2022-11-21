package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	// gorm
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	// gorm.Model
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// type Todo struct {
// 	gorm.Model
// 	Name    string `json:"name"`
// 	Content string `json:"content"`
// 	Status  int    `json:"status"`
// }

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	return db
}

// 一覧取得
func dbGetAll() *gorm.DB {
	db := GetDBConnection()

	// defer db.Close()
	println(db)
	var users []User
	result := db.Find(&users)
	println("user:", result.Value)

	return result
}

// func getUsers() []*User {
func getUsers() []*User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

// Create
func createUser(name string, status int) []error {
	db := GetDBConnection()
	defer db.Close()
	// Insert処理
	if err := db.Create(&User{
		ID:     0,
		Name:   name,
		Status: 0,
	}).GetErrors(); err != nil {
		return err
	}
	return nil
}

func main() {
	db.Initialize()
	deger db.Close()

	router.Router()
	// r := gin.Default()
	// r.LoadHTMLGlob("templates/index.html")

	// 一覧
	// r.GET("/userlist", func(c *gin.Context) {
	// 	// users := dbGetAll()
	// 	users := getUsers()

	// 	if users != nil {
	// 		fmt.Println("Attempt to read users and Suceed to Load users!!", users[0])
	// 	}

	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"user": users,
	// 	})
	// })

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/users", func(c *gin.Context) {
	// 	// users := getUsers()
	// 	users := dbGetAll()
	// 	if users != nil {
	// 		fmt.Printf("Load users Success!!")
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user": users,
	// 	})
	// })

	// r.POST("/new", func(c *gin.Context) {
	// 	var form User
	// 	db := GetDBConnection()
	// 	defer db.Close()

	// 	// バリデーション処理
	// 	if err := c.Bind(&form); err != nil {
	// 		c.HTML(http.StatusBadRequest, "index.html", gin.H{"err": err})
	// 		c.Abort()
	// 	} else {
	// 		name := c.PostForm("name")
	// 		println(name)
	// 		status := 0
	// 		// 登録ユーザーが重複していた場合にはじく処理
	// 		if err := createUser(name, status); err != nil {
	// 			c.HTML(http.StatusBadRequest, "index.html", gin.H{"err": err})
	// 		}
	// 		c.Redirect(302, "/userlist")
	// 	}
	// })

	http.HandleFunc("/users", userPage)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	r.Run(":8080")
}
