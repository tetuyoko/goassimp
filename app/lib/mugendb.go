package mugendb

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var (
    Db  gorm.DB
)

func InitDB() {
    var err error
    Db, err = gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
    checkErr(err, "mysql Open failed.")
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

