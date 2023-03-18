package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tolumadamori/sch_manager/pkg/config"
	"github.com/tolumadamori/sch_manager/pkg/models"
	"gorm.io/gorm"
)

// Wrapper function to connect to the DB and check for errors while connecting.
func dbConnecter() *gorm.DB {
	db, err := config.ConnectDB()
	if err != nil {
		log.Println(err)
	}

	return db
}

// GetStudents			godoc
// @Summary				Get Students
// @Description			Returns all the existing students.
// @Produce				application/json
// @Tags				Student
// @Success				200 {object} models.StudentSuccessResponse
// @Router				/student [get]
func GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []models.Student

		db := dbConnecter()
		if err := db.Find(&students).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Students fetched successfully", "data": students})
	}

}

// GetStudentbyID 		godoc
// @Summary				Get Student by ID
// @Description			Returns a Student. Student to return is specified by Matric Number in path Param.
// @Param				matric_number path string true "Matric Number"
// @Produce				application/json
// @Tags				Student
// @Success				200 {object} models.StudentSuccessResponse
// @Router				/student/{matric_number} [get]
func GetStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		db := dbConnecter()
		if err := db.Where("matric_number=?", c.Param("matric_number")).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Student fetched successfully", "data": student})
	}
}

// CreateStudent	godoc
// @Summary			Create Student
// @Description		Creates a Student in the DB. Student Details are passed in the request body.
// @Param			Student body models.SwagStudent true "Student to Create"
// @Produce			application/json
// @Tags			Student
// @Success			201 {object} models.StudentSuccessResponse
// @Router			/student [post]
func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		newStudent := models.Student{Name: student.Name, Email: student.Email, MatricNumber: student.MatricNumber, Department: student.Department, Level: student.Level}

		db := dbConnecter()

		if err := db.Create(&newStudent).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": true, "message": "Student created successfully", "data": student})

	}

}

// UpdateStudent	godoc
// @Summary			Update Student
// @Description		Updates a Student. Student is specified by Matric Number. Update details are passed in the request body. All Fields are not required.
// @Param			matric_number path string true "Matric Number"
// @Param			UpdateStudent body models.SwagUpdateStudent true "Update Details"
// @Tags			Student
// @Produce			application/json
// @Success			200 {object} models.StudentSuccessResponse
// @Router			/student/{matric_number} [patch]
func UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		db := dbConnecter()

		if err := db.Where("matric_number=?", c.Param("matric_number")).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		var updatedStudent models.Student
		if err := c.ShouldBindJSON(&updatedStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		if err := db.Model(&student).Updates(models.Student{Name: updatedStudent.Name, Email: updatedStudent.Email, Department: updatedStudent.Department, Level: updatedStudent.Level}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Student updated successfully", "data": student})
	}

}

// DeleteStudent	godoc
// @Summary			Delete Student
// @Description		Deletes a Student. Student to delete is specified by matric number in path param.
// @Param			matric_number path string true "Matric Number"
// @Produce			application/json
// @Tags			Student
// @Success			200 {object} models.GenericSuccessResponse
// @Router			/student/{matric_number} [delete]
func DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		db := dbConnecter()

		if err := db.Where("matric_number=?", c.Param("matric_number")).Delete(&student).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Student deleted successfully"})

	}
}
