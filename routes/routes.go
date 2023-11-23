package routes

import (
	"management-school/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowStudents)
	r.GET("/students/:id", controllers.ShowStudent)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/students/find", controllers.FindStudent)
	r.Run("localhost:8000")
}
