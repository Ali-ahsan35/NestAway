package controllers

import (
	"encoding/json"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
	beego.Controller
}

func (c *PropertyDetailsController) Get() {
	ids := c.GetString("ids") // comma-separated IDs

	apiURL := "https://smartours.com/api/property/bookmark/v1?propertyIdList=" + ids

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