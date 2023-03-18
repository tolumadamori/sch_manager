package routes

import (
	"github.com/tolumadamori/sch_manager/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// Course routes.
func CourseRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/course", controllers.GetCourses())          //Lists all courses created.
	incomingRoutes.GET("/course/:ID", controllers.GetCourse())       //Returns the course specified by course code in the path.
	incomingRoutes.POST("/course", controllers.CreateCourse())       //Creates a new course in the DB.
	incomingRoutes.PATCH("/course/:ID", controllers.UpdateCourse())  //Updates the course details. New details are passed in the course body.
	incomingRoutes.DELETE("/course/:ID", controllers.DeleteCourse()) //Deletes the course with the course code pass in the path.
}
