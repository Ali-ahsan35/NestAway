package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
)

type RefineController struct {
	web.Controller
}

func (c *RefineController) Get() {
	c.TplName = "refine.tpl"
}


func (c *RefineController) GetBreadcrumb() {
	keyword := c.GetString("keyword")

	apiURL := "https://smartours.com/api/location/v1?keyword=" +
		url.QueryEscape(keyword) + "&isLocationEntity=true"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	// Add headers
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}