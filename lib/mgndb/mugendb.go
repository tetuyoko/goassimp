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

//		mugendb.InitDB(user, password, host, dbname)
func InitDB(user string, password string, host string, dbname string) {
	fmt.Println("called")
	var err error

	// connect mysql
	// "root:@/godb?charset=utf8&parseTime=True&loc=Local"
	op := user + ":" + password + "@" + host + "/?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(op)
	Db, err = gorm.Open("mysql", op)
	checkErr(err, "mysql Open failed.")

	// create db if not exists
	Db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname + " DEFAULT CHARACTER SET utf8;")

	// connect db
	op = user + ":" + password + "@" + host + "/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql",op )
	checkErr(err, "mysql Connect failed.")

	// Db.DB().Ping()
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	// migration
	Db.AutoMigrate(&models.User{})
	insertUser()
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func insertUser() {
	// get
	Db.Create(&models.User{Name: "Jinzhu"})
	user := models.User{}
	Db.First(&user)
	log.Fatalln(user.Name, nil)
}
