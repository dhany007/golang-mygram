package testsapi

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func TemporaryCreatePhoto(router *gin.Engine, token string) (photoId int) {
	photo := `{
    "title": "photo indah",
    "caption": "ini adalah photo indah",
    "photo_url": "https://img.okezone.com/content/2018/08/27/406/1942193/daerah-daerah-misterius-di-pegunungan-himalaya-salah-satunya-tempat-tinggal-yeti-JqpTeHvIOb.jpg"
	}`
	bodyPhoto := strings.NewReader(photo)

	tokenUser := "Bearer " + token
	requestPhoto := httptest.NewRequest("POST", "http://localhost:3000/photos/", bodyPhoto)
	requestPhoto.Header.Add("Content-Type", "application/json")
	requestPhoto.Header.Add("Authorization", tokenUser)

	recorderPhoto := httptest.NewRecorder()

	router.ServeHTTP(recorderPhoto, requestPhoto)

	responseAddPhoto := recorderPhoto.Result()

	body, _ := io.ReadAll(responseAddPhoto.Body)
	responseBody := map[string]interface{}{}
	json.Unmarshal(body, &responseBody)

	tempPhotoId := responseBody["id"]
	photoId = int(tempPhotoId.(float64))
	return
}
