package main

import (
	"net/http"
	"tuition-api/api"
	"tuition-api/data"

	"github.com/gin-gonic/gin"
)

/*
Given a students query as text, handles the query by:

-- Parsing student query by deserializing the json. Formatting it as a prompt for the Teacher.

-- Submitting prompt to Teacher, a logical wrapper around our LLM

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

	var teacherResponse, err = api.GetResponseFromTeacher(studentQueryData)

	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}

	c.JSON(http.StatusOK, teacherResponse)
}
