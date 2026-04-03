package controllers

import (
	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
	beego.Controller
}

func (c *PropertyDetailsController) Get() {
	ids := c.GetString("ids")
	baseURL, _ := beego.AppConfig.String("api_base_url")
	result,err := requests.FetchPropertyDetails(baseURL,ids)
	if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

	c.Data["json"] = result
	c.ServeJSON()
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
