package controllers

import (
	"management-school/database"
	"management-school/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func ShowStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(404, gin.H{
			"statusCode": 404,
			"message":    "Student not found"})
		return
	}

	c.JSON(200, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(404, gin.H{
			"statusCode": 404,
			"message":    "Student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	database.DB.Save(&student)

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(404, gin.H{
			"statusCode": 404,
			"message":    "Student not found"})
		return
	}

	database.DB.Delete(&student)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"message":    "Student deleted successfully"})
}
