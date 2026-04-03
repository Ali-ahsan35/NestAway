package controllers

import (
	"smartours/requests"

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
	amount := c.GetString("amount")
	selectedCurrency := c.GetString("selectedCurrency")
	guests := c.GetString("pax")
	dateStart := c.GetString("dateStart")
	dateEnd := c.GetString("dateEnd")
	if order == "" {
		order = "1" // default: Most Popular
	}

	baseURL, _ := beego.AppConfig.String("api_base_url")

	result, err := requests.FetchProperties(baseURL,requests.PropertyParams{
		Category:         category,
		Order:            order,
		Amenities:        amenities,
		EcoFriendly:      ecoFriendly,
		Amount:           amount,
		SelectedCurrency: selectedCurrency,
		Guests:           guests,
		DateStart:        dateStart,
		DateEnd:          dateEnd,
	})
	if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

	c.Data["json"] = result
	c.ServeJSON()
}
