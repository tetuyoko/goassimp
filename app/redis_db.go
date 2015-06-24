package app

import (
	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

type RedisDB struct {
	pool *pools.ResourcePool
}

type pooledConn struct {
	*ResourceConn
	pool *pools.ResourcePool
}

func (wc *pooledConn) Close() {
	wc.pool.Put(wc.ResourceConn)
}

func InitRedis() {
	pool = newPool(":6379")
	defer pool.Close()
	db := newRedisDB(pool)
	defer db.Close()
}

func newRedisDB(pool *pools.ResourcePool) *RedisDB {
	return &RedisDB{pool}
}

func (db *RedisDB) conn() (*pooledConn, error) {
	ctx := context.TODO()
	r, err := db.pool.Get(ctx)
	if err != nil {
		return nil, err
	}
	c := r.(*ResourceConn)
	return &pooledConn{c, db.pool}, nil
}
func (db *RedisDB) Close() {
	db.pool.Close()
}

func (db *RedisDB) Ping() (string, error) {
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

//func (db *RedisDB) LoadUser(id int) (*User, error) {
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
