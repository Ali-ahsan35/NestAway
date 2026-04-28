package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CategoryData struct {
	LocationName  string
	PropertyCount string
	Breadcrumbs   []interface{}
	Items         []map[string]interface{}
	Sections      []interface{}
}

func FetchCategoryPage(baseURL string, slug string) (CategoryData, error) {
	apiURL := baseURL + "/api/v1/category/details/" + slug
	return fetchCategoryPageByURL(apiURL)
}

func FetchCategoryPageWithType(baseURL string, slug string, params map[string]string) (CategoryData, error) {
	apiURL, err := url.Parse(baseURL + "/api/v1/category/details/" + slug)
	if err != nil {
		return CategoryData{}, err
	}

	query := apiURL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	apiURL.RawQuery = query.Encode()

	return fetchCategoryPageByURL(apiURL.String())
}

func fillItemsFromSections(items []map[string]interface{}, sections []interface{}, target int) []map[string]interface{} {
	if len(items) >= target || len(sections) == 0 {
		return items
	}

	seen := map[string]struct{}{}
	for _, item := range items {
		if id, ok := item["ID"].(string); ok && id != "" {
			seen[id] = struct{}{}
		}
	}

	merged := append([]map[string]interface{}{}, items...)
	for _, sec := range sections {
		sectionMap, ok := sec.(map[string]interface{})
		if !ok {
			continue
		}

		processedItems, ok := sectionMap["ProcessedItems"].([]map[string]interface{})
		if !ok {
			continue
		}

		for _, item := range processedItems {
			if len(merged) >= target {
				return merged
			}

			id, _ := item["ID"].(string)
			if id != "" {
				if _, exists := seen[id]; exists {
					continue
				}
				seen[id] = struct{}{}
			}

			merged = append(merged, item)
		}
	}

	return merged
}

func fetchCategoryPageByURL(apiURL string) (CategoryData, error) {

	fmt.Println("Calling our API:", apiURL)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return CategoryData{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error calling our API:", err)
		return CategoryData{}, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Our API status:", resp.StatusCode)

	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	// extract the data
	geoInfo, _ := result["GeoInfo"].(map[string]interface{})
	propertyCount := ""
	locationName := ""
	breadcrumbs := []interface{}{}

	if geoInfo != nil {
		if count, ok := geoInfo["PropertyCount"].(float64); ok {
			propertyCount = fmt.Sprintf("%.0f+", count)
		}
		if name, ok := geoInfo["ShortName"].(string); ok {
			locationName = name
		}
		if bc, ok := geoInfo["Breadcrumbs"].([]interface{}); ok {
			breadcrumbs = bc
		}
	}

	items := []interface{}{}
	if result, ok := result["Result"].(map[string]interface{}); ok {
		if itemsList, ok := result["Items"].([]interface{}); ok {
			items = itemsList
		}
	}

	// extract first 3 amenities
	processedItems := []map[string]interface{}{}
	for i, item := range items {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		itemMap["Index"] = i
		prop, _ := itemMap["Property"].(map[string]interface{})
		if prop != nil {
			amenitiesMap, _ := prop["Amenities"].(map[string]interface{})
			top3 := []string{}
			for _, v := range amenitiesMap {
				if len(top3) >= 3 {
					break
				}
				top3 = append(top3, v.(string))
			}
			prop["TopAmenities"] = top3
			itemMap["Property"] = prop
		}
		processedItems = append(processedItems, itemMap)
	}

	sections := []interface{}{}
	if resultMap, ok := result["Result"].(map[string]interface{}); ok {
		if sectionsList, ok := resultMap["Sections"].([]interface{}); ok {
			// Process each section's items same as main items
			for i, section := range sectionsList {
				sectionMap, ok := section.(map[string]interface{})
				if !ok {
					continue
				}
				sectionItems, _ := sectionMap["Items"].([]interface{})
				processedSectionItems := []map[string]interface{}{}
				for _, item := range sectionItems {
					itemMap, ok := item.(map[string]interface{})
					if !ok {
						continue
					}
					itemMap["Index"] = i
					prop, _ := itemMap["Property"].(map[string]interface{})
					if prop != nil {
						amenitiesMap, _ := prop["Amenities"].(map[string]interface{})
						top3 := []string{}
						for _, v := range amenitiesMap {
							if len(top3) >= 3 {
								break
							}
							top3 = append(top3, v.(string))
						}
						prop["TopAmenities"] = top3
						itemMap["Property"] = prop
					}
					processedSectionItems = append(processedSectionItems, itemMap)
				}
				sectionMap["ProcessedItems"] = processedSectionItems

				title, _ := sectionMap["Title"].(string)
				subTitle, _ := sectionMap["SubTitle"].(string)
				sectionMap["Title"] = strings.ReplaceAll(title, "{{.Location}}", locationName)
				sectionMap["SubTitle"] = strings.ReplaceAll(subTitle, "{{.Location}}", locationName)

				sections = append(sections, sectionMap)
			}
		}
	}
	processedItems = fillItemsFromSections(processedItems, sections, 24)

	return CategoryData{
		LocationName:  locationName,
		PropertyCount: propertyCount,
		Breadcrumbs:   breadcrumbs,
		Items:         processedItems,
		Sections:      sections,
	}, nil
}
