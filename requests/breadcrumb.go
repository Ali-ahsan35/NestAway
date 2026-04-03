package requests

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func FetchBreadcrumb(baseURL string,keyword string)(map[string]interface{}, error) {
	apiURL := baseURL+"/api/location/v1?keyword=" +
		url.QueryEscape(keyword) + "&isLocationEntity=true"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil,err
	}

	return result, nil
}