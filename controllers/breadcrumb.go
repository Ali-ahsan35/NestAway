package controllers

import (
	"smartours/requests"
	beego "github.com/beego/beego/v2/server/web"
)

type BreadcrumbController struct {
	beego.Controller
	Fetcher requests.BreadcrumbFetcher
}

func (c *BreadcrumbController) GetBreadcrumb() {
	keyword := c.GetString("keyword")
	baseURL, _ := beego.AppConfig.String("api_base_url")

	result, err := c.Fetcher.FetchBreadcrumb(baseURL, keyword)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}