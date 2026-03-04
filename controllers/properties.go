package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertiesController struct {
	beego.Controller
}

func (c *PropertiesController) Get() {
    category := c.GetString("category")
    order := c.GetString("order")
    if order == "" {
        order = "1" // default: Most Popular
    }

    apiURL := "https://smartours.com/api/properties/category/v1?order=" + order +
        "&category=" + url.QueryEscape(category) +
        "&limit=192&items=1&locations=BD&device=desktop&page=1"

    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }
    req.Header.Set("X-Requested-With", "XMLHttpRequest")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)

    c.Data["json"] = result
    c.ServeJSON()
}