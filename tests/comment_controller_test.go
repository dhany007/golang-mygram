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

func TestCreateCommentSucces(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	photoId := TemporaryCreatePhoto(router, tokenUser)

	comment := fmt.Sprintf(`{
    "message": "salam kenal dari user 1",
    "photo_id": %d
	}`, photoId)
	bodyComment := strings.NewReader(comment)

	token := "Bearer " + tokenUser
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/comments/", bodyComment)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)
}

func TestCreateCommentFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	photoId := TemporaryCreatePhoto(router, tokenUser)
	token := "Bearer " + tokenUser

	testCases := []struct {
		desc    string
		request string
		token   string
	}{
		{
			desc: "photo id required",
			request: `{
				"message": "salam kenal dari user 1"
			}`,
			token: token,
		},
		{
			desc: "message required",
			request: fmt.Sprintf(`{
				"photo_id": %d
			}`, photoId),
			token: token,
		},
		{
			desc: "token invalid",
			request: fmt.Sprintf(`{
				"message": "salam kenal dari user 1",
				"photo_id": %d
			}`, photoId),
			token: "token salah",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			bodyComment := strings.NewReader(tC.request)

			request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/comments/", bodyComment)
			request.Header.Add("Content-Type", "application/json")
			request.Header.Add("Authorization", tC.token)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)

			response := recorder.Result()
			assert.Equal(t, 400, response.StatusCode)
		})
	}
}

func TestGetCommentSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	photoId := TemporaryCreatePhoto(router, tokenUser)

	comment := fmt.Sprintf(`{
    "message": "salam kenal dari user 1",
    "photo_id": %d
	}`, photoId)
	bodyComment := strings.NewReader(comment)

	token := "Bearer " + tokenUser
	requestAdd := httptest.NewRequest(http.MethodPost, "http://localhost:3000/comments/", bodyComment)
	requestAdd.Header.Add("Content-Type", "application/json")
	requestAdd.Header.Add("Authorization", token)

	recorderAdd := httptest.NewRecorder()
	router.ServeHTTP(recorderAdd, requestAdd)

	requestUpdate := httptest.NewRequest(http.MethodGet, "http://localhost:3000/comments/", nil)
	requestUpdate.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, requestUpdate)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestGetCommentFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	_ = TemporaryCreatePhoto(router, tokenUser)

	token := "Bearer " + tokenUser

	requestUpdate := httptest.NewRequest(http.MethodGet, "http://localhost:3000/comments/", nil)
	requestUpdate.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, requestUpdate)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateCommentSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	photoId := TemporaryCreatePhoto(router, tokenUser)

	comment := fmt.Sprintf(`{
    "message": "salam kenal dari user 1",
    "photo_id": %d
	}`, photoId)
	bodyComment := strings.NewReader(comment)

	token := "Bearer " + tokenUser
	requestAdd := httptest.NewRequest(http.MethodPost, "http://localhost:3000/comments/", bodyComment)
	requestAdd.Header.Add("Content-Type", "application/json")
	requestAdd.Header.Add("Authorization", token)

	recorderAdd := httptest.NewRecorder()
	router.ServeHTTP(recorderAdd, requestAdd)
	responseAdd := recorderAdd.Result()

	body, _ := io.ReadAll(responseAdd.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	commentId := responseBody["id"]
	commentId = int(commentId.(float64))

	commentUpdate := `{
    "message": "salam kenal dari user 1 update"
	}`
	bodyCommentUpdate := strings.NewReader(commentUpdate)

	url := fmt.Sprintf("http://localhost:3000/comments/%d", commentId)

	requestUpdate := httptest.NewRequest(http.MethodPut, url, bodyCommentUpdate)
	requestUpdate.Header.Add("Content-Type", "application/json")
	requestUpdate.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, requestUpdate)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateCommentFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	_ = TemporaryCreatePhoto(router, tokenUser)

	token := "Bearer " + tokenUser

	commentId := 0

	commentUpdate := `{
    "message": "salam kenal dari user 1 update"
	}`
	bodyCommentUpdate := strings.NewReader(commentUpdate)

	url := fmt.Sprintf("http://localhost:3000/comments/%d", commentId)

	requestUpdate := httptest.NewRequest(http.MethodPut, url, bodyCommentUpdate)
	requestUpdate.Header.Add("Content-Type", "application/json")
	requestUpdate.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, requestUpdate)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestDeleteCommentSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	photoId := TemporaryCreatePhoto(router, tokenUser)

	comment := fmt.Sprintf(`{
    "message": "salam kenal dari user 1",
    "photo_id": %d
	}`, photoId)
	bodyComment := strings.NewReader(comment)

	token := "Bearer " + tokenUser
	requestAdd := httptest.NewRequest(http.MethodPost, "http://localhost:3000/comments/", bodyComment)
	requestAdd.Header.Add("Content-Type", "application/json")
	requestAdd.Header.Add("Authorization", token)

	recorderAdd := httptest.NewRecorder()
	router.ServeHTTP(recorderAdd, requestAdd)
	responseAdd := recorderAdd.Result()

	body, _ := io.ReadAll(responseAdd.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	commentId := responseBody["id"]
	commentId = int(commentId.(float64))

	url := fmt.Sprintf("http://localhost:3000/comments/%d", commentId)

	requestDelete := httptest.NewRequest(http.MethodDelete, url, nil)
	requestDelete.Header.Add("Content-Type", "application/json")
	requestDelete.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, requestDelete)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCommentFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	tokenUser := GetUserCredential(router)
	_ = TemporaryCreatePhoto(router, tokenUser)

	token := "Bearer " + tokenUser
	url := fmt.Sprintf("http://localhost:3000/comments/%d", 0)

	requestDelete := httptest.NewRequest(http.MethodDelete, url, nil)
	requestDelete.Header.Add("Content-Type", "application/json")
	requestDelete.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, requestDelete)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}
