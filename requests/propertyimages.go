package requests

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func FetchPropertyImages(propertyId string) (map[string]interface{}, error) {
	apiURL := "https://presto:TRAV3LA1@ownerdirect.beta.123presto.com/api/property/images/v1?propertyId="+url.QueryEscape(propertyId)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/json")

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