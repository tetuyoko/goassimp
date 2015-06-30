package controllers

import (
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/go-sql-driver/mysql"
    r "github.com/revel/revel"
    "goassimp/app/models" // revel new APP_NAME の APP_NAME
    "log"
)

var (
    DbMap *gorp.DbMap // このデータベースマッパーからSQLを流す
)

func InitDB() {
    // connect mysql
    db, err := sql.Open("mysql", "root:@/")
    checkErr(err, "sql.Open failed")
    DbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

    // create db if not exists
    _, err = DbMap.Exec("CREATE DATABASE IF NOT EXISTS godb DEFAULT CHARACTER SET utf8;")
    checkErr(err, "create db failed")

    // connect db
    db, err = sql.Open("mysql", "root:@/godb")
    checkErr(err, "sql.Open Failed.")
    DbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

    // ここで好きにテーブルを定義する
    // add table
     DbMap.AddTableWithName(models.User{}, "users").SetKeys(true, "Id")

    DbMap.CreateTables()
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

type GorpController struct {
    *r.Controller
    Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
    txn, err := DbMap.Begin()
    checkErr(err, "DbMap.Begin() failed")
    c.Txn = txn
    return nil
}

func (c *GorpController) Commit() r.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

func (c *GorpController) Rollback() r.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}


