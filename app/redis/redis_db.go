package redis

import (
	//	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
	"github.com/garyburd/redigo/redis"
)

type RedisDB struct {
	pool *pools.ResourcePool
}

var redisDB *RedisDB

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
	defer p.Close()
	redisDB = newRedisDB(p)
	defer redisDB.Close()
	redisDB.Ping()
}

func newRedisDB(pool *pools.ResourcePool) *RedisDB {
	return &RedisDB{pool}
}

func (db *RedisDB) Ping() (string, error) {
	pc, err := db.conn()
	defer pc.Put()
	if err != nil {
		return "", err
	}
	info, err :=  redis.String(pc.Do("INFO"))
	//info, err := redis.String(pc.Do("INFO"))
	if err != nil {
	   return "", err
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
