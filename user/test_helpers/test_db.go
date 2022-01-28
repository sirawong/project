package testhelper

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"user/entities"

	gomocket "github.com/Selvatico/go-mocket"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetMockShowtimeEntity for unit test
func GetMockShowtimeEntity() entities.User {
	seats := [][]int32{}
	for i := 0; i < 2; i++ {
		seats[i] = []int32{0}
	}

	mockShowtimeEntity := entities.User{
		ID:       "",
		Name:     "",
		Username: "",
		Email:    "",
		Password: "",
		Role:     "",
		Phone:    "",
		Imageurl: "",
		Tokens:   []*entities.Token{},
	}
	return mockShowtimeEntity
}

func SetupMockDB() *gorm.DB {
	gomocket.Catcher.Register()

	db, err := gorm.Open(gomocket.DriverName, "")
	if err != nil {
		log.Fatalf("error mocking up the database %s", err)
	}

	db.LogMode(true)

	return db
}

func MakeStubContext(method string, url string, params string) (c *gin.Context) {
	const MIMEJSON = "application/json"

	body := bytes.NewBufferString(params)

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request, _ = http.NewRequest(method, url, body)
	context.Request.Header.Add("Content-Type", MIMEJSON)

	return context
}
