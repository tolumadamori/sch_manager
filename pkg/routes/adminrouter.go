package routes

import (
	"github.com/tolumadamori/sch_manager/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// Admin Routes.
func AdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/admin/liststudentcourses/:matric_number", controllers.ListCourses())         // lists all the courses taken by a student.
	incomingRoutes.GET("/admin/listcoursetudents/:ID", controllers.ListStudents())                    //lists the students taking a particular course when the course ID is passed as a path param
	incomingRoutes.POST("/admin/updatestudentcourses/:matric_number", controllers.UpdateCourseList()) //updates the list of courses taken by a student using the course IDs passed in the request body.
	incomingRoutes.POST("/admin/addcourse/:matric_number/:ID", controllers.AddCourse())               //Adds the course specified by ID in the path param to the student specified by matric number in the path param.
	incomingRoutes.DELETE("/admin/removecourse/:matric_number/:ID", controllers.RemoveCourse())       // Removes the course specified by ID in the Path params from the list of courses taken by the specified student.

}
