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
	r.Run("localhost:8000")
}
