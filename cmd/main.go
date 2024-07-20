package main

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//Setup Router
	router := gin.Default()

	apiGroup := router.Group("/api/v1")
	{

		testgroup := apiGroup.Group("/test")
		{
			testgroup.GET("/ping", handlePing)
		}
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
