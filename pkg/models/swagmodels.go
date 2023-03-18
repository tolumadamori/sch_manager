package models

/*-------------------------- Swagger Generic Structs ---------------------------*/
type GenericSuccessResponse struct {
	Status  bool
	Message string
}

/*-------------------------- Swagger Course Structs ---------------------------*/

type CourseSuccessResponse struct {
	Status  bool
	Message string
	Data    Course
}

type SwagCourse struct {
	ID         uint   `gorm:"primarykey" json:"ID"`
	CourseName string `json:"course_name"`
	CourseCode string `json:"course_code"`
	Lecturer   string `json:"lecturer"`
}

type SwagUpdateCourse struct {
	CourseName string `json:"course_name"`
	CourseCode string `json:"course_code"`
	Lecturer   string `json:"lecturer"`
}

/*-------------------------- Swagger Student Structs ---------------------------*/

type StudentSuccessResponse struct {
	Status  bool
	Message string
	Data    Student
}

type SwagStudent struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	MatricNumber uint   `gorm:"primarykey;index" json:"matric_number"`
	Department   string `json:"department"`
	Level        string `json:"level"`
}

type SwagUpdateStudent struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Level      string `json:"level"`
}

/*-------------------------- Swagger Admin Structs ---------------------------*/
type AdminCourseList struct {
	Status  bool
	Message string
	Data    []Course
}

type AdminStudentList struct {
	Status  bool
	Message string
	Data    []Student
}

type CourseList []uint
