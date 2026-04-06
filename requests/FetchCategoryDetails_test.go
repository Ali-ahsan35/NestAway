package requests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCategoryDetails_Success(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), "/api/v1/category/details/")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"GeoInfo": map[string]interface{}{
				"ShortName":     "Barcelona",
				"PropertyCount": 172,
			},
		})
	}))
	defer fakeServer.Close()

	result, err := FetchCategoryDetails(fakeServer.URL, "spain:catalonia:barcelona")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	geoInfo := result["GeoInfo"].(map[string]interface{})
	assert.Equal(t, "Barcelona", geoInfo["ShortName"])
}

func TestFetchCategoryDetails_APIError(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer fakeServer.Close()

	result, err := FetchCategoryDetails(fakeServer.URL, "spain:catalonia:barcelona")

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestFetchCategoryDetails_EmptySlug(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{})
	}))
	defer fakeServer.Close()

	result, err := FetchCategoryDetails(fakeServer.URL, "")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}