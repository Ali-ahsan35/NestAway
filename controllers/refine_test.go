package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestRefineController_Get(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/refine?search=Barcelona", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &RefineController{}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "refine.tpl", controller.TplName)
	assert.Equal(t, "Barcelona", controller.Data["Keyword"])
}

func TestRefineController_DefaultKeyword(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/refine", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &RefineController{}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	assert.Equal(t, "refine.tpl", controller.TplName)
	assert.Equal(t, "Barcelona, Spain", controller.Data["Keyword"])
}