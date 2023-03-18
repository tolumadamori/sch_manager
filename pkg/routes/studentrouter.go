package routes

import (
	"github.com/tolumadamori/sch_manager/pkg/controllers"

	"github.com/gin-gonic/gin"
)

// Student routes.
func StudentRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/student", controllers.GetStudents())                     //Retrieves all the students created in the DB.
	incomingRoutes.GET("/student/:matric_number", controllers.GetStudent())       //Retrieves a particular student in the DB when the Matric number is passed as a path variable.
	incomingRoutes.POST("/student", controllers.CreateStudent())                  //Creates a new student in the DB. Details are passed in request body.
	incomingRoutes.PATCH("/student/:matric_number", controllers.UpdateStudent())  //Updates a student with the matric number passed as a path param and the student details in the request body.
	incomingRoutes.DELETE("/student/:matric_number", controllers.DeleteStudent()) //Deletes the student with the matric number passed in the request body.
}
