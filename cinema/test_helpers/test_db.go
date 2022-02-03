package testhelper

import (
	"bytes"
	"cinema/entities"
	"log"
	"net/http"
	"net/http/httptest"

	gomocket "github.com/Selvatico/go-mocket"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetMockCinemaEntity for unit test
func GetMockCinemaEntity() entities.Cinema {
	seats := [][]int{}
	for i := 0; i < 2; i++ {
		seats[i] = []int{0}
	}

	mockCinemaEntity := entities.Cinema{
		ID:             "1",
		Name:           "dev",
		TicketPrice:    99,
		City:           "test",
		Seats:          seats,
		SeatsAvailable: 99,
		Image:          "url",
	}
	return mockCinemaEntity
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
