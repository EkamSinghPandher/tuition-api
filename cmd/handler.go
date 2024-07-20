package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
A sanity check endpoint to ensure Router is working.
*/
func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
