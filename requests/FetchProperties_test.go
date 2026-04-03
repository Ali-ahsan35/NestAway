package requests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchProperties_Success(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.String(), "/api/properties/category/v1")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Result": map[string]interface{}{
				"ItemIDs": []string{"EP-123", "EP-456"},
			},
		})
	}))
	defer fakeServer.Close()

	result, err := FetchProperties(fakeServer.URL, PropertyParams{
		Category: "spain",
		Order:    "1",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	resultMap := result["Result"].(map[string]interface{})
	assert.NotNil(t, resultMap["ItemIDs"])
}

func TestFetchProperties_APIError(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer fakeServer.Close()

	result, err := FetchProperties(fakeServer.URL, PropertyParams{
		Category: "spain",
		Order:    "1",
	})

	// 500 with empty body → result is nil, no Go error
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestFetchProperties_WithFilters(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check filters are in the URL
		assert.Contains(t, r.URL.String(), "amenities=1-2-3")
		assert.Contains(t, r.URL.String(), "pax=2")
		assert.Contains(t, r.URL.String(), "dateStart=2026-04-01")
		assert.Contains(t, r.URL.String(), "dateEnd=2026-04-05")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Result": map[string]interface{}{
				"ItemIDs": []string{"EP-123"},
			},
		})
	}))
	defer fakeServer.Close()

	result, err := FetchProperties(fakeServer.URL, PropertyParams{
		Category:  "spain",
		Order:     "1",
		Amenities: "1-2-3",
		Guests:    "2",
		DateStart: "2026-04-01",
		DateEnd:   "2026-04-05",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}