package main

import (
	"net/http"
	"tuition-api/data"

	"github.com/gin-gonic/gin"
)

/*
Given a students query as text, handles the query by:

-- Parsing student query by deserializing the json. Formatting it as a prompt for the Teacher.

-- Submititng prompt to Teacher, a logical wrapper around our LLM

-- Formatting response from Teacher
*/
func handleStudentQuery(c *gin.Context) {
	var studentQueryData data.StudentQuery
	bindErr := c.ShouldBindJSON(&studentQueryData)

	if bindErr != nil {
		r := data.Message{Message: "Bad Request Body"}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": studentQueryData.StudentQuestion})
}
