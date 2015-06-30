package mugendb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goassimp/app/models"
	"log"
)

var (
	Db gorm.DB
)

func InitDB() {
	fmt.Println("called")
	var err error

	// connect mysql
	Db, err = gorm.Open("mysql", "root:@/?charset=utf8&parseTime=True&loc=Local")
	checkErr(err, "mysql Open failed.")

	// create db if not exists
	Db.Exec("CREATE DATABASE IF NOT EXISTS godb DEFAULT CHARACTER SET utf8;")

	// connect db
	Db, err = gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
	checkErr(err, "mysql Connect failed.")

	// Db.DB().Ping()
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	// migration
	Db.AutoMigrate(&models.User{})

	// get
	//Db.Create(&models.User{Name: "Jinzhu"})
	//user := models.User{}
	//Db.First(&user)
	//log.Fatalln(user.Name, nil)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
