package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type CategoryDetailsController struct {
	beego.Controller
}

func (c *CategoryDetailsController) Get() {
	rawSlug := c.Ctx.Input.Param(":splat")
	slug := strings.ToLower(strings.ReplaceAll(rawSlug, "/", ":"))

	if slug == "" {
		c.Data["json"] = map[string]string{"error": "country is required"}
		c.ServeJSON()
		return
	}

	apiURL := "https://presto:TRAV3LA1@ownerdirect.beta.123presto.com/api/v1/category/details/" + slug +
		"?aggsAvgPrice=1" +
		"&aggsAvgRating=1" +
		"&aggsAvgRoomSize=1" +
		"&aggsCategory=1" +
		"&device=desktop" +
		"&items=1" +
		"&locations=US" +
		"&sections=1"

	fmt.Println("CategoryDetails URL:", apiURL)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Request create error:", err)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	// Set required headers
	req.Header.Set("Origin", "123presto-MS-ROW.com")
	req.Header.Set("User-Agent", "desktop")
	req.Header.Set("X-Api-key", "sparrowxkey@w3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(bodyBytes[:min(len(bodyBytes), 500)]))

	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	c.Data["json"] = result
	c.ServeJSON()
}
