package requests

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type PropertyParams struct {
	Category         string
	Order            string
	Amenities        string
	EcoFriendly      string
	Amount           string
	SelectedCurrency string
	Guests           string
	DateStart        string
	DateEnd          string
}

func FetchProperties(baseURL string, params PropertyParams) (map[string]interface{}, error) {

	apiURL := baseURL + "/api/properties/category/v1?order=" + params.Order + "&category=" + url.QueryEscape(params.Category) +
    "&limit=192&items=1&locations=BD&device=desktop&page=1"

	if params.Amenities != "" {
		apiURL += "&amenities=" + params.Amenities
	}
	if params.EcoFriendly == "true" {
		apiURL += "&ecoFriendly=true"
	}
	if params.Amount != "" {
		apiURL += "&amount=" + params.Amount
		if params.SelectedCurrency != "" {
			apiURL += "&selectedCurrency=" + params.SelectedCurrency
		} else {
			apiURL += "&selectedCurrency=BDT"
		}
	}
	if params.Guests != "" {
		apiURL += "&pax=" + params.Guests
	}
	if params.DateStart != "" {
		apiURL += "&dateStart=" + params.DateStart
	}
	if params.DateEnd != "" {
		apiURL += "&dateEnd=" + params.DateEnd
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
