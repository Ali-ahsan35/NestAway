package controllers

import (
	"fmt"
	// "smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type RefineController struct {
	beego.Controller
}

func (c *RefineController) Get() {
	keyword := c.GetString("search", "Barcelona, Spain")
	fmt.Println("Search keyword:", keyword)
	c.Data["Keyword"] = keyword
	c.TplName = "refine.tpl"
}

// func (c *RefineController) GetBreadcrumb() {
// 	keyword := c.GetString("keyword", "Barcelona, Spain")

// 	c.Data["Keyword"] = keyword
// 	baseURL, _ := beego.AppConfig.String("api_base_url")
// 	result,err:=requests.FetchBreadcrumb(baseURL,keyword)
// 	if err != nil {
//         c.Data["json"] = map[string]string{"error": err.Error()}
//         c.ServeJSON()
//         return
//     }
// 	c.Data["json"] = result
// 	c.ServeJSON()
// }
