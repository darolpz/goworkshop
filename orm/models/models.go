package models

import (
	"time"
)

type Teacher struct {
	ID        uint `gorm:"primary_key"`
	FirstName string
	Lastname  string
	Classes   []Class
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Student struct {
	ID        uint `gorm:"primary_key"`
	FirstName string
	Lastname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Classes   []*Class `gorm:"many2many:class_students;"`
}

type Class struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Teacher   Teacher
	TeacherID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Students  []*Student `gorm:"many2many:class_students;"`
}

type ClassStudents struct {
	ClassID   uint
	StudentID uint
}
