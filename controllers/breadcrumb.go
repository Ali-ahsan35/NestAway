package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BreadcrumbController struct {
	beego.Controller
}

func (c *BreadcrumbController) Get() {
	keyword := c.GetString("keyword")
	encodedKeyword := url.QueryEscape(keyword)

	apiURL := "https://smartours.com/api/location/v1?keyword=" + encodedKeyword + "&isLocationEntity=true"

	resp, err := http.Get(apiURL)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid JSON"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}