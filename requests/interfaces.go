package requests

type BreadcrumbFetcher interface {
	FetchBreadcrumb(baseURL, keyword string) (map[string]interface{}, error)
}

type PropertiesFetcher interface {
	FetchProperties(baseURL string, params PropertyParams) (map[string]interface{}, error)
}

type PropertyDetailsFetcher interface {
	FetchPropertyDetails(baseURL, ids string) (map[string]interface{}, error)
}

type PropertyImagesFetcher interface {
	FetchPropertyImages(baseURL, propertyId string) (map[string]interface{}, error)
}

type CategoryDetailsFetcher interface {
	FetchCategoryDetails(baseURL, slug, pt, limit string) (map[string]interface{}, error)
}

type CategoryPageFetcher interface {
	FetchCategoryPage(baseURL, slug string) (CategoryData, error)
}

type CategoryPageByTypeFetcher interface {
	FetchCategoryPageWithType(baseURL, slug string, pt int) (CategoryData, error)
}
