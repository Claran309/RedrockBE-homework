package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Student 学生模型
type Student struct {
	ID    int    `gorm:"primary_key;auto_increment;column:student_id"`
	Name  string `gorm:"column:name"`
	Grade string `gorm:"column:grade"`
	Class string `gorm:"column:class"`
	//关联
	Enrollments []Enrollment `gorm:"foreignKey:StudentID"`
}

// Course 课程模型
type Course struct {
	ID      int    `gorm:"primary_key;auto_increment;column:course_id"`
	Name    string `gorm:"column:name"`
	Capital int    `gorm:"column:capital"`
	Enroll  int    `gorm:"column:enroll"`
	//关联
	Enrollments []Enrollment `gorm:"foreignKey:CourseID"`
}

// Enrollment 选课记录模型
type Enrollment struct {
	StudentID int `gorm:"column:student_id"`
	CourseID  int `gorm:"column:course_id"`
	//定义外键关联
	Student Student `gorm:"foreignKey:StudentID;references:StudentID"`
	Course  Course  `gorm:"foreignKey:CourseID;references:CourseID"`
}

func Init() {
	dsn := "claran:chr070309@tcp(localhost:3306)/Redrock_lesson3?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //Debug result
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&Student{}, &Course{}); err != nil {
		log.Fatal("11111111111111111", err)
	}
	if err := DB.AutoMigrate(&Enrollment{}); err != nil {
		log.Fatal("222222222222222222", err)
	}
}
