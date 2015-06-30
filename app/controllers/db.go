package controllers

import (
	//    "fmt"
	"github.com/revel/revel"
	//    "goassimp/app/models"
)

type DB struct {
	*revel.Controller
}

func (c DB) Index() revel.Result {
	//    for i := 0; i < 10; i++ {
	//        DbMap.Insert(&models.User{0, fmt.Sprintf("user%d", i)})
	//    }
	//
	//    rows, _ := DbMap.Select(models.User{}, "select * from user")
	//    for _, row := range rows {
	//        user := row.(*models.User)
	//        fmt.Printf("%d, %s\n", user.Id, user.Name)
	//    }

	return c.Render()
}
