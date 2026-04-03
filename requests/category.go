package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func FetchCategoryDetails(baseURL string,slug string) (map[string]interface{}, error) {
    apiURL := baseURL+"/api/v1/category/details/" + slug +
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
		return nil, err
	}

	// Set required headers
	req.Header.Set("Origin", "123presto-MS-ROW.com")
	req.Header.Set("User-Agent", "desktop")
	req.Header.Set("X-Api-key", "sparrowxkey@w3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Response:", string(bodyBytes[:min(len(bodyBytes), 500)]))

	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	return result, nil
}