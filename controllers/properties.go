package controllers

import (
	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertiesController struct {
	beego.Controller
	Fetcher requests.PropertiesFetcher
}

func (c *PropertiesController) Get() {
	category := c.GetString("category")
	order := c.GetString("order")
	amenities := c.GetString("amenities")
	ecoFriendly := c.GetString("ecoFriendly")
	amount := c.GetString("amount")
	selectedCurrency := c.GetString("selectedCurrency")
	guests := c.GetString("pax")
	dateStart := c.GetString("dateStart")
	dateEnd := c.GetString("dateEnd")
	extraParams := map[string]string{}
	for key, values := range c.Ctx.Request.URL.Query() {
		if len(values) == 0 {
			continue
		}
		extraParams[key] = values[0]
	}
	delete(extraParams, "category")
	delete(extraParams, "order")
	delete(extraParams, "amenities")
	delete(extraParams, "ecoFriendly")
	delete(extraParams, "amount")
	delete(extraParams, "selectedCurrency")
	delete(extraParams, "pax")
	delete(extraParams, "dateStart")
	delete(extraParams, "dateEnd")
	if len(extraParams) == 0 {
		extraParams = nil
	}
	if order == "" {
		order = "1"
	}

	baseURL, _ := beego.AppConfig.String("api_base_url")
	result, err := c.Fetcher.FetchProperties(baseURL, requests.PropertyParams{
		Category:         category,
		Order:            order,
		Amenities:        amenities,
		EcoFriendly:      ecoFriendly,
		Amount:           amount,
		SelectedCurrency: selectedCurrency,
		Guests:           guests,
		DateStart:        dateStart,
		DateEnd:          dateEnd,
		ExtraParams:      extraParams,
	})
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = result
	c.ServeJSON()
}
