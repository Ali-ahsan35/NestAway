package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchPropertyDetails(baseURL string,ids string) (map[string]interface{}, error) {
	apiURL := baseURL+"/api/property/bookmark/v1?propertyIdList=" + ids

	fmt.Println("PropertyDetails URL:", apiURL[:min(len(apiURL), 200)])

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil,err
	}
	defer resp.Body.Close()

	// Read body
	bodyBytes, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Read error:", err2)
	} else {
		fmt.Println("Response status:", resp.StatusCode)
		fmt.Println("Response body:", string(bodyBytes[:min(len(bodyBytes), 500)]))
	}

	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	return result,nil
}