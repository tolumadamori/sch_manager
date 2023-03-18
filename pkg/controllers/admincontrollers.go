package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tolumadamori/sch_manager/pkg/models"
)

// GetStudentCourses	godoc
// @Summary				Get Student Courses
// @Description			Returns all the courses assigned to a Student. The student is specified by matric number in the path.
// @Param				matric_number path string true "Matric Number"
// @Produce				application/json
// @Tags				Admin
// @Success				200 {object} models.AdminCourseList
// @Router				/admin/liststudentcourses/{matric_number} [get]
func ListCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student

		db := dbConnecter()

		if err := db.Where("matric_number=?", c.Param("matric_number")).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}
		var coursesTaken = []models.Course{}

		if err := db.Model(&student).Association("Courses").Find(&coursesTaken); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "List of courses returned successfully", "data": coursesTaken})

	}
}

// GetCourseStudents	godoc
// @Summary				Get Course Assignees
// @Description			Returns all the students assigned to a course. The course is specified by ID in the path.
// @Param				ID path string true "Course ID"
// @Produce				application/json
// @Tags				Admin
// @Success				200 {object} models.AdminStudentList
// @Router				/admin/listcoursetudents/{ID} [get]
func ListStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		db := dbConnecter()

		if err := db.Where("ID=?", c.Param("ID")).First(&course).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}
		var studentsTakingCourse = []models.Student{}

		if err := db.Model(&course).Association("Students").Find(&studentsTakingCourse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "List of students returned successfully", "data": studentsTakingCourse})

	}
}

// AddCourseToStudent	godoc
// @Summary				Add Course to Student
// @Description			Assign a single specified course to a student. Student and course are specified as path params.
// @Param				matric_number path string true "Matric Number"
// @Param				ID path string true "Course ID"
// @Produce				application/json
// @Tags				Admin
// @Success				200 {object} models.GenericSuccessResponse
// @Router				/admin/addcourse/{matric_number}/{ID} [post]
func AddCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		var student models.Student

		db := dbConnecter()
		if err := db.Where("matric_number=?", c.Param("matric_number")).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		var course models.Course

		if err := db.Where("ID =?", c.Param("ID")).First(&course).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		if err := db.Model(&student).Association("Courses").Append(&course); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course " + c.Param("ID") + " added successfully to student " + c.Param("matric_number")})
	}
}

// UpdateCourseList	godoc
// @Summary				Update Course List
// @Description			Assign multiple courses specified by Course ID to a student. Course IDs are passed as a comma delimited array in the request body. The difference between this route and the AddCourse route is that multiple courses can be added at once.
// @Param				matric_number path string true "Matric Number"
// @Param				Course_List body models.CourseList true "Course List"
// @Produce				application/json
// @Tags				Admin
// @Success				200 {object} models.GenericSuccessResponse
// @Router				/admin/updatestudentcourses/{matric_number} [post]
func UpdateCourseList() gin.HandlerFunc {
	return func(c *gin.Context) {

		var student models.Student

		db := dbConnecter()

		if err := db.Where("matric_number=?", c.Param("matric_number")).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		var listOfCourses []int
		if err := c.ShouldBindJSON(&listOfCourses); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		var coursesToAppend = []models.Course{}

		for _, v := range listOfCourses {

			var course models.Course
			if err := db.Where("ID =?", v).First(&course).Error; err != nil {
				c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
				return
			} else {
				coursesToAppend = append(coursesToAppend, course)
			}
		}

		if err := db.Model(&student).Association("Courses").Append(&coursesToAppend); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": fmt.Sprintf("The following courses %v were added successfully to student %v", listOfCourses, c.Param("matric_number"))})
	}
}

// RemoveCourseFromStudent	godoc
// @Summary				Remove Course From Student
// @Description			Removes a single specified course from a student. Student and course are specified as path params.
// @Param				matric_number path string true "Matric Number"
// @Param				ID path string true "Course ID"
// @Produce				application/json
// @Tags				Admin
// @Success				200 {object} models.GenericSuccessResponse
// @Router				/admin/removecourse/{matric_number}/{ID} [delete]
func RemoveCourse() gin.HandlerFunc {
	return func(c *gin.Context) {

		var student models.Student

		db := dbConnecter()

		if err := db.Where("matric_number=?", c.Param("matric_number")).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		var course = []models.Course{}

		if err := db.Where("ID =?", c.Param("ID")).First(&course).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		if err := db.Model(&student).Association("Courses").Delete(&course); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course " + c.Param("ID") + " removed successfully from student " + c.Param("matric_number")})
	}
}
