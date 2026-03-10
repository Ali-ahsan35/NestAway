package routers

import (
	"smartours/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/refine", &controllers.RefineController{})
	beego.Router("/api/breadcrumb", &controllers.RefineController{}, "get:GetBreadcrumb")
	beego.Router("/api/properties", &controllers.PropertiesController{}, "get:Get")
	beego.Router("/api/propertydetails", &controllers.PropertyDetailsController{}, "get:Get")
	beego.Router("/api/v1/category/details/?:country", &controllers.CategoryDetailsController{}, "get:Get")
	beego.Router("/all/:country", &controllers.AllController{}, "get:Get")
	beego.Router("/redirect-partner", &controllers.RedirectPartnerController{}, "get:Get")
}
