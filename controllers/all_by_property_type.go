package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"smartours/requests"

	beego "github.com/beego/beego/v2/server/web"
)

type AllByPropertyTypeController struct {
	beego.Controller
	Fetcher requests.CategoryPageByTypeFetcher
}

// propertyTypesFilePath is kept as a variable to make testing easier.
var propertyTypesFilePath = filepath.Join("static", "data", "property_types.json")

func loadPropertyTypeMap(path string) (map[string]int, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw := map[string]int{}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	normalized := make(map[string]int, len(raw))
	for key, value := range raw {
		normalized[strings.ToLower(strings.TrimSpace(key))] = value
	}

	return normalized, nil
}

func (c *AllByPropertyTypeController) Get() {
	rawSlug := strings.Trim(c.Ctx.Input.Param(":splat"), "/")
	propertyType := strings.ToLower(strings.TrimSpace(c.Ctx.Input.Param(":propertyType")))

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

	typeMap, err := loadPropertyTypeMap(propertyTypesFilePath)
	if err != nil {
		c.Data["Error"] = "failed to load property type mappings"
		c.Data["Country"] = rawSlug
		c.TplName = "all_by_type.tpl"
		return
	}

	pt, ok := typeMap[propertyType]
	if !ok {
		c.Data["Error"] = fmt.Sprintf("unsupported property type: %s", propertyType)
		c.Data["Country"] = rawSlug
		c.Data["PropertyType"] = propertyType
		c.TplName = "all_by_type.tpl"
		return
	}

	localURL, _ := beego.AppConfig.String("local_base_url")

	data, err := c.Fetcher.FetchCategoryPageWithType(localURL, rawSlug, pt)
	if err != nil {
		c.Data["Error"] = err.Error()
		c.Data["Country"] = rawSlug
		c.Data["PropertyType"] = propertyType
		c.Data["PropertyTypeID"] = strconv.Itoa(pt)
		c.TplName = "all_by_type.tpl"
		return
	}

	c.Data["Items"] = data.Items
	c.Data["Country"] = rawSlug
	c.Data["PropertyType"] = propertyType
	c.Data["PropertyTypeID"] = strconv.Itoa(pt)
	c.Data["LocationName"] = data.LocationName
	c.Data["PropertyCount"] = data.PropertyCount
	c.Data["Breadcrumbs"] = data.Breadcrumbs
	c.TplName = "all_by_type.tpl"
}
