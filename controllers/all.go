package controllers

import (
	"smartours/requests"
	beego "github.com/beego/beego/v2/server/web"
)

type AllController struct {
    beego.Controller
}

func (c *AllController) Get() {
	rawSlug := c.Ctx.Input.Param(":splat")

	localURL, _ := beego.AppConfig.String("local_base_url")

	data, err := requests.FetchCategoryPage(localURL,rawSlug)
    if err != nil {
        c.Data["Error"] = err.Error()
        c.Data["Country"] = rawSlug
        c.TplName = "all.tpl"
        return
    }

	

	c.Data["Sections"] = data.Sections
	c.Data["Items"] = data.Items
	c.Data["Country"] = rawSlug
	c.Data["LocationName"] = data.LocationName
	c.Data["PropertyCount"] = data.PropertyCount
	c.Data["Breadcrumbs"] = data.Breadcrumbs
	c.TplName = "all.tpl"

}