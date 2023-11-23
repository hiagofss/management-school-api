package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"management-school/controllers"
	"management-school/database"
	"management-school/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var studentMock models.Student

func CreateMockStudent() {
	student := models.Student{
		Name:     "John Doe",
		Email:    "john.doe@test.com",
		Document: "12345678901"}
	database.DB.Create(&student)
	studentMock = student
}

func DeleteMockStudent() {
	var student models.Student
	database.DB.Unscoped().Delete(&student, studentMock.ID)
}

func SetupRouterTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestCreateStudent(t *testing.T) {
	database.ConnectDatabase()

	r := SetupRouterTest()
	r.POST("/students", controllers.CreateStudent)

	// Test 1: Valid student
	w := httptest.NewRecorder()
	jsonStr := []byte(`{"name": "John Doe", "email": "john.doe@test.com", "document": "12345678901"}`)
	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var student models.Student

	if err := json.NewDecoder(w.Body).Decode(&student); err != nil {
		fmt.Println(err)
	}

	database.DB.Unscoped().Delete(&student, student.ID)
}

func TestGetStudents(t *testing.T) {
	database.ConnectDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupRouterTest()
	r.GET("/students", controllers.ShowStudents)

	// Test 1: Valid student
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/students", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupRouterTest()
	r.GET("/students/:id", controllers.ShowStudent)

	// Test 1: Valid student
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/students/"+strconv.FormatUint(uint64(studentMock.ID), 10), nil)
	r.ServeHTTP(w, req)

	var student models.Student
	json.Unmarshal(w.Body.Bytes(), &student)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, studentMock.ID, student.ID)
}

func TestGetStudentByDocument(t *testing.T) {
	database.ConnectDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupRouterTest()
	r.GET("/students/find", controllers.FindStudent)

	// Test 1: Valid student
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/students/find?document="+studentMock.Document, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupRouterTest()
	r.PUT("/students/:id", controllers.UpdateStudent)

	// Test 1: Valid student
	w := httptest.NewRecorder()
	jsonStr := []byte(`{"name": "John Doe Renamend", "email": "test.update@test.com", "document": "12345678999"}`)
	req, _ := http.NewRequest("PUT", "/students/"+strconv.FormatUint(uint64(studentMock.ID), 10), bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)
	var student models.Student
	json.Unmarshal(w.Body.Bytes(), &student)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "John Doe Renamend", student.Name)
	assert.Equal(t, "12345678999", student.Document)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()

	r := SetupRouterTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/students/"+strconv.FormatUint(uint64(studentMock.ID), 10), nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
