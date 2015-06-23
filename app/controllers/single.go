package controllers

import (
	"bytes"
	"github.com/revel/revel"
	//	"goassimp/app/routes"
	//	"image"
	//	_ "image/jpeg"
	//	_ "image/png"
	"io"
	"os"
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

type Single struct {
	App
}

func (c *Single) Upload() revel.Result {
	return c.Render()
}

func (c *Single) HandleUpload(avatar []byte) revel.Result {
	dstDir := "public/tmp"
	dstName := c.Params.Files["avatar"][0].Filename
	dst, err := os.Create(dstDir + "/" + dstName)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, bytes.NewReader(avatar))
	if err != nil {
		panic(err)
	}

	return c.RenderJson(map[string]interface{}{
		"Status": "Successfully uploaded",
	})
}
