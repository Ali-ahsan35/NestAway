package routers

import (
	"encoding/json"
	"os"
	"sort"

	"smartours/controllers"
	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

const subcategoryMappingPath = "static/data/property_types.json"

func loadSubcategoryRoutes() []string {
	body, err := os.ReadFile(subcategoryMappingPath)
	if err != nil {
		return []string{}
	}

	var mapping map[string]map[string]string
	if err := json.Unmarshal(body, &mapping); err != nil {
		return []string{}
	}

	routes := make([]string, 0, len(mapping))
	for key := range mapping {
		routes = append(routes, key)
	}
	sort.Strings(routes)
	return routes
}

func init() {
	// beego.Router("/", &controllers.MainController{})
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
	beego.Router("/api/v1/category/details/*", &controllers.CategoryDetailsController{
		Fetcher: &requests.CategoryDetailsFetcherImpl{},
	}, "get:Get")
	beego.Router("/redirect-partner", &controllers.RedirectPartnerController{}, "get:Get")
	beego.Router("/api/property/images", &controllers.PropertyImagesController{
		Fetcher: &requests.PropertyImagesFetcherImpl{},
	}, "get:Get")

	// Sub-category routes — defined dynamically
	subCategories := loadSubcategoryRoutes()

	allByTypeController := &controllers.AllByPropertyTypeController{
		Fetcher: &requests.CategoryPageByTypeFetcherImpl{},
	}

	for _, sc := range subCategories {
		beego.Router("/all/*/"+sc, allByTypeController, "get:Get")
	}

	// General /all/* route must come AFTER sub-category routes
	beego.Router("/all/*", &controllers.AllController{
		Fetcher: &requests.CategoryPageFetcherImpl{},
	}, "get:Get")
}
