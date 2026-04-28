package requests

type BreadcrumbFetcherImpl struct{}

func (b *BreadcrumbFetcherImpl) FetchBreadcrumb(baseURL, keyword string) (map[string]interface{}, error) {
	return FetchBreadcrumb(baseURL, keyword)
}

type PropertiesFetcherImpl struct{}

func (p *PropertiesFetcherImpl) FetchProperties(baseURL string, params PropertyParams) (map[string]interface{}, error) {
	return FetchProperties(baseURL, params)
}

type PropertyImagesFetcherImpl struct{}

func (p *PropertyImagesFetcherImpl) FetchPropertyImages(baseURL string, propertyId string) (map[string]interface{}, error) {
	return FetchPropertyImages(baseURL, propertyId)
}

type PropertyDetailsFetcherImpl struct{}

func (p *PropertyDetailsFetcherImpl) FetchPropertyDetails(baseURL string, ids string) (map[string]interface{}, error) {
	return FetchPropertyDetails(baseURL, ids)
}

type CategoryDetailsFetcherImpl struct{}

func (p *CategoryDetailsFetcherImpl) FetchCategoryDetails(baseURL string, slug string, pt string, limit string) (map[string]interface{}, error) {
	return FetchCategoryDetails(baseURL, slug, pt, limit)
}

type CategoryPageFetcherImpl struct{}

func (c *CategoryPageFetcherImpl) FetchCategoryPage(baseURL, slug string) (CategoryData, error) {
	return FetchCategoryPage(baseURL, slug)
}

// type CategoryPageByTypeFetcherImpl struct{}

// func (c *CategoryPageByTypeFetcherImpl) FetchCategoryPageWithType(baseURL, slug string, pt int) (CategoryData, error) {
// 	return FetchCategoryPageWithType(baseURL, slug, pt)
// }

type CategoryPageByTypeFetcherImpl struct{}

func (c *CategoryPageByTypeFetcherImpl) FetchCategoryPageWithType(baseURL, slug string, params map[string]string) (CategoryData, error) {
    return FetchCategoryPageWithType(baseURL, slug, params)
}
