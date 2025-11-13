package main

import (
	"Redrock-lesson1/Redrock-lesson4/lv1/database"
	"Redrock-lesson1/Redrock-lesson4/lv1/service"
	"fmt"
)

func main() {
	database.Init()

	err := demo()
	if err != nil {
		panic(err)
	}
}

func demo() error {
	var isSeedSample bool
	for {
		fmt.Println("\n================================================================")
		fmt.Printf("\nAll information of the students and courses\n\n")
		students, courses, err := service.CheckInfo()
		if err != nil {
			return err
		}
		fmt.Printf("Students information:\n")
		for _, student := range students {
			fmt.Printf("ID: %d, name: %s, grade: %s, class: %s\n", student.ID, student.Name, student.Grade, student.Class)
		}
		fmt.Printf("\nCourses information:\n")
		for _, course := range courses {
			fmt.Printf("ID: %d, name: %s, capital: %d, enrollment: %d\n", course.ID, course.Name, course.Capital, course.Enroll)
		}
		fmt.Println("\n\nPress\n`1` to seed sample\n'2' to new student\n'3' to new course\n'4' to add enrollment\n'5' to drop enrollment\n'6' to check enrollments\n'-1' to delete all information\n'0' to leave the demo")
		var choose int
		fmt.Scan(&choose)
		switch choose {
		case 1:
			if isSeedSample {
				fmt.Println("You have already did this")
				continue
			}
			isSeedSample = true
			service.SeedSample()
			fmt.Println("Seeded successfully")
		case 2:
			var (
				name  string
				grade string
				class string
			)
			fmt.Printf("Input the student name, grade and class:\n")
			fmt.Scan(&name, &grade, &class)
			err := service.AddStudent(name, grade, class)
			if err != nil {
				fmt.Println("\nFailed AddStudent:", err)
			}
			fmt.Println("\nAdded successfully")
		case 3:
			var (
				name    string
				capital int
			)
			fmt.Printf("Input the course name and capital:\n")
			fmt.Scan(&name, &capital)
			err := service.AddCourse(name, capital)
			if err != nil {
				fmt.Println("\nFailed AddCourse:", err)
			}
			fmt.Println("\nAdded successfully")
		case 4:
			var (
				studentID int
				courseID  int
			)
			fmt.Printf("Input the studentID and courseID:\n")
			fmt.Scan(&studentID, &courseID)
			err := service.EnrollCourse(studentID, courseID)
			if err != nil {
				fmt.Println("\nFailed Enroll:", err)
			}
			fmt.Println("\nEnrolled successfully")
		case 5:
			var (
				studentID int
				courseID  int
			)
			fmt.Printf("Input the studentID and courseID:\n")
			fmt.Scan(&studentID, &courseID)
			err := service.DropCourse(studentID, courseID)
			if err != nil {
				fmt.Println("\nFailed DropCourse:", err)
			}
			fmt.Println("\nDrop successfully")
		case 6:
			var enrollments []database.Enrollment
			enrollments, err := service.CheckEnrollment()
			if err != nil {
				fmt.Println("\nFailed CheckEnrollment:", err)
			}
			fmt.Println("\nStudentID   --->   courseID")
			for _, enrollment := range enrollments {
				fmt.Printf("%d           --->          %d", enrollment.StudentID, enrollment.CourseID)
			}
		case 0:
			return nil
		case -1:
			database.DB.Where("1 = 1").Delete(&database.Course{})
			database.DB.Where("1 = 1").Delete(&database.Student{})
			database.DB.Where("1 = 1").Delete(&database.Enrollment{})
			database.DB.Exec("ALTER TABLE students AUTO_INCREMENT = 1")
			database.DB.Exec("ALTER TABLE courses AUTO_INCREMENT = 1")
		}
	}
}
