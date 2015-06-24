package app

import (
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

type DB struct {
	p *pools.ResourcePool
}

func newDB(p *pools.ResourcePool) *DB {
	return &DB{p}
}

type pooledConn struct {
	*ResourceConn
	p *pools.ResourcePool
}

func (wc *pooledConn) Close() {
	wc.p.Put(wc.ResourceConn)
}

func (db *DB) conn() (*pooledConn, error) {
	ctx := context.TODO()
	r, err := db.p.Get(ctx)
	if err != nil {
		return nil, err
	}
	c := r.(*ResourceConn)
	return &pooledConn{c, db.p}, nil
}

func (db *DB) Close() {
	db.p.Close()
}

func (db *DB) Ping() (string, error) {
	c, err := db.conn()
	if err != nil {
		return "", err
	}
	defer c.Close()

	name, err := redis.String(c.Do("INFO"))
	if err != nil {
		return "", err
	}
	return name, err
}

//func (db *DB) LoadUser(id int) (*User, error) {
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
