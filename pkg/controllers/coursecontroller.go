package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tolumadamori/sch_manager/pkg/models"
)

// GetCourses			godoc
// @Summary				Get Courses
// @Description			Returns all the existing courses.
// @Produce				application/json
// @Tags				Course
// @Success				200 {object} models.CourseSuccessResponse
// @Router				/course [get]
func GetCourses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var courses []models.Course

		db := dbConnecter()

		if err := db.Find(&courses).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Courses fetched successfully", "data": courses})
	}

}

// GetCoursebyID 		godoc
// @Summary				Get Course by ID
// @Description			Returns a course. Course to return is specified by ID in path Param.
// @Param				ID path string true "Course ID"
// @Produce				application/json
// @Tags				Course
// @Success				200 {object} models.CourseSuccessResponse
// @Router				/course/{ID} [get]
func GetCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		db := dbConnecter()
		if err := db.Where("ID=?", c.Param("ID")).First(&course).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course fetched successfully", "data": course})

	}
}

// CreateCourse		godoc
// @Summary			Create Course
// @Description		Creates a course in the DB. Course Details are passed in the request body.
// @Param			course body models.SwagCourse true "Course to Create"
// @Produce			application/json
// @Tags			Course
// @Success			201 {object} models.CourseSuccessResponse
// @Router			/course [post]
func CreateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		if err := c.ShouldBindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		newcourse := models.Course{ID: course.ID, CourseName: course.CourseName, CourseCode: course.CourseCode, Lecturer: course.Lecturer}

		db := dbConnecter()

		if err := db.Create(&newcourse).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": true, "message": "Course created successfully", "data": newcourse})
	}
}

// UpdateCourse		godoc
// @Summary			Update Course
// @Description		Updates a course. Course is specified by ID. Update details are passed in the request body.
// @Param			ID path string true "Course ID"
// @Param			Updatecourse body models.SwagUpdateCourse true "Update Details"
// @Tags			Course
// @Produce			application/json
// @Success			200 {object} models.CourseSuccessResponse
// @Router			/course/{ID} [patch]
func UpdateCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		db := dbConnecter()

		if err := db.Where("ID=?", c.Param("ID")).First(&course).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
			return
		}

		var updatedcourse models.Course
		if err := c.ShouldBindJSON(&updatedcourse); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		if err := db.Model(&course).Updates(models.Course{CourseName: updatedcourse.CourseName, CourseCode: updatedcourse.CourseCode, Lecturer: updatedcourse.Lecturer}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course updated successfully", "data": course})
	}
}

// DeleteCourse		godoc
// @Summary			Delete Course
// @Description		Deletes a course. Course to delete is specified by ID in path param.
// @Param			ID path string true "Course id"
// @Produce			application/json
// @Tags			Course
// @Success			200 {object} models.GenericSuccessResponse
// @Router			/course/{ID} [delete]
func DeleteCourse() gin.HandlerFunc {
	return func(c *gin.Context) {
		var course models.Course

		db := dbConnecter()

		if err := db.Where("ID=?", c.Param("ID")).Delete(&course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course deleted successfully"})
	}
}
