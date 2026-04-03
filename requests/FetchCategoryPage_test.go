package requests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCategoryPage_Success(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"GeoInfo": map[string]interface{}{
				"ShortName":     "Barcelona",
				"PropertyCount": float64(172),
				"Breadcrumbs": []interface{}{
					map[string]interface{}{
						"Name": "Spain",
						"Slug": "spain",
					},
					map[string]interface{}{
						"Name": "Catalonia",
						"Slug": "spain/catalonia",
					},
				},
			},
			"Result": map[string]interface{}{
				"Items":    []interface{}{},
				"Sections": []interface{}{},
			},
		})
	}))
	defer fakeServer.Close()

	data, err := FetchCategoryPage(fakeServer.URL, "spain/catalonia/barcelona")

	assert.NoError(t, err)
	assert.Equal(t, "Barcelona", data.LocationName)
	assert.Equal(t, "172+", data.PropertyCount)
	assert.Equal(t, 2, len(data.Breadcrumbs))
}

func TestFetchCategoryPage_APIError(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer fakeServer.Close()

	data, err := FetchCategoryPage(fakeServer.URL, "spain/catalonia/barcelona")

	assert.NoError(t, err)
	assert.Equal(t, "", data.LocationName)
	assert.Equal(t, "", data.PropertyCount)
	assert.Equal(t, 0, len(data.Breadcrumbs))
}

func TestFetchCategoryPage_EmptyGeoInfo(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"GeoInfo": map[string]interface{}{},
			"Result": map[string]interface{}{
				"Items":    []interface{}{},
				"Sections": []interface{}{},
			},
		})
	}))
	defer fakeServer.Close()

	data, err := FetchCategoryPage(fakeServer.URL, "spain/catalonia/barcelona")

	assert.NoError(t, err)
	assert.Equal(t, "", data.LocationName)
	assert.Equal(t, "", data.PropertyCount)
	assert.Equal(t, 0, len(data.Breadcrumbs))
	assert.Equal(t, 0, len(data.Items))
}