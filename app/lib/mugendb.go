package mugendb

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func InitDB() {
   db, err := gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
   checkErr(err, "mysql Open failed.")
   db.DB().Ping()
   db.DB().SetMaxIdleConns(10)
   db.DB().SetMaxOpenConns(100)
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

