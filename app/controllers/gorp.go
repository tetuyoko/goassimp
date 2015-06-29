package controllers

import (
    "database/sql"
    "github.com/coopernurse/gorp"
    _ "github.com/mattn/go-sqlite3"
    r "github.com/revel/revel"
    "goassimp/app/models" // revel new APP_NAME の APP_NAME
)

var (
    DbMap *gorp.DbMap // このデータベースマッパーからSQLを流す
)

func InitDB() {
    db, err := sql.Open("sqlite3", "./app.db")
    if err != nil {
        panic(err.Error())
    }
    DbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

    // ここで好きにテーブルを定義する
    t := DbMap.AddTable(models.User{}).SetKeys(true, "Id")
    t.ColMap("Name").MaxSize = 20

    DbMap.CreateTables()
}

type GorpController struct {
    *r.Controller
    Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
    txn, err := DbMap.Begin()
    if err != nil {
        panic(err)
    }
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


