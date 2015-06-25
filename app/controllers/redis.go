package controllers

import (
    "fmt"
    "github.com/revel/revel"
    "goassimp/app/mgnredis"
)

// redisの疎通テスト用

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
        "Status":   "Success",
    })
}

func (c *Redis) Set() revel.Result {
    info, err := mgnredis.RedisDb.Set(c.Params.Values["key"][0], c.Params.Values["val"][0])
    if err != nil {
        panic(err)
    }
    str := fmt.Sprintf("%s", info)

    return c.RenderJson(map[string]interface{}{
        "reply":    str,
        "Status":   "Success",
    })
}

func (c *Redis) Get(key string) revel.Result {
    return c.RenderJson(map[string]interface{}{
        "key":    key,
        "val":    "val",
        "Status":   "Success",
    })
}
