package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"smartours/mocks"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestBreadcrumbController_Success(t *testing.T) {
	// 1. Create mock
	mockFetcher := mocks.NewBreadcrumbFetcher(t)

	// 2. Tell mock what to return
	fakeData := map[string]interface{}{
		"GeoInfo": map[string]interface{}{
			"ShortName": "Barcelona",
		},
	}
	mockFetcher.On("FetchBreadcrumb", "", "Barcelona").Return(fakeData, nil)

	// 3. Create fake HTTP request and response
	req := httptest.NewRequest(http.MethodGet, "/api/breadcrumb?keyword=Barcelona", nil)
	w := httptest.NewRecorder()

	// 4. Setup Beego context
	ctx := context.NewContext()
	ctx.Reset(w, req)

	// 5. Create controller with mock and context
	controller := &BreadcrumbController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	// 6. Call controller method
	controller.GetBreadcrumb()

	// 7. Assert response
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.NotNil(t, result)

	geoInfo := result["GeoInfo"].(map[string]interface{})
	assert.Equal(t, "Barcelona", geoInfo["ShortName"])
}

func TestBreadcrumbController_Error(t *testing.T) {
	mockFetcher := mocks.NewBreadcrumbFetcher(t)

	mockFetcher.On("FetchBreadcrumb", "", "Barcelona").Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/api/breadcrumb?keyword=Barcelona", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &BreadcrumbController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.GetBreadcrumb()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, assert.AnError.Error(), result["error"])
}

func TestBreadcrumbController_EmptyKeyword(t *testing.T) {
	mockFetcher := mocks.NewBreadcrumbFetcher(t)

	fakeData := map[string]interface{}{
		"GeoInfo": map[string]interface{}{},
	}
	mockFetcher.On("FetchBreadcrumb", "", "").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/breadcrumb", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &BreadcrumbController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.GetBreadcrumb()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func init() {
	 beego.AppConfig.Set("api_base_url", "")
}