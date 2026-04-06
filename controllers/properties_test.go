package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"smartours/mocks"
	"smartours/requests"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestPropertiesController_Success(t *testing.T) {
	mockFetcher := mocks.NewPropertiesFetcher(t)

	fakeData := map[string]interface{}{
		"Result": map[string]interface{}{
			"ItemIDs": []interface{}{"EP-123", "EP-456"},
		},
	}

	mockFetcher.On("FetchProperties", "", requests.PropertyParams{
		Category: "spain",
		Order:    "1",
	}).Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/properties?category=spain&order=1", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertiesController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.NotNil(t, result)
	resultMap := result["Result"].(map[string]interface{})
	assert.NotNil(t, resultMap["ItemIDs"])
}

func TestPropertiesController_Error(t *testing.T) {
	mockFetcher := mocks.NewPropertiesFetcher(t)

	mockFetcher.On("FetchProperties", "", requests.PropertyParams{
		Category: "spain",
		Order:    "1",
	}).Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/api/properties?category=spain&order=1", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertiesController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, assert.AnError.Error(), result["error"])
}

func TestPropertiesController_WithFilters(t *testing.T) {
	mockFetcher := mocks.NewPropertiesFetcher(t)

	fakeData := map[string]interface{}{
		"Result": map[string]interface{}{
			"ItemIDs": []interface{}{"EP-123"},
		},
	}

	mockFetcher.On("FetchProperties", "", requests.PropertyParams{
		Category:  "spain",
		Order:     "1",
		Amenities: "1-2-3",
		Guests:    "2",
		DateStart: "2026-04-01",
		DateEnd:   "2026-04-05",
	}).Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet,
		"/api/properties?category=spain&order=1&amenities=1-2-3&pax=2&dateStart=2026-04-01&dateEnd=2026-04-05",
		nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertiesController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.NotNil(t, result)
}

func init() {
	beego.AppConfig.Set("api_base_url", "")
}