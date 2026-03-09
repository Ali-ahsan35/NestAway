package controllers

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "strings"

    beego "github.com/beego/beego/v2/server/web"
)

type AllController struct {
    beego.Controller
}

func (c *AllController) Get() {
    country := strings.ToLower(c.Ctx.Input.Param(":country"))

    apiURL := "http://localhost:8080/api/v1/category/details/" + country

    fmt.Println("Calling our API:", apiURL)

    req, err := http.NewRequest("GET", apiURL, nil)
    if err != nil {
        c.Data["Error"] = err.Error()
        c.Data["Country"] = country
        c.TplName = "all.tpl"
        return
    }

	// req.Header.Set("X-Requested-With", "XMLHttpRequest")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error calling our API:", err)
        c.Data["Error"] = err.Error()
        c.Data["Country"] = country
        c.TplName = "all.tpl"
        return
    }
    defer resp.Body.Close()

    bodyBytes, _ := io.ReadAll(resp.Body)
    fmt.Println("Our API status:", resp.StatusCode)

    var result map[string]interface{}
    json.Unmarshal(bodyBytes, &result)

    // After unmarshaling result, extract the data
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

	c.Data["Country"] = country
	c.Data["PropertyCount"] = propertyCount
	c.Data["LocationName"] = locationName
	c.Data["Breadcrumbs"] = breadcrumbs
	c.Data["LocationName"] = locationName
	c.Data["CategoryData"] = result
	c.TplName = "all.tpl"
}