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

func TestFetchCategoryPage_WithItems(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"GeoInfo": map[string]interface{}{
				"ShortName":     "Barcelona",
				"PropertyCount": float64(172),
				"Breadcrumbs":   []interface{}{},
			},
			"Result": map[string]interface{}{
				"Items": []interface{}{
					map[string]interface{}{
						"ID": "EP-123",
						"Property": map[string]interface{}{
							"PropertyName": "Test Villa",
							"Amenities": map[string]interface{}{
								"1": "Pool",
								"2": "WiFi",
								"3": "Kitchen",
								"4": "Parking",
							},
						},
					},
				},
				"Sections": []interface{}{},
			},
		})
	}))
	defer fakeServer.Close()

	data, err := FetchCategoryPage(fakeServer.URL, "spain/catalonia/barcelona")

	assert.NoError(t, err)
	assert.Equal(t, 1, len(data.Items))

	item := data.Items[0]
	prop := item["Property"].(map[string]interface{})
	topAmenities := prop["TopAmenities"].([]string)
	assert.Len(t, topAmenities, 3)

	assert.Equal(t, 0, item["Index"])
}

func TestFetchCategoryPage_WithSections(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"GeoInfo": map[string]interface{}{
				"ShortName":     "Barcelona",
				"PropertyCount": float64(50),
				"Breadcrumbs":   []interface{}{},
			},
			"Result": map[string]interface{}{
				"Items": []interface{}{},
				"Sections": []interface{}{
					map[string]interface{}{
						"Title":    "Top Villas in {{.Location}}",
						"SubTitle": "Best villas in {{.Location}}",
						"Count":    float64(10),
						"Items": []interface{}{
							map[string]interface{}{
								"ID": "EP-456",
								"Property": map[string]interface{}{
									"PropertyName": "Section Villa",
									"Amenities": map[string]interface{}{
										"1": "Pool",
										"2": "WiFi",
									},
								},
							},
						},
					},
				},
			},
		})
	}))
	defer fakeServer.Close()

	data, err := FetchCategoryPage(fakeServer.URL, "spain/catalonia/barcelona")

	assert.NoError(t, err)
	assert.Equal(t, 1, len(data.Sections))

	section := data.Sections[0].(map[string]interface{})
	assert.Equal(t, "Top Villas in Barcelona", section["Title"])
	assert.Equal(t, "Best villas in Barcelona", section["SubTitle"])

	sectionItems := section["ProcessedItems"].([]map[string]interface{})
	assert.Len(t, sectionItems, 1)
}