package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	// run url mapping
	mapUrls()
	// start at port :8080
	router.Run(":8080")
}