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
	"gorm.io/gorm"
)

func truncateTable(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}

func TestRegisterUserSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	requestBody := strings.NewReader(`{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/register", requestBody)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, response.StatusCode)
	assert.Equal(t, "dhany2@gmail.com", responseBody["email"])
	assert.Equal(t, "dhany2", responseBody["username"])
}

func TestRegisterUserFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	testCases := []struct {
		desc        string
		requestBody string
		expected    int
	}{
		{
			desc: "email valid",
			requestBody: `{
				"age": 8,
				"email": "dhany2",
				"password": "dhany2",
				"username": "dhany2"
			}`,
			expected: 400,
		},
		{
			desc: "email required",
			requestBody: `{
				"age": 8,
				"password": "dhany2",
				"username": "dhany2"
			}`,
			expected: 400,
		},
		{
			desc: "username required",
			requestBody: `{
				"age": 8,
				"email": "dhany2@gmail.com",
				"password": "dhany2",
			}`,
			expected: 400,
		},
		{
			desc: "age min 8",
			requestBody: `{
				"age": 5,
				"email": "dhany2@gmail.com",
				"password": "dhany2",
				"username": "dhany2"
			}`,
			expected: 400,
		},
		{
			desc: "password length min 6",
			requestBody: `{
				"age": 10,
				"email": "dhany2@gmail.com",
				"password": "dha",
				"username": "dhany2"
			}`,
			expected: 400,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			requestBody := strings.NewReader(tC.requestBody)
			request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/register", requestBody)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)
			response := recorder.Result()
			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}

func TestLoginUserSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	requestBodyRegister := strings.NewReader(`{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`)

	requestRegister := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/register", requestBodyRegister)
	recorderRegister := httptest.NewRecorder()
	router.ServeHTTP(recorderRegister, requestRegister)

	requestBodyLogin := strings.NewReader(`{
    "email": "dhany2@gmail.com",
    "password": "dhany2"
	}`)

	requestLogin := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/login", requestBodyLogin)
	recorderLogin := httptest.NewRecorder()

	router.ServeHTTP(recorderLogin, requestLogin)

	response := recorderLogin.Result()

	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
}

func TestLoginUserFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	requestBodyRegister := strings.NewReader(`{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`)

	requestRegister := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/register", requestBodyRegister)
	recorderRegister := httptest.NewRecorder()
	router.ServeHTTP(recorderRegister, requestRegister)

	testCases := []struct {
		desc        string
		requestBody string
		expected    int
	}{
		{
			desc: "user not found",
			requestBody: `{
				"email": "dhany1@gmail.com",
				"password": "dhany2",
			}`,
			expected: 400,
		},
		{
			desc: "email required",
			requestBody: `{
				"password": "dhany2"
			}`,
			expected: 400,
		},
		{
			desc: "password not match",
			requestBody: `{
				"email": "dhany2@gmail.com",
				"password": "dhany",
			}`,
			expected: 400,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			requestBody := strings.NewReader(tC.requestBody)
			request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/login", requestBody)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)
			response := recorder.Result()
			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	registerUser1 := `{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`
	userId := RegisterUserTest(registerUser1, router)

	loginUser1 := `{
    "email": "dhany2@gmail.com",
    "password": "dhany2"
	}`

	tokenUser1 := LoginUserTest(loginUser1, router)

	requestBodyUpdate := strings.NewReader(`{
    "email": "dhanyupdate@gmail.com",
    "username": "dhanyupdate"
	}`)

	url := fmt.Sprintf("http://localhost:3000/users/%d", userId)
	token := fmt.Sprintf("Bearer %s", tokenUser1)

	request := httptest.NewRequest(http.MethodPut, url, requestBodyUpdate)
	request.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
}

func TestUpdateUserFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	registerUser1 := `{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`
	userId1 := RegisterUserTest(registerUser1, router)

	loginUser1 := `{
    "email": "dhany2@gmail.com",
    "password": "dhany2"
	}`
	tokenUser1 := LoginUserTest(loginUser1, router)

	registerUser2 := `{
    "age": 8,
    "email": "dhany1@gmail.com",
    "password": "dhany1",
    "username": "dhany1"
	}`
	userId2 := RegisterUserTest(registerUser2, router)

	loginUser2 := `{
    "email": "dhany1@gmail.com",
    "password": "dhany1"
	}`
	tokenUser2 := LoginUserTest(loginUser2, router)

	requestBodyUpdate := strings.NewReader(`{
    "email": "dhanyupdate@gmail.com",
    "username": "dhanyupdate"
	}`)

	testCases := []struct {
		desc     string
		token    string
		id       int
		expected int
	}{
		{
			desc:     "user not found",
			token:    tokenUser1,
			id:       1,
			expected: 400,
		},
		{
			desc:     "token invalid",
			token:    "token salah",
			id:       userId1,
			expected: 400,
		},
		{
			desc:     "unauthorized user",
			token:    tokenUser2,
			id:       userId1,
			expected: 400,
		},
		{
			desc:     "unauthorized user",
			token:    tokenUser1,
			id:       userId2,
			expected: 400,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:3000/users/%d", tC.id)
			token := fmt.Sprintf("Bearer %s", tC.token)

			request := httptest.NewRequest(http.MethodPut, url, requestBodyUpdate)
			request.Header.Add("Authorization", token)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)
			response := recorder.Result()
			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	registerUser1 := `{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`
	userId := RegisterUserTest(registerUser1, router)

	loginUser1 := `{
    "email": "dhany2@gmail.com",
    "password": "dhany2"
	}`

	tokenUser1 := LoginUserTest(loginUser1, router)

	url := fmt.Sprintf("http://localhost:3000/users/%d", userId)
	token := fmt.Sprintf("Bearer %s", tokenUser1)

	request := httptest.NewRequest(http.MethodDelete, url, nil)
	request.Header.Add("Authorization", token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteUserFailed(t *testing.T) {
	dbTest := StartDBTest()
	router := routers.StartServer(dbTest)
	truncateTable(dbTest)

	registerUser1 := `{
    "age": 8,
    "email": "dhany2@gmail.com",
    "password": "dhany2",
    "username": "dhany2"
	}`
	userId1 := RegisterUserTest(registerUser1, router)

	loginUser1 := `{
    "email": "dhany2@gmail.com",
    "password": "dhany2"
	}`
	tokenUser1 := LoginUserTest(loginUser1, router)

	registerUser2 := `{
    "age": 8,
    "email": "dhany1@gmail.com",
    "password": "dhany1",
    "username": "dhany1"
	}`
	userId2 := RegisterUserTest(registerUser2, router)

	loginUser2 := `{
    "email": "dhany1@gmail.com",
    "password": "dhany1"
	}`
	tokenUser2 := LoginUserTest(loginUser2, router)

	testCases := []struct {
		desc     string
		token    string
		id       int
		expected int
	}{
		{
			desc:     "user not found",
			token:    tokenUser1,
			id:       1,
			expected: 400,
		},
		{
			desc:     "token invalid",
			token:    "token salah",
			id:       userId1,
			expected: 400,
		},
		{
			desc:     "unauthorized user",
			token:    tokenUser2,
			id:       userId1,
			expected: 400,
		},
		{
			desc:     "unauthorized user",
			token:    tokenUser1,
			id:       userId2,
			expected: 400,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:3000/users/%d", tC.id)
			token := fmt.Sprintf("Bearer %s", tC.token)

			request := httptest.NewRequest(http.MethodDelete, url, nil)
			request.Header.Add("Authorization", token)

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, request)
			response := recorder.Result()
			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}
