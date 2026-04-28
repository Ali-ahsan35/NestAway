package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"smartours/requests"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

type fakeCategoryPageByTypeFetcher struct {
	data      requests.CategoryData
	err       error
	gotBase   string
	gotSlug   string
	gotParams map[string]string
	callCount int
}

func (f *fakeCategoryPageByTypeFetcher) FetchCategoryPageWithType(baseURL, slug string, params map[string]string) (requests.CategoryData, error) {
	f.callCount++
	f.gotBase = baseURL
	f.gotSlug = slug
	f.gotParams = params
	return f.data, f.err
}

func writeTypeMapFile(t *testing.T, content string) string {
	t.Helper()
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "property_types.json")
	err := os.WriteFile(filePath, []byte(content), 0o644)
	assert.NoError(t, err)
	return filePath
}

func TestAllByPropertyTypeController_Success(t *testing.T) {
	oldPath := propertyTypesFilePath
	propertyTypesFilePath = writeTypeMapFile(t, `{"cabins":{"pt":"3"},"resorts":{"pt":"7"}}`)
	defer func() { propertyTypesFilePath = oldPath }()

	fetcher := &fakeCategoryPageByTypeFetcher{
		data: requests.CategoryData{
			LocationName:  "USA",
			PropertyCount: "10+",
			Breadcrumbs:   []interface{}{},
			Items:         []map[string]interface{}{},
			Sections:      []interface{}{},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/all/usa/cabins", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "usa")
	ctx.Input.SetParam(":propertyType", "cabins")

	controller := &AllByPropertyTypeController{Fetcher: fetcher}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "all_by_type.tpl", controller.TplName)
	assert.Equal(t, "USA", controller.Data["LocationName"])
	assert.Equal(t, "cabins", controller.Data["PropertyType"])
	assert.Equal(t, "usa", fetcher.gotSlug)
	assert.Equal(t, map[string]string{"pt": "3"}, fetcher.gotParams)
	assert.Equal(t, 1, fetcher.callCount)
}

func TestAllByPropertyTypeController_UnsupportedType(t *testing.T) {
	oldPath := propertyTypesFilePath
	propertyTypesFilePath = writeTypeMapFile(t, `{"cabins":{"pt":"3"}}`)
	defer func() { propertyTypesFilePath = oldPath }()

	fetcher := &fakeCategoryPageByTypeFetcher{}

	req := httptest.NewRequest(http.MethodGet, "/all/usa/hotels", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "usa")
	ctx.Input.SetParam(":propertyType", "hotels")

	controller := &AllByPropertyTypeController{Fetcher: fetcher}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "all_by_type.tpl", controller.TplName)
	assert.Equal(t, "unsupported property type: hotels", controller.Data["Error"])
	assert.Equal(t, 0, fetcher.callCount)
}

func TestAllByPropertyTypeController_MissingSlug(t *testing.T) {
	fetcher := &fakeCategoryPageByTypeFetcher{}

	req := httptest.NewRequest(http.MethodGet, "/all//cabins", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)
	ctx.Input.SetParam(":splat", "")
	ctx.Input.SetParam(":propertyType", "cabins")

	controller := &AllByPropertyTypeController{Fetcher: fetcher}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "all_by_type.tpl", controller.TplName)
	assert.Equal(t, "location slug is required", controller.Data["Error"])
	assert.Equal(t, 0, fetcher.callCount)
}
