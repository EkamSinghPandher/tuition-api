package main

import (
	// "github.com/gin-contrib/cors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables from environment file
	loadErr := godotenv.Load(".env")

	if loadErr != nil {
		fmt.Println("Unabe to fetch .env file")
		return
	}

	//Setup Router
	router := gin.Default()

	apiGroup := router.Group("/api/v1")
	{

		testgroup := apiGroup.Group("/test")
		{
			testgroup.GET("/ping", handlePing)
		}

		studentgroup := apiGroup.Group("/student")
		{
			studentgroup.POST("/query", handleStudentQuery)
		}
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
