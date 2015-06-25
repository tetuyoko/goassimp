package mgnredis

import (
	"fmt"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

type RedisDB struct {
	pool *pools.ResourcePool
}

var RedisDb *RedisDB

type pooledConn struct {
	ResourceConn
	pool *pools.ResourcePool
}

func (wc pooledConn) Put() {
	wc.pool.Put(wc.ResourceConn)
}

func (db *RedisDB) Close() {
	db.pool.Close()
}

func InitRedis() {
	p := newPool(":6379")
	//defer p.Close()
	RedisDb = newRedisDB(p)
	//defer RedisDb.Close()
	info, err := RedisDb.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", info)
}

func newRedisDB(pool *pools.ResourcePool) *RedisDB {
	return &RedisDB{pool}
}

func (db *RedisDB) Ping() (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()
	info, err := pc.Do("INFO")
	//info, err := redis.String(pc.Do("INFO"))
	if err != nil {
		panic(err)
	}
	return info, err
}

func (db *RedisDB) Get() (interface{}, error) {
	pc, err := db.conn()
	if err != nil {
		panic(err)
	}
	defer pc.Put()
	info, err := pc.Do("INFO")
	//info, err := redis.String(pc.Do("INFO"))
	if err != nil {
		panic(err)
	}
	return info, err
}

func (db *RedisDB) conn() (*pooledConn, error) {
	ctx := context.TODO()
	r, err := db.pool.Get(ctx)
	if err != nil {
		return nil, err
	}
	c := r.(ResourceConn)
	return &pooledConn{c, db.pool}, nil
}

// func (db *RedisDB) LoadUser(id int) (*User, error) {
//	c, err := db.conn()
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	name, err := redis.String(c.Do("GET", UserKey(id)))
//	if err != nil {
//		return nil, err
//	}
//
//	return &User{ID: id, Name: name}, nil
//}
//func (db *DB) SaveUser(u *User) error {
//	c, err := db.conn()
//	if err != nil {
//		return err
//	}
//	defer c.Close()
//
//	_, err = c.Do("SET", u.Key(), u.Name)
//	return err
//}