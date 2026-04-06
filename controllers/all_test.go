package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"smartours/mocks"
	"smartours/requests"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestAllController_Success(t *testing.T) {
	mockFetcher := mocks.NewCategoryPageFetcher(t)

	fakeData := requests.CategoryData{
		LocationName:  "Barcelona",
		PropertyCount: "172+",
		Breadcrumbs:   []interface{}{},
		Items:         []map[string]interface{}{},
		Sections:      []interface{}{},
	}

	mockFetcher.On("FetchCategoryPage", "", "spain/catalonia/barcelona").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/all/spain/catalonia/barcelona", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "spain/catalonia/barcelona")

	controller := &AllController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "all.tpl", controller.TplName)
	assert.Equal(t, "Barcelona", controller.Data["LocationName"])
	assert.Equal(t, "172+", controller.Data["PropertyCount"])
}

func TestAllController_Error(t *testing.T) {
	mockFetcher := mocks.NewCategoryPageFetcher(t)

	mockFetcher.On("FetchCategoryPage", "", "spain/catalonia/barcelona").Return(requests.CategoryData{}, assert.AnError)

	req := httptest.NewRequest(http.MethodGet, "/all/spain/catalonia/barcelona", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "spain/catalonia/barcelona")

	controller := &AllController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "all.tpl", controller.TplName)
	assert.Equal(t, assert.AnError.Error(), controller.Data["Error"])
}

func TestAllController_EmptySlug(t *testing.T) {
	mockFetcher := mocks.NewCategoryPageFetcher(t)

	fakeData := requests.CategoryData{
		LocationName:  "",
		PropertyCount: "",
		Breadcrumbs:   []interface{}{},
		Items:         []map[string]interface{}{},
		Sections:      []interface{}{},
	}

	mockFetcher.On("FetchCategoryPage", "", "").Return(fakeData, nil)

	req := httptest.NewRequest(http.MethodGet, "/all/", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "")

	controller := &AllController{
		Fetcher: mockFetcher,
	}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "all.tpl", controller.TplName)
}