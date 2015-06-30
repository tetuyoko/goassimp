package controllers

// redisの疎通テスト用
import (
	"fmt"

	"github.com/revel/revel"
	"goassimp/lib/mgnredis"
)

type Redis struct {
	App
}

func (c *Redis) Ping() revel.Result {
	info, err := mgnredis.RedisDb.Ping()
	if err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%s", info)

	return c.RenderJson(map[string]interface{}{
		"Ping":   str,
		"Status": "Success",
	})
}

func (c *Redis) Set() revel.Result {
	info, err := mgnredis.RedisDb.Set(c.Params.Values["key"][0], c.Params.Values["val"][0])
	if err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%s", info)

	return c.RenderJson(map[string]interface{}{
		"reply": str,
	})
}

func (c *Redis) Hset() revel.Result {
	info, err := mgnredis.RedisDb.HSet(
		c.Params.Values["key"][0],
		c.Params.Values["field"][0],
		c.Params.Values["val"][0])
	if err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%s", info)

	return c.RenderJson(map[string]interface{}{
		"reply": str,
	})
}

func (c *Redis) Hget(key string) revel.Result {
	info, err := mgnredis.RedisDb.HGetAll(key)
	if err != nil {
		panic(err)
	}

	//str := fmt.Sprintf("%s", info)
	return c.RenderJson(map[string]interface{}{
		"key":          key,
		"field_values": info,
		"Status":       "Success",
	})
}

func (c *Redis) Get(key string) revel.Result {
	info, err := mgnredis.RedisDb.Get(key)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", info)
	return c.RenderJson(map[string]interface{}{
		"key":    key,
		"val":    str,
		"Status": "Success",
	})
}
