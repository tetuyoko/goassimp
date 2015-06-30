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
		errs := fmt.Sprintf("%s", err)
		return c.RenderJson(map[string]interface{}{
			"Status": errs,
		})
	}
	return c.Render(convert_logs)
}

func (c *Convert) Convert(userfile []byte) revel.Result {
	var status string
	uuid := get8UUID()

	// 画像保存
	dstDir := "public/tmp/" + string(uuid)
	if c.Params.Files["userfile"] == nil {
		return c.RenderJson(map[string]interface{}{
			"Status": "not found",
		})
	}
	dstName := c.Params.Files["userfile"][0].Filename

	// create dir
	if err := os.Mkdir(dstDir, 0777); err != nil {
		errs := fmt.Sprintf("%s", err)
		status = "failed err" + errs
		return c.RenderJson(map[string]interface{}{
			"Status": status,
		})
		fmt.Println(err)
	}

	// create file
	pth := dstDir + "/" + dstName

	dst, err := os.Create(pth)
	defer dst.Close()
	if err != nil {
		errs := fmt.Sprintf("%s", err)
		status = "failed err" + errs
		return c.RenderJson(map[string]interface{}{
			"Status": status,
		})
	}

	_, err = io.Copy(dst, bytes.NewReader(userfile))
	if err != nil {
		errs := fmt.Sprintf("%s", err)
		status = "failed err" + errs
		return c.RenderJson(map[string]interface{}{
			"Status": status,
		})
	}

	// ログ保存
	con := models.ConvertLog{UUID: uuid , Url: pth}
	status = "Status: Successfully uploaded"

	if err := mugendb.Db.Save(&con).Error; err != nil {
		errs := fmt.Sprintf("%s", err)
		status = "failed err" + errs
		return c.RenderJson(map[string]interface{}{
			"Status": status,
		})
	}

	return c.RenderJson(con)
}

func get8UUID() string {
	u1 := uuid.NewV4()
	str := fmt.Sprintf("%s", u1)
	return str[0:8]
}

//return c.RenderJson(map[string]interface{}{
//	"Status": status,
//})
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

