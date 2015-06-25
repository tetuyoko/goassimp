package redis

import (
	"log"

	//	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

type RedisDB struct {
	pool *pools.ResourcePool
}

var redisDB *RedisDB

type pooledConn struct {
	ResourceConn
	pool *pools.ResourcePool
}

func (wc pooledConn) Close() {
	wc.pool.Put(wc.ResourceConn)
}

func newRedisDB(pool *pools.ResourcePool) *RedisDB {
	return &RedisDB{pool}
}

func InitRedis() {
	p := newPool(":6379")
	defer p.Close()
	redisDB = newRedisDB(p)
	defer redisDB.Close()
	myPing()
}

func myPing() {
	//ctx := context.TODO()
	//r, err := redisDB.pool.Get(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer redisDB.pool.Put(r)

	//c := r.(ResourceConn)

	//pc := pooledConn{c, redisDB.pool}

	pc, err := redisDB.conn()
	if err != nil {
		pc.Close()
		log.Fatal(err)
	}
	n, err := pc.Do("INFO")
	log.Printf("info=%s", n)
	pc.Close()
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

func (db *RedisDB) Close() {
	db.pool.Close()
}

//func Ping( p pool) {
//	ctx := context.TODO()
//	r, err := p.Get(ctx)
//	if err != nil {
//		log.Fatal(err)
//		//return nil, err
//	}
//	defer p.Put(r)
//	c := r.(ResourceConn)
//	n, err := c.Do("INFO")
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("info=%s", n)
//}

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
