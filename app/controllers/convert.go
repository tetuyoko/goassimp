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
	var convert_logs []models.ConvertLog
	if err := mugendb.Db.Order("created_at desc").Limit(10).Find(&convert_logs).Error; err != nil {
		panic(err)
	}
	return c.Render(convert_logs)
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
	con := models.ConvertLog{UUID: get8UUID(), Url: pth}
	status := "Status: Successfully uploaded"

	if err := mugendb.Db.Save(&con).Error; err != nil {
		errs := fmt.Sprintf("%s", err)
		status = "failed err" + errs
	}

	//return c.RenderJson(map[string]interface{}{
	//	"Status": status,
	//})
	fmt.Println(status)

	//buf := bytes.NewBuffer([]byte(`{
	//	"test": {
	//		"array": [1, "2", 3],
	//		"arraywithsubs": [
	//			{"subkeyone": 1},
	//			{"subkeytwo": 2, "subkeythree": 3}
	//		],
	//		"bignum": 8000000000
	//	}
	//}`))
	//js, err := NewFromReader(buf)

	return c.RenderJson(con)
}

func get8UUID() string {
	u1 := uuid.NewV4()
	str := fmt.Sprintf("%s", u1)
	return str[0:8]
}
