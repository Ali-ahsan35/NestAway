package controllers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type AllByPropertyTypeController struct {
	beego.Controller
	Fetcher requests.CategoryPageByTypeFetcher
}

var propertyTypesFilePath = filepath.Join("static", "data", "property_types.json")

func loadSubcategoryMapping(path string) (map[string]map[string]string, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := map[string]map[string]string{}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	normalized := make(map[string]map[string]string, len(raw))
	for key, value := range raw {
		normalized[strings.ToLower(strings.TrimSpace(key))] = value
	}

	return normalized, nil
}

func (c *AllByPropertyTypeController) Get() {
	path := strings.Trim(c.Ctx.Request.URL.Path, "/")
	segments := strings.Split(path, "/")
	if len(segments) < 3 || strings.ToLower(segments[0]) != "all" {
		c.Data["Error"] = "invalid sub-category route"
		c.TplName = "all_by_type.tpl"
		return
	}

	rawSlug := strings.Join(segments[1:len(segments)-1], "/")
	propertyType := strings.ToLower(strings.TrimSpace(segments[len(segments)-1]))

	if rawSlug == "" {
		c.Data["Error"] = "location slug is required"
		c.TplName = "all_by_type.tpl"
		return
	}

	if propertyType == "" {
		c.Data["Error"] = "property type is required"
		c.Data["Country"] = rawSlug
		c.TplName = "all_by_type.tpl"
		return
	}

	mapping, err := loadSubcategoryMapping(propertyTypesFilePath)
	if err != nil {
		c.Data["Error"] = "failed to load subcategory mappings"
		c.Data["Country"] = rawSlug
		c.TplName = "all_by_type.tpl"
		return
	}

	params, ok := mapping[propertyType]
	if !ok {
		c.Data["Error"] = fmt.Sprintf("unsupported property type: %s", propertyType)
		c.Data["Country"] = rawSlug
		c.Data["PropertyType"] = propertyType
		c.TplName = "all_by_type.tpl"
		return
	}

	localURL, _ := beego.AppConfig.String("local_base_url")

	data, err := c.Fetcher.FetchCategoryPageWithType(localURL, rawSlug, params)
	if err != nil {
		c.Data["Error"] = err.Error()
		c.Data["Country"] = rawSlug
		c.Data["PropertyType"] = propertyType
		c.TplName = "all_by_type.tpl"
		return
	}

	// Build refine URL for "View More Properties"
	refineURL := buildRefineURL(rawSlug, data.LocationFullName, params)

	displayType := titleCase(strings.ReplaceAll(propertyType, "-", " "))

	c.Data["Items"] = data.Items
	c.Data["Sections"] = data.Sections
	c.Data["Country"] = rawSlug
	c.Data["PropertyType"] = displayType
	c.Data["LocationName"] = data.LocationName
	c.Data["PropertyCount"] = data.PropertyCount
	c.Data["Breadcrumbs"] = data.Breadcrumbs
	c.Data["RefineURL"] = refineURL
	c.TplName = "all_by_type.tpl"
}

func buildRefineURL(slug string, locationFullName string, params map[string]string) string {
	location := strings.TrimSpace(locationFullName)
	if location == "" {
		// Convert slug to location name for search param
		// usa/texas -> usa texas
		location = strings.ReplaceAll(slug, "/", " ")
	}
	query := url.Values{}
	query.Set("search", location)
	for key, value := range params {
		query.Set(key, value)
	}
	return "/refine?" + query.Encode()
}

func titleCase(input string) string {
	parts := strings.Fields(strings.ToLower(input))
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}
