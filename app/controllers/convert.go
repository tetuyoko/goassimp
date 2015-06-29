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
	return c.Render()
}

func (c *Convert) Convert(source[]byte) revel.Result {
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

	fmt.Println("Status: Successfully uploaded")

	return c.RenderJson(map[string]interface{}{
		"Status": "Successfully uploaded",
	})
}
