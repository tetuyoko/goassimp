package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"goassimp/app/models"
	"goassimp/app/routes"

	"goassimp/app/lib/mgnredis"
)

type Sand struct {
	App
}

func (c *Sand) List() revel.Result {
	greeting := "Super Mother Fucing Hage!!"
	unko := "unko"
	return c.RenderJson(map[string]interface{}{
		"Count":    4,
		"unko":     unko,
		"greeting": greeting,
		"Status":   "Successfully uploaded",
	})
}

func (c *Sand) Show(id int) revel.Result {
	oem := models.Oembed{Version: 1, Type: "fuga"}
	return c.RenderJson(oem)
}

func (c *Sand) Cancel(id int) revel.Result {
	c.Flash.Success(fmt.Sprintln("Booking cancelled for confirmation number", id))
	return c.Redirect(routes.Sand.Index())
}

func (c *Sand) Index() revel.Result {
	greeting := "Super Mother Fucing Hage!!"
	info, err := mgnredis.RedisDb.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", info)
	return c.Render(greeting, "unko")
}
