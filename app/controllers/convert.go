package controllers

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/revel/revel"
	"goassimp/lib/mgndb"
	"goassimp/app/models"

	"github.com/satori/go.uuid"
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

	// ログ保存
	con := models.ConvertLog{UUID: get8UUID(), Path: pth}
	status := "Status: Successfully uploaded"

	if err := mugendb.Db.Save(&con).Error; err != nil {
		errs := fmt.Sprintf("%s", err)
		status = "failed err" + errs
	}

	return c.RenderJson(map[string]interface{}{
		"Status": status,
	})
}

func get8UUID() string {
	u1 := uuid.NewV4()
	str := fmt.Sprintf("%s", u1)
	return str[0:8]
}


