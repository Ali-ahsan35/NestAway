package controllers

import (
	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type BreadcrumbController struct {
	beego.Controller
}

func (c *BreadcrumbController) Get() {
	keyword := c.GetString("keyword")

	baseURL, _ := beego.AppConfig.String("api_base_url")

	result,err:=requests.FetchBreadcrumb(baseURL,keyword)
	if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

	c.Data["json"] = result
	c.ServeJSON()
}
