package testsapi

import (
	"encoding/json"
	"final/routers"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSocialMediaSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	socialMedia := `{
    "name": "instagram",
    "social_media_url": "ini adalah contoh url"
	}`
	bodyRequest := strings.NewReader(socialMedia)

	token := "Bearer " + tokenUser
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/socialmedias/", bodyRequest)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)
}

func TestCreateSocialMediaFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	testCases := []struct {
		desc    string
		request string
		token   string
	}{
		{
			desc: "token invalid",
			request: `{
				"name": "instagram",
				"social_media_url": "ini adalah contoh url"
			}`,
			token: "token salah",
		},
		{
			desc: "name required",
			request: `{
				"social_media_url": "ini adalah contoh url"
			}`,
			token: token,
		},
		{
			desc: "socialmedia url required",
			request: `{
				"name": "instagram",
			}`,
			token: token,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bodyRequest := strings.NewReader(tC.request)
			request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/socialmedias/", bodyRequest)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("Authorization", tC.token)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			response := recorder.Result()
			assert.Equal(t, 400, response.StatusCode)
		})
	}
}

func TestGetSocialMediaSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	socialMedia := `{
    "name": "instagram",
    "social_media_url": "ini adalah contoh url"
	}`
	bodyRequest := strings.NewReader(socialMedia)

	token := "Bearer " + tokenUser

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/socialmedias/", bodyRequest)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	requestGet := httptest.NewRequest(http.MethodGet, "http://localhost:3000/socialmedias/", nil)
	requestGet.Header.Add("Authorization", token)

	recorderGet := httptest.NewRecorder()
	router.ServeHTTP(recorderGet, requestGet)
	responseGet := recorderGet.Result()
	assert.Equal(t, 200, responseGet.StatusCode)
}

func TestGetSocialMediaFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	requestGet := httptest.NewRequest(http.MethodGet, "http://localhost:3000/socialmedias/", nil)
	requestGet.Header.Add("Authorization", token)

	recorderGet := httptest.NewRecorder()
	router.ServeHTTP(recorderGet, requestGet)
	responseGet := recorderGet.Result()
	assert.Equal(t, 400, responseGet.StatusCode)
}

func TestUpdateSocialMediaSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	socialMedia := `{
    "name": "instagram",
    "social_media_url": "ini adalah contoh url"
	}`
	bodyRequest := strings.NewReader(socialMedia)

	token := "Bearer " + tokenUser

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/socialmedias/", bodyRequest)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	responseAdd := recorder.Result()

	body, _ := io.ReadAll(responseAdd.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)
	socialMediaId := responseBody["id"]

	socialMediaId = int(socialMediaId.(float64))
	urlSocialMedia := fmt.Sprintf("http://localhost:3000/socialmedias/%d", socialMediaId)

	socialMediaUpdate := `{
    "name": "instagram update",
    "social_media_url": "ini adalah contoh url update"
	}`
	bodyRequestUpdate := strings.NewReader(socialMediaUpdate)

	requestUpdate := httptest.NewRequest(http.MethodPut, urlSocialMedia, bodyRequestUpdate)
	requestUpdate.Header.Add("Authorization", token)

	recorderUpdate := httptest.NewRecorder()
	router.ServeHTTP(recorderUpdate, requestUpdate)
	responseUpdate := recorderUpdate.Result()
	assert.Equal(t, 200, responseUpdate.StatusCode)
}

func TestUpdateSocialMediaFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	socialMediaId := 0
	urlSocialMedia := fmt.Sprintf("http://localhost:3000/socialmedias/%d", socialMediaId)

	socialMediaUpdate := `{
    "name": "instagram update",
    "social_media_url": "ini adalah contoh url update"
	}`
	bodyRequestUpdate := strings.NewReader(socialMediaUpdate)

	requestUpdate := httptest.NewRequest(http.MethodPut, urlSocialMedia, bodyRequestUpdate)
	requestUpdate.Header.Add("Authorization", token)

	recorderUpdate := httptest.NewRecorder()
	router.ServeHTTP(recorderUpdate, requestUpdate)

	responseUpdate := recorderUpdate.Result()
	assert.Equal(t, 400, responseUpdate.StatusCode)
}

func TestDeleteSocialMediaSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	socialMedia := `{
    "name": "instagram",
    "social_media_url": "ini adalah contoh url"
	}`
	bodyRequest := strings.NewReader(socialMedia)

	token := "Bearer " + tokenUser

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/socialmedias/", bodyRequest)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	responseAdd := recorder.Result()

	body, _ := io.ReadAll(responseAdd.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)
	socialMediaId := responseBody["id"]

	socialMediaId = int(socialMediaId.(float64))
	urlSocialMedia := fmt.Sprintf("http://localhost:3000/socialmedias/%d", socialMediaId)

	requestDelete := httptest.NewRequest(http.MethodDelete, urlSocialMedia, nil)
	requestDelete.Header.Add("Authorization", token)

	recorderDelete := httptest.NewRecorder()
	router.ServeHTTP(recorderDelete, requestDelete)
	responseDelete := recorderDelete.Result()
	assert.Equal(t, 200, responseDelete.StatusCode)
}

func TestDeleteSocialMediaFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	socialMediaId := 0
	urlSocialMedia := fmt.Sprintf("http://localhost:3000/socialmedias/%d", socialMediaId)

	requestDelete := httptest.NewRequest(http.MethodDelete, urlSocialMedia, nil)
	requestDelete.Header.Add("Authorization", token)

	recorderDelete := httptest.NewRecorder()
	router.ServeHTTP(recorderDelete, requestDelete)
	responseDelete := recorderDelete.Result()
	assert.Equal(t, 400, responseDelete.StatusCode)
}
