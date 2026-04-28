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

func TestFetchCategoryPageWithType_SendsPTQuery(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "3", r.URL.Query().Get("pt"))
		assert.Equal(t, "24", r.URL.Query().Get("limit"))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"GeoInfo": map[string]interface{}{
				"ShortName":     "USA",
				"PropertyCount": float64(10),
				"Breadcrumbs":   []interface{}{},
			},
			"Result": map[string]interface{}{
				"Items":    []interface{}{},
				"Sections": []interface{}{},
			},
		})
	}))
	defer fakeServer.Close()

	data, err := FetchCategoryPageWithType(fakeServer.URL, "usa", map[string]string{"pt": "3"})

	assert.NoError(t, err)
	assert.Equal(t, "USA", data.LocationName)
}

func TestFillItemsFromSections_FillsUpToTarget(t *testing.T) {
	baseItems := []map[string]interface{}{
		{"ID": "A"},
		{"ID": "B"},
	}

	sections := []interface{}{
		map[string]interface{}{
			"ProcessedItems": []map[string]interface{}{
				{"ID": "C"},
				{"ID": "D"},
				{"ID": "E"},
			},
		},
	}

	result := fillItemsFromSections(baseItems, sections, 4)
	assert.Equal(t, 4, len(result))
	assert.Equal(t, "A", result[0]["ID"])
	assert.Equal(t, "B", result[1]["ID"])
}

func TestFillItemsFromSections_DeduplicatesByID(t *testing.T) {
	baseItems := []map[string]interface{}{
		{"ID": "A"},
	}

	sections := []interface{}{
		map[string]interface{}{
			"ProcessedItems": []map[string]interface{}{
				{"ID": "A"},
				{"ID": "B"},
				{"ID": "C"},
			},
		},
	}

	result := fillItemsFromSections(baseItems, sections, 3)
	assert.Equal(t, 3, len(result))
	assert.Equal(t, "A", result[0]["ID"])
	assert.Equal(t, "B", result[1]["ID"])
	assert.Equal(t, "C", result[2]["ID"])
}
