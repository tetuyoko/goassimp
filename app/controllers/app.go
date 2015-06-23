package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Unko() revel.Result {
	greeting := "Super Mother Fucing Hage!!"
	unko := "unko"
	return c.Render(greeting, unko)
}
