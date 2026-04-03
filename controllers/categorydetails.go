package controllers

import (
	"smartours/requests"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type CategoryDetailsController struct {
	beego.Controller
}

func (c *CategoryDetailsController) Get() {
	rawSlug := c.Ctx.Input.Param(":splat")
	slug := strings.ToLower(strings.ReplaceAll(rawSlug, "/", ":"))

	if slug == "" {
		c.Data["json"] = map[string]string{"error": "country is required"}
		c.ServeJSON()
		return
	}
	baseURL, _ := beego.AppConfig.String("api_base_url")
	result, err := requests.FetchCategoryDetails(baseURL,slug)
	if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

	c.Data["json"] = result
	c.ServeJSON()
}
