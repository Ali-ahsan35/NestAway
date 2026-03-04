package controllers

import (
	"encoding/json"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertiesController struct {
	beego.Controller
}

func (c *PropertiesController) Get() {
	category := c.GetString("category")

	apiURL := "https://smartours.com/api/properties/category/v1" +
		"?order=1" +
		"&category=" + category +
		"&limit=192" +
		"&items=1" +
		"&locations=BD" +
		"&device=desktop" +
		"&page=1"

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	c.Data["json"] = result
	c.ServeJSON()
}