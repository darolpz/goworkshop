package main

import (
	"fmt"
	"strconv"

	"github.com/darolpz/workshop/15_orm/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Conexion a db
	db, err := gorm.Open("mysql", "dummy:12345@/workshop?charset=utf8&parseTime=True&loc=Local")
	//Si falla imprime el error
	if err != nil {
		fmt.Println(err)
	}
	// Al finalizar la ejecucion del bloque se cierra la conexion
	defer db.Close()

	//Migracion de db
	db.DropTableIfExists(models.ClassStudents{}, models.Student{}, models.Class{}, models.Teacher{})
	db.AutoMigrate(models.Teacher{}, models.Class{}, models.Student{})

	// Se crean las foreign keys
	db.AutoMigrate(&models.Class{}).AddForeignKey("teacher_id", "teachers(id)", "CASCADE", "CASCADE")
	db.Table("class_students").AddForeignKey("class_id", "classes(id)", "CASCADE", "CASCADE")
	db.Table("class_students").AddForeignKey("student_id", "students(id)", "CASCADE", "CASCADE")

	// Creo maestros
	teacher1 := models.Teacher{FirstName: "Ricardo", Lastname: "Fort"}
	teacher2 := models.Teacher{FirstName: "Carlin", Lastname: "Calvo"}
	db.Create(&teacher1)
	db.Create(&teacher2)

	// Creo clases
	class1 := models.Class{Name: "Como manejar un roll royce", TeacherID: teacher1.ID}
	class2 := models.Class{Name: "Como hackear todo", TeacherID: teacher2.ID}
	db.Create(&class1)
	db.Create(&class2)

	//Creo alumnos
	var students [3]models.Student
	for index := range students {
		students[index].FirstName = "Student"
		students[index].Lastname = "Number " + strconv.Itoa(index+1)
		db.Create(&students[index])
	}
	//Asocio alumnos con clases
	db.Model(&class1).Association("Students").Append(students[0])
	db.Model(&class2).Association("Students").Append(students[0])
	db.Model(&class1).Association("Students").Append(students[1])
	db.Model(&class2).Association("Students").Append(students[2])

	//Count de relaciones
	countClass := db.Model(&class1).Association("Students").Count()
	countStudent := db.Model(&students[1]).Association("Classes").Count()
	fmt.Printf("Cantidad de alumnos en clase 1: %v \n", countClass)
	fmt.Printf("Cantidad de clases a las que asiste alumno 2: %v \n", countStudent)

	// Clases a las que asiste alumno 1 y 2
	var classes0, classes1 []models.Class
	db.Model(&students[0]).Related(&classes0, "Classes")
	db.Model(&students[1]).Related(&classes1, "Classes")

	fmt.Printf("Clases a las que asiste %v: %v \n", students[0].FirstName+students[0].Lastname, classes0)
	fmt.Printf("Clases a las que asiste %v: %v \n", students[1].FirstName+students[1].Lastname, classes1)

	// Alumnos que asisten a clase 2
	var students1 []models.Student
	db.Model(&class2).Related(&students1, "Students")
	fmt.Printf("Alumnos de clase 2: %v", students1)

}
