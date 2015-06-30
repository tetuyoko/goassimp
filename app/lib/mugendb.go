package mugendb

import (
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

func InitDB() {
   db, _ := gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
   db.DB().Ping()
}
