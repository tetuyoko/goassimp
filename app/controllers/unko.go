package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"goassimp/app/models"
	"goassimp/app/routes"

	"goassimp/app/mgnredis"
)

type Unko struct {
	App
}

func (c *Unko) List() revel.Result {
	greeting := "Super Mother Fucing Hage!!"
	unko := "unko"
	return c.RenderJson(map[string]interface{}{
		"Count":    4,
		"unko":     unko,
		"greeting": greeting,
		"Status":   "Successfully uploaded",
	})
}

func (c *Unko) Show(id int) revel.Result {
	oem := models.Oembed{Version: 1, Type: "fuga"}
	return c.RenderJson(oem)
}

func (c *Unko) Cancel(id int) revel.Result {
	c.Flash.Success(fmt.Sprintln("Booking cancelled for confirmation number", id))
	return c.Redirect(routes.Unko.Index())
	//return c.RenderText("unko")
}

func (c *Unko) Index() revel.Result {
	greeting := "Super Mother Fucing Hage!!"
	info, err := mgnredis.RedisDb.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", info)
	return c.Render(greeting, "unko")
}

//func (c *Single) HandleUpload(avatar []byte) revel.Result {
//	// Validation rules.
//	c.Validation.Required(avatar)
//	c.Validation.MinSize(avatar, 2*KB).
//		Message("Minimum a file size of 2KB expected")
//	c.Validation.MaxSize(avatar, 2*MB).
//		Message("File cannot be larger than 2MB")
//
//	// Check format of the file.
//	conf, format, err := image.DecodeConfig(bytes.NewReader(avatar))
//	c.Validation.Required(err == nil).Key("avatar").
//		Message("Incorrect file format")
//	c.Validation.Required(format == "jpeg" || format == "png").Key("avatar").
//		Message("JPEG or PNG file format is expected")
//
//	// Check resolution.
//	c.Validation.Required(conf.Height >= 150 && conf.Width >= 150).Key("avatar").
//		Message("Minimum allowed resolution is 150x150px")
//
//	// Handle errors.
//	if c.Validation.HasErrors() {
//		c.Validation.Keep()
//		c.FlashParams()
//		return c.Redirect(routes.Single.Upload())
//	}
//
//	return c.RenderJson(FileInfo{
//		ContentType: c.Params.Files["avatar"][0].Header.Get("Content-Type"),
//		Filename:    c.Params.Files["avatar"][0].Filename,
//		RealFormat:  format,
//		Resolution:  fmt.Sprintf("%dx%d", conf.Width, conf.Height),
//		Size:        len(avatar),
//		Status:      "Successfully uploaded",
//	})
//}
