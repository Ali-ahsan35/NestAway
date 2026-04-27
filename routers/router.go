package routers

import (
	"smartours/controllers"
	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/refine", &controllers.RefineController{})
	beego.Router("/api/breadcrumb", &controllers.BreadcrumbController{
		Fetcher: &requests.BreadcrumbFetcherImpl{},
	}, "get:GetBreadcrumb")
	beego.Router("/api/properties", &controllers.PropertiesController{
		Fetcher: &requests.PropertiesFetcherImpl{},
	}, "get:Get")
	beego.Router("/api/propertydetails", &controllers.PropertyDetailsController{
		Fetcher: &requests.PropertyDetailsFetcherImpl{},
	}, "get:Get")
	beego.Router("/all/*/:propertyType", &controllers.AllByPropertyTypeController{
		Fetcher: &requests.CategoryPageByTypeFetcherImpl{},
	}, "get:Get")
	beego.Router("/all/*", &controllers.AllController{
		Fetcher: &requests.CategoryPageFetcherImpl{},
	}, "get:Get")
	beego.Router("/api/v1/category/details/*", &controllers.CategoryDetailsController{
		Fetcher: &requests.CategoryDetailsFetcherImpl{},
	}, "get:Get")
	beego.Router("/redirect-partner", &controllers.RedirectPartnerController{}, "get:Get")
	beego.Router("/api/property/images", &controllers.PropertyImagesController{
		Fetcher: &requests.PropertyImagesFetcherImpl{},
	}, "get:Get")
}
