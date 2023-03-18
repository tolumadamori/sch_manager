package models

import "time"

//Course Model
type Course struct {
	ID         uint   `gorm:"primarykey" json:"ID"`
	CourseName string `json:"course_name"`
	CourseCode string `json:"course_code"`
	Lecturer   string `json:"lecturer"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Students   []Student `gorm:"many2many:student_courses;foreignkey:ID;association_foreignkey:matric_number;association_jointable_foreignkey:matric_number;jointable_foreignkey:ID;" json:"-"`
}
