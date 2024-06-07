package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department int    `json:"department"`
	Math       int    `json:"math"`
	English    int    `json:"english"`
	IT         int    `json:"it"`
}

var db *gorm.DB
var err error

func init() {
	dsn := "root:@tcp(localhost:3306)/python_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Student{})
}

func getStudents(c *gin.Context) {
	var students []Student
	db.Find(&students)
	c.JSON(http.StatusOK, students)
}

func createStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&student)
	c.JSON(http.StatusOK, student)
}

func main() {
	r := gin.Default()

	r.GET("/students", getStudents)
	r.POST("/students", createStudent)

	fmt.Println("Starting server at port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
