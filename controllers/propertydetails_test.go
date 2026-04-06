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

func TestPropertyDetailsController_Success(t *testing.T) {
	mockFetcher := mocks.NewPropertyDetailsFetcher(t)

	fakeData := map[string]interface{}{
		"Items": []interface{}{
			map[string]interface{}{
				"ID": "EP-123",
				"Property": map[string]interface{}{
					"PropertyName": "Test Villa",
				},
			},
		},
	}

	mockFetcher.On("FetchPropertyDetails", "", "EP-123").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/propertydetails?ids=EP-123", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertyDetailsController{
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
	items := result["Items"].([]interface{})
	assert.Len(t, items, 1)
}

func TestPropertyDetailsController_Error(t *testing.T) {
	mockFetcher := mocks.NewPropertyDetailsFetcher(t)

	mockFetcher.On("FetchPropertyDetails", "", "EP-123").Return(nil, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/api/propertydetails?ids=EP-123", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertyDetailsController{
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

func TestPropertyDetailsController_MultipleIDs(t *testing.T) {
	mockFetcher := mocks.NewPropertyDetailsFetcher(t)

	fakeData := map[string]interface{}{
		"Items": []interface{}{
			map[string]interface{}{
				"ID": "EP-123",
				"Property": map[string]interface{}{
					"PropertyName": "Test Villa 1",
				},
			},
			map[string]interface{}{
				"ID": "EP-456",
				"Property": map[string]interface{}{
					"PropertyName": "Test Villa 2",
				},
			},
			map[string]interface{}{
				"ID": "EP-789",
				"Property": map[string]interface{}{
					"PropertyName": "Test Villa 3",
				},
			},
		},
	}

	mockFetcher.On("FetchPropertyDetails", "", "EP-123,EP-456,EP-789").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/propertydetails?ids=EP-123,EP-456,EP-789", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &PropertyDetailsController{
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
	items := result["Items"].([]interface{})
	assert.Len(t, items, 3)
}