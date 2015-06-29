package controllers

import (
	"bytes"
	"fmt"
	"github.com/revel/revel"
	//"goassimp/app/routes"
	"io"
	"os"
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

func (c *Convert) Convert(source[]byte) revel.Result {
	// 画像保存
	dstDir := "public/tmp"
	dstName := c.Params.Files["source"][0].Filename
	dst, err := os.Create(dstDir + "/" + dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, bytes.NewReader(source))
	if err != nil {
		panic(err)
	}

	// 情報登録
	// uuid, path, created_at
	// Zset順


	fmt.Println("Status: Successfully uploaded")

	return c.RenderJson(map[string]interface{}{
		"Status": "Successfully uploaded",
	})
}
