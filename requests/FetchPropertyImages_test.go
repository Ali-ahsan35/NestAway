package requests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFetchPropertyImages_Success(t *testing.T) {
    fakeJSON := `{
        "Error": null,
        "Images": [
            "img1.jpg",
            "img2.jpg"
        ],
        "Message": "",
        "Success": true
    }`

    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "/api/property/images/v1", r.URL.Path)

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(fakeJSON))
    }))
    defer server.Close()

    result, err := FetchPropertyImages(server.URL, "EP-120752254")

    assert.NoError(t, err)
    assert.NotNil(t, result)

    images := result["Images"].([]interface{})
    assert.Len(t, images, 2)
    assert.Equal(t, "img1.jpg", images[0])
    assert.True(t, result["Success"].(bool))
}

func TestFetchPropertyImages_APIError(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
    }))
    defer server.Close()

    result, err := FetchPropertyImages(server.URL, "EP-120752254")

    assert.Error(t, err)
    assert.Nil(t, result)
}
func TestFetchPropertyImages_EmptyImages(t *testing.T) {
    fakeJSON := `{
        "Error": null,
        "Images": [],
        "Message": "",
        "Success": true
    }`

    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(fakeJSON))
    }))
    defer server.Close()

    result, err := FetchPropertyImages(server.URL, "EP-120752254")

    assert.NoError(t, err)

    images := result["Images"].([]interface{})
    assert.Len(t, images, 0)
}