package routes

import (
	"api_rest_gin/controllers"

	"github.com/gin-gonic/gin"
)

func Handle() {
	r := gin.Default()
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/students/:document", controllers.GetStudentByDocument)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.Run()
}
