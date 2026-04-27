package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"smartours/mocks"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestCategoryDetailsController_Success(t *testing.T) {
	mockFetcher := mocks.NewCategoryDetailsFetcher(t)

	fakeData := map[string]interface{}{
		"GeoInfo": map[string]interface{}{
			"ShortName":     "Barcelona",
			"PropertyCount": float64(172),
		},
	}

	mockFetcher.On("FetchCategoryDetails", "", "spain:catalonia:barcelona", "", "").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/details/spain/catalonia/barcelona", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "spain/catalonia/barcelona")

	controller := &CategoryDetailsController{
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
	geoInfo := result["GeoInfo"].(map[string]interface{})
	assert.Equal(t, "Barcelona", geoInfo["ShortName"])
}

func TestCategoryDetailsController_Error(t *testing.T) {
	mockFetcher := mocks.NewCategoryDetailsFetcher(t)

	mockFetcher.On("FetchCategoryDetails", "", "spain:catalonia:barcelona", "", "").Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/details/spain/catalonia/barcelona", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "spain/catalonia/barcelona")

	controller := &CategoryDetailsController{
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

func TestCategoryDetailsController_EmptySlug(t *testing.T) {
	mockFetcher := mocks.NewCategoryDetailsFetcher(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/details/", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "")

	controller := &CategoryDetailsController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "country is required", result["error"])
}

func TestCategoryDetailsController_WithPT(t *testing.T) {
	mockFetcher := mocks.NewCategoryDetailsFetcher(t)

	fakeData := map[string]interface{}{
		"GeoInfo": map[string]interface{}{
			"ShortName": "USA",
		},
	}

	mockFetcher.On("FetchCategoryDetails", "", "usa", "3", "").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/details/usa?pt=3", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "usa")

	controller := &CategoryDetailsController{
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
	geoInfo := result["GeoInfo"].(map[string]interface{})
	assert.Equal(t, "USA", geoInfo["ShortName"])
}

func TestCategoryDetailsController_WithPTAndLimit(t *testing.T) {
	mockFetcher := mocks.NewCategoryDetailsFetcher(t)

	fakeData := map[string]interface{}{
		"GeoInfo": map[string]interface{}{
			"ShortName": "USA",
		},
	}

	mockFetcher.On("FetchCategoryDetails", "", "usa", "3", "24").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/category/details/usa?pt=3&limit=24", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "usa")

	controller := &CategoryDetailsController{
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
	geoInfo := result["GeoInfo"].(map[string]interface{})
	assert.Equal(t, "USA", geoInfo["ShortName"])
}
