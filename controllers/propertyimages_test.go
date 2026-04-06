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

func TestPropertyImagesController_Success(t *testing.T) {
	mockFetcher := mocks.NewPropertyImagesFetcher(t)

	fakeData := map[string]interface{}{
		"Images": []interface{}{
			"img1.jpg",
			"img2.jpg",
		},
		"Success": true,
	}

	mockFetcher.On("FetchPropertyImages", "", "EP-123").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/property/images?propertyId=EP-123", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertyImagesController{
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
	images := result["Images"].([]interface{})
	assert.Len(t, images, 2)
}

func TestPropertyImagesController_Error(t *testing.T) {
	mockFetcher := mocks.NewPropertyImagesFetcher(t)

	mockFetcher.On("FetchPropertyImages", "", "EP-123").Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/api/property/images?propertyId=EP-123", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertyImagesController{
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

func TestPropertyImagesController_EmptyPropertyId(t *testing.T) {
	mockFetcher := mocks.NewPropertyImagesFetcher(t)

	fakeData := map[string]interface{}{
		"Images":  []interface{}{},
		"Success": true,
	}

	mockFetcher.On("FetchPropertyImages", "", "").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/property/images", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertyImagesController{
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