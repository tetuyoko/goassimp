package mugendb

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

	"goassimp/app/models"
)

var (
	Db gorm.DB
)

func InitDB(user string, password string, host string, dbname string) {
	var err error

	// connect mysql
	// "root:@/godb?charset=utf8&parseTime=True&loc=Local"
	op := user + ":" + password + "@" + host + "/?charset=utf8&parseTime=True&loc=Local"
	log.Println(op)
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
	Db.AutoMigrate(&models.User{}, &models.ConvertLog{})
	insertUser()
	insertConvertLog()
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func insertUser() {
	Db.Create(&models.User{Name: "Jinzha"})
	user := models.User{}
	Db.First(&user)
	log.Println(user.Name, nil)
}

func insertConvertLog() {
	Db.Create(&models.ConvertLog{UUID: "Jingia", Path: "path/to/tmp"})
	c := models.ConvertLog{}
	Db.First(&c)
	log.Println(c.UUID, nil)
	log.Println(c.Path, nil)
}
