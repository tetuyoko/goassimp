package mugendb

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

var (
    Db  gorm.DB
)

func InitDB() {
    fmt.Println("called")
    var err error

    // connect mysql
    Db, err = gorm.Open("mysql", "root:@/?charset=utf8&parseTime=True&loc=Local")
    checkErr(err, "mysql Open failed.")

    // create db if not exists
    Db.Exec( "CREATE DATABASE IF NOT EXISTS godb DEFAULT CHARACTER SET utf8;")

    // connect db
    Db, err = gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
    checkErr(err, "mysql Connect failed.")

    Db.DB().Ping()
    Db.DB().SetMaxIdleConns(10)
    Db.DB().SetMaxOpenConns(100)
    Db.SingularTable(true)

     // migration
     // db.CreateTable(&User{})
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

