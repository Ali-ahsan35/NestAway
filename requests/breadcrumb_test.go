package requests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFetchBreadcrumb_Success(t *testing.T) {
	fakeJSON := `{
		"GeoInfo": {
			"ShortName": "Barcelona",
			"LocationSlug": "barcelona-spain"
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/location/v1", r.URL.Path)
		assert.Equal(t, "Barcelona", r.URL.Query().Get("keyword"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakeJSON))
	}))
	defer server.Close()

	result, err := FetchBreadcrumb(server.URL, "Barcelona")

	assert.NoError(t, err)
	assert.NotNil(t, result)

	geoInfo := result["GeoInfo"].(map[string]interface{})
	assert.Equal(t, "Barcelona", geoInfo["ShortName"])
}

func TestFetchBreadcrumb_EmptyKeyword(t *testing.T) {
	fakeJSON := `{"GeoInfo": {}}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "", r.URL.Query().Get("keyword"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fakeJSON))
	}))
	defer server.Close()

	result, err := FetchBreadcrumb(server.URL, "")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestFetchBreadcrumb_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	result, err := FetchBreadcrumb(server.URL, "Barcelona")

	assert.Error(t, err)
	assert.Nil(t, result)
}