package controllers

import (
	"bytes"
	"fmt"
	"github.com/revel/revel"
	"goassimp/app/routes"
	"io"
	"os"
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

	fmt.Println("Status: Successfully uploaded")

	return c.Redirect(routes.Single.Upload())

	//	return c.RenderJson(map[string]interface{}{
	//		"Status": "Successfully uploaded"
	//	})
}
