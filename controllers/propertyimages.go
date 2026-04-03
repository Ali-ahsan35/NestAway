package controllers

import (
	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyImagesController struct{
	beego.Controller
}

func (c* PropertyImagesController) Get() {
	propertyId := c.GetString("propertyId")

	baseURL, _ := beego.AppConfig.String("api_base_url")

	result,err:=requests.FetchPropertyImages(baseURL,propertyId)
	if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

	c.Data["json"] = result
	c.ServeJSON()
}