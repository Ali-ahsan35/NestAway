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
	amenities := c.GetString("amenities")
	ecoFriendly := c.GetString("ecoFriendly")
	// minPrice := c.GetString("minPrice")
	// maxPrice := c.GetString("maxPrice")
	amount := c.GetString("amount")
	selectedCurrency := c.GetString("selectedCurrency")
	guests := c.GetString("pax")
	dateStart := c.GetString("dateStart")
	dateEnd := c.GetString("dateEnd")
	if order == "" {
		order = "1" // default: Most Popular
	}

	apiURL := "https://presto:TRAV3LA1@ownerdirect.beta.123presto.com/api/properties/category/v1?order=" + order +
		"&category=" + url.QueryEscape(category) +
		"&limit=192&items=1&locations=BD&device=desktop&page=1"

	if amenities != "" {
		apiURL += "&amenities=" + amenities
	}
	if ecoFriendly == "true" {
		apiURL += "&ecoFriendly=true"
	}
	if amount != "" {
		apiURL += "&amount=" + amount
		if selectedCurrency != "" {
			apiURL += "&selectedCurrency=" + selectedCurrency
		} else {
			apiURL += "&selectedCurrency=BDT"
		}
	}
	if guests != "" {
		apiURL += "&pax=" + guests
	}
	if dateStart != "" {
		apiURL += "&dateStart=" + dateStart
	}
	if dateEnd != "" {
		apiURL += "&dateEnd=" + dateEnd
	}

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
