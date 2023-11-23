package routes

import (
	"management-school/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.NoRoute(controllers.NotFound)

	r.GET("/students", controllers.ShowStudents)
	r.GET("/students/:id", controllers.ShowStudent)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/students/find", controllers.FindStudent)
	r.GET("/", controllers.RenderIndexPage)
	r.GET("/view/students", controllers.RenderStudentsPage)
	r.Run("localhost:8000")
}
