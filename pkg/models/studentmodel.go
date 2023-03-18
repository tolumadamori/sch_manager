package models

import "time"

//Student Model
type Student struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	MatricNumber uint   `gorm:"primarykey;index" json:"matric_number"`
	Department   string `json:"department"`
	Level        string `json:"level"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Courses      []Course `gorm:"many2many:student_courses;foreignkey:matric_number;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:matric_number;" json:"-"`
}
