package main

import (
	"github.com/aynp/urlshortner/controllers"
	"github.com/aynp/urlshortner/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	r := gin.Default()

	// health check
	r.GET("/health", controllers.HealthCheck)

	// create our shortend url
	r.POST("/v1/tiny/create", controllers.CreateShortURL)

	// redirect to original url
	r.GET("/v1/tiny/redirect/:path", controllers.Redirect)

	r.Run()
}
