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

func TestCreatePhotoSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	photo := `{
    "title": "photo indah",
    "caption": "ini adalah photo indah",
    "photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
	}`
	bodyPhoto := strings.NewReader(photo)

	token := "Bearer " + tokenUser
	requestPhoto := httptest.NewRequest("POST", "http://localhost:3000/photos/", bodyPhoto)
	requestPhoto.Header.Add("Content-Type", "application/json")
	requestPhoto.Header.Add("Authorization", token)

	recorderPhoto := httptest.NewRecorder()

	router.ServeHTTP(recorderPhoto, requestPhoto)

	responsePhoto := recorderPhoto.Result()
	assert.Equal(t, 201, responsePhoto.StatusCode)
}

func TestCreatePhotoFailed(t *testing.T) {
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
				"title": "photo indah",
				"caption": "ini adalah photo indah",
				"photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
			}`,
			token: "salah",
		},
		{
			desc: "title required",
			request: `{
				"caption": "ini adalah photo indah",
				"photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
			}`,
			token: token,
		},
		{
			desc: "photo url required",
			request: `{
				"title": "photo indah",
				"caption": "ini adalah photo indah",
			}`,
			token: token,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bodyPhoto := strings.NewReader(tC.request)

			requestPhoto := httptest.NewRequest("POST", "http://localhost:3000/photos/", bodyPhoto)
			requestPhoto.Header.Add("Content-Type", "application/json")
			requestPhoto.Header.Add("Authorization", tC.token)

			recorderPhoto := httptest.NewRecorder()
			router.ServeHTTP(recorderPhoto, requestPhoto)

			responsePhoto := recorderPhoto.Result()
			assert.Equal(t, 400, responsePhoto.StatusCode)
		})
	}
}

func TestGetPhotoSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	photo := `{
    "title": "photo indah",
    "caption": "ini adalah photo indah",
    "photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
	}`
	bodyPhoto := strings.NewReader(photo)

	token := "Bearer " + tokenUser
	requestPhoto := httptest.NewRequest(http.MethodPost, "http://localhost:3000/photos/", bodyPhoto)
	requestPhoto.Header.Add("Content-Type", "application/json")
	requestPhoto.Header.Add("Authorization", token)
	recorderPhoto := httptest.NewRecorder()
	router.ServeHTTP(recorderPhoto, requestPhoto)

	requestGetPhoto := httptest.NewRequest(http.MethodGet, "http://localhost:3000/photos/", nil)
	requestGetPhoto.Header.Add("Authorization", token)
	recorderGetPhoto := httptest.NewRecorder()
	router.ServeHTTP(recorderGetPhoto, requestGetPhoto)
	responseGetPhoto := recorderGetPhoto.Result()
	assert.Equal(t, 200, responseGetPhoto.StatusCode)
}

func TestGetPhotoFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	requestGetPhoto := httptest.NewRequest(http.MethodGet, "http://localhost:3000/photos/", nil)
	requestGetPhoto.Header.Add("Authorization", token)
	recorderGetPhoto := httptest.NewRecorder()

	router.ServeHTTP(recorderGetPhoto, requestGetPhoto)

	responseGetPhoto := recorderGetPhoto.Result()
	assert.Equal(t, 400, responseGetPhoto.StatusCode)
}

func TestUpdatePhotoSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	photo := `{
    "title": "photo indah",
    "caption": "ini adalah photo indah",
    "photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
	}`
	bodyPhoto := strings.NewReader(photo)

	token := "Bearer " + tokenUser
	requestPhoto := httptest.NewRequest(http.MethodPost, "http://localhost:3000/photos/", bodyPhoto)
	requestPhoto.Header.Add("Content-Type", "application/json")
	requestPhoto.Header.Add("Authorization", token)
	recorderPhoto := httptest.NewRecorder()

	router.ServeHTTP(recorderPhoto, requestPhoto)
	responseAddPhoto := recorderPhoto.Result()

	body, _ := io.ReadAll(responseAddPhoto.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	photoId := responseBody["id"]
	photoId = int(photoId.(float64))
	urlPhoto := fmt.Sprintf("http://localhost:3000/photos/%d", photoId)

	photoUpdate := `{
    "title": "photo update",
    "caption": "ini adalah photo update",
    "photo_url": "url photo update"
	}`
	bodyUpdatePhoto := strings.NewReader(photoUpdate)

	requestUpdatePhoto := httptest.NewRequest(http.MethodPut, urlPhoto, bodyUpdatePhoto)
	requestUpdatePhoto.Header.Add("Authorization", token)

	recorderUpdatePhoto := httptest.NewRecorder()
	router.ServeHTTP(recorderUpdatePhoto, requestUpdatePhoto)

	responseUpdatePhoto := recorderUpdatePhoto.Result()
	assert.Equal(t, 200, responseUpdatePhoto.StatusCode)
}

func TestUpdatePhotoFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	urlPhoto := "http://localhost:3000/photos/1"

	photoUpdate := `{
    "title": "photo update",
    "caption": "ini adalah photo update",
    "photo_url": "url photo update"
	}`
	bodyUpdatePhoto := strings.NewReader(photoUpdate)

	// photo not found
	requestUpdatePhoto := httptest.NewRequest(http.MethodPut, urlPhoto, bodyUpdatePhoto)
	requestUpdatePhoto.Header.Add("Authorization", token)

	recorderUpdatePhoto := httptest.NewRecorder()
	router.ServeHTTP(recorderUpdatePhoto, requestUpdatePhoto)

	responseUpdatePhoto := recorderUpdatePhoto.Result()
	assert.Equal(t, 400, responseUpdatePhoto.StatusCode)
}

func TestDeletePhotoSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)

	photo := `{
    "title": "photo indah",
    "caption": "ini adalah photo indah",
    "photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
	}`
	bodyPhoto := strings.NewReader(photo)

	token := "Bearer " + tokenUser
	requestPhoto := httptest.NewRequest(http.MethodPost, "http://localhost:3000/photos/", bodyPhoto)
	requestPhoto.Header.Add("Content-Type", "application/json")
	requestPhoto.Header.Add("Authorization", token)
	recorderPhoto := httptest.NewRecorder()

	router.ServeHTTP(recorderPhoto, requestPhoto)
	responseAddPhoto := recorderPhoto.Result()

	body, _ := io.ReadAll(responseAddPhoto.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	photoId := responseBody["id"]
	photoId = int(photoId.(float64))
	urlPhoto := fmt.Sprintf("http://localhost:3000/photos/%d", photoId)

	requestUpdatePhoto := httptest.NewRequest(http.MethodDelete, urlPhoto, nil)
	requestUpdatePhoto.Header.Add("Authorization", token)

	recorderUpdatePhoto := httptest.NewRecorder()
	router.ServeHTTP(recorderUpdatePhoto, requestUpdatePhoto)

	responseUpdatePhoto := recorderUpdatePhoto.Result()
	assert.Equal(t, 200, responseUpdatePhoto.StatusCode)
}

func TestDeletePhotoFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	token := "Bearer " + tokenUser

	urlPhoto := "http://localhost:3000/photos/1"

	requestUpdatePhoto := httptest.NewRequest(http.MethodDelete, urlPhoto, nil)
	requestUpdatePhoto.Header.Add("Authorization", token)

	recorderUpdatePhoto := httptest.NewRecorder()
	router.ServeHTTP(recorderUpdatePhoto, requestUpdatePhoto)

	responseUpdatePhoto := recorderUpdatePhoto.Result()
	assert.Equal(t, 400, responseUpdatePhoto.StatusCode)
}
