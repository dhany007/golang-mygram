package testsapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func RegisterUserTest(request string, router *gin.Engine) (userId int) {
	requestBodyRegister := strings.NewReader(request)

	requestRegister := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/register", requestBodyRegister)
	recorderRegister := httptest.NewRecorder()
	router.ServeHTTP(recorderRegister, requestRegister)

	responseRegister := recorderRegister.Result()
	defer responseRegister.Body.Close()

	body, _ := io.ReadAll(responseRegister.Body)
	response := map[string]interface{}{}
	json.Unmarshal(body, &response)

	userId = int(response["id"].(float64))
	return
}

func LoginUserTest(request string, router *gin.Engine) (tokenUser string) {
	requestBodyLogin := strings.NewReader(request)

	requestLogin := httptest.NewRequest(http.MethodPost, "http://localhost:3000/users/login", requestBodyLogin)
	recorderLogin := httptest.NewRecorder()

	router.ServeHTTP(recorderLogin, requestLogin)
	responseLogin := recorderLogin.Result()
	defer responseLogin.Body.Close()

	body, _ := io.ReadAll(responseLogin.Body)
	responseBodyLogin := map[string]interface{}{}
	json.Unmarshal(body, &responseBodyLogin)

	tokenUser = responseBodyLogin["token"].(string)
	return
}

func GetUserCredential(router *gin.Engine) (token string) {
	registerUserTest := `{
		"age": 8,
		"email": "dhany2@gmail.com",
		"password": "dhany2",
		"username": "dhany2"
	}`

	_ = RegisterUserTest(registerUserTest, router)

	bodyLoginUser := `{
		"email": "dhany2@gmail.com",
		"password": "dhany2"
	}`

	token = LoginUserTest(bodyLoginUser, router)
	return
}
