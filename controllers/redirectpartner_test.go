package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestRedirectPartnerController_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet,
		"/redirect-partner?direct_redirect_url=https://www.expedia.com&property_id=EP-123&feed=24",
		nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &RedirectPartnerController{}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	resp := w.Result()
	assert.Equal(t, http.StatusFound, resp.StatusCode) // 302
	assert.Equal(t, "https://www.expedia.com", resp.Header.Get("Location"))
}

func TestRedirectPartnerController_EmptyURL(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/redirect-partner", nil)
	w := httptest.NewRecorder()

	ctx := context.NewContext()
	ctx.Reset(w, req)

	controller := &RedirectPartnerController{}
	controller.Ctx = ctx
	controller.Data = make(map[interface{}]interface{})

	controller.Get()

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "No redirect URL provided", w.Body.String())
}