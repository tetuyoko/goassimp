package controllers

import (
	"bytes"
	"fmt"
	"github.com/revel/revel"
	//"goassimp/app/routes"
	"github.com/satori/go.uuid"
	"goassimp/lib/mgnredis"
	"io"
	"os"
	"time"
)

type Convert struct {
	App
}

func (c *Convert) List() revel.Result {
	var id string = c.Params.Get("id")
	fmt.Println(id)
	// 情報出力
	// Zset順
	return c.Render()
}

func (c *Convert) Convert(source []byte) revel.Result {
	// 画像保存
	dstDir := "public/tmp"
	dstName := c.Params.Files["source"][0].Filename
	pth := dstDir + "/" + dstName
	dst, err := os.Create(pth)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, bytes.NewReader(source))
	if err != nil {
		panic(err)
	}

	// TODO: おとなしくDBにしよ。。
	// 情報登録
	// uuid, path, created_at
	// Zset順
	uuid := get8UUID()

	_, err = mgnredis.RedisDb.HSet(uuid, "path", pth)
	if err != nil {
		panic(err)
	}
	t := time.Now()

	str := fmt.Sprintf("%s", t.Unix())

	_, err = mgnredis.RedisDb.HSet(uuid, "created_at", str)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status: Successfully uploaded")

	return c.RenderJson(map[string]interface{}{
		"Status": "Successfully uploaded",
	})
}

func get8UUID() string {
	u1 := uuid.NewV4()
	str := fmt.Sprintf("%s", u1)
	return str[0:8]
}
