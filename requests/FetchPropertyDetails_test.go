package requests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchPropertyDetails_Success(t *testing.T)  {
	fakeServer :=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		assert.Contains(t, r.URL.String(),"/api/property/bookmark/v1")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Items": []interface{}{
                map[string]interface{}{
                    "ID": "EP-123",
                    "Property": map[string]interface{}{
                        "PropertyName": "Test Villa",
                    },
                },
            },
		})
	}))
	defer fakeServer.Close()

	result, err := FetchPropertyDetails(fakeServer.URL, "EP-123")

    assert.NoError(t, err)
    assert.NotNil(t, result)
    items := result["Items"].([]interface{})
    assert.Len(t, items, 1)
    item := items[0].(map[string]interface{})
    assert.Equal(t, "EP-123", item["ID"])
}

func TestFetchPropertyDetails_APIError (t *testing.T)  {
	fakeServer := httptest.NewServer(http.HandlerFunc(func( w http.ResponseWriter, r *http.Request)  {
		w.WriteHeader(500)
	}))
	defer fakeServer.Close()

	result, err:=FetchPropertyDetails(fakeServer.URL, "EP-123")
	assert.NoError(t, err)
    assert.Nil(t, result)
}

func TestFetchPropertyDetails_MultipleIDs(t *testing.T) {
    fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        assert.Contains(t, r.URL.String(), "propertyIdList=")
        assert.Contains(t, r.URL.String(), "EP-123")
        assert.Contains(t, r.URL.String(), "EP-456")
        assert.Contains(t, r.URL.String(), "EP-789")

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
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
        })
    }))
    defer fakeServer.Close()

    result, err := FetchPropertyDetails(fakeServer.URL, "EP-123,EP-456,EP-789")

    assert.NoError(t, err)
    assert.NotNil(t, result)
    items := result["Items"].([]interface{})
    assert.Len(t, items, 3)
}