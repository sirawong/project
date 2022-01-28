package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	cors "github.com/rs/cors/wrapper/gin"

	"github.com/gin-gonic/gin"
)

var serviceName string = "media"
var folder string = "files"
var path string = "file"
var rootFolder string = "root/files"
var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// type File struct {
// 	FileName string `bson:"file_name"`
// 	URL      string `bson:"url"`
// 	Link     string `bson:"link" json:"link"`
// }

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	router := gin.Default()
	// router.Use(cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{
	// 		http.MethodHead,
	// 		http.MethodGet,
	// 		http.MethodPost,
	// 		http.MethodPut,
	// 		http.MethodPatch,
	// 		http.MethodDelete},
	// 	ExposedHeaders:   []string{"X-Content-Length"},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: false,
	// }))
	router.Use(cors.AllowAll())
	group := router.Group("media")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	group.Static(path, folder)
	group.GET("/health", func(c *gin.Context) {
		f, err := os.Open(folder)
		if err != nil {
			f.Close()
			response := Response{
				Code:    http.StatusNotFound,
				Data:    []string{},
				Message: "file not found",
			}
			c.JSON(http.StatusNotFound, response)
			return
		}
		response := Response{
			Code:    http.StatusOK,
			Data:    []string{},
			Message: "found a file",
		}
		c.JSON(http.StatusOK, response)

	})

	group.POST("/upload", func(c *gin.Context) {
		// single file
		log.Println("check")
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// generate filename
		filename := RandString(10) + "-" + file.Filename

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, folder+"/"+filename)

		u := &url.URL{
			Scheme: "http",
			Host:   c.Request.Host,
			Path:   serviceName + "/" + path + "/" + filename,
		}

		log.Println(u)
		// response := File{
		// 	FileName: filename,
		// 	URL:      u.String(),
		// 	Link:     u.String(),
		// }
		c.JSON(http.StatusOK, u.String())
	})

	router.Run(":9000")
}

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
