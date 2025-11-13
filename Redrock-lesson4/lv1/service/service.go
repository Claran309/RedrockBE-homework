package service

import (
	"Redrock-lesson1/Redrock-lesson4/lv1/database"
	"errors"

	"gorm.io/gorm"
)

func SeedSample() {
	student1 := database.Student{
		Name:  "claran1",
		Grade: "1",
		Class: "4",
	}
	student2 := database.Student{
		Name:  "claran2",
		Grade: "2",
		Class: "4",
	}
	student3 := database.Student{
		Name:  "claran3",
		Grade: "3",
		Class: "4",
	}
	student4 := database.Student{
		Name:  "claran4",
		Grade: "4",
		Class: "4",
	}
	database.DB.Create(&student1)
	database.DB.Create(&student2)
	database.DB.Create(&student3)
	database.DB.Create(&student4)
	course1 := database.Course{
		Name:    "Math",
		Capital: 50,
		Enroll:  0,
	}
	course2 := database.Course{
		Name:    "English",
		Capital: 50,
		Enroll:  0,
	}
	course3 := database.Course{
		Name:    "Chinese",
		Capital: 50,
		Enroll:  0,
	}
	database.DB.Create(&course1)
	database.DB.Create(&course2)
	database.DB.Create(&course3)
}

// DropCourse 退课
func DropCourse(StudentID, CourseID int) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		//是否存在记录
		var enrollment database.Enrollment
		if err := tx.Where("student_id = ? AND course_id = ?", StudentID, CourseID).
			First(&enrollment).Error; err != nil {
			return errors.New("enrollment Not Found")
		}

		//删除
		if err := tx.Delete(&enrollment).Error; err != nil {
			return errors.New("delete failed")
		}

		//更新人数
		if err := tx.Model(&database.Course{}).
			Where("course_id = ?", CourseID).
			Update("enroll", gorm.Expr("enroll - ?", 1)).Error; err != nil {
			return errors.New("update failed")
		}

		return nil
	})
}

// EnrollCourse 选课
func EnrollCourse(StudentID, CourseID int) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 检查学生是否存在
		var student database.Student
		if err := tx.First(&student, StudentID).Error; err != nil {
			return errors.New("student Not Found")
		}

		// 检查课程是否存在
		var course database.Course
		if err := tx.First(&course, CourseID).Error; err != nil {
			return errors.New("course Not Found")
		}

		// 检查课程是否已满
		if course.Enroll >= course.Capital {
			return errors.New("course is full")
		}

		// 是否重复选择
		var exists int64
		if err := tx.Model(&database.Enrollment{}).
			Where("student_id = ? AND course_id = ?", StudentID, CourseID).
			Count(&exists).Error; err != nil {
			return err
		}
		if exists >= 1 {
			return errors.New("enrollment exists")
		}

		// 创建选课关系
		enrollment := database.Enrollment{
			StudentID: StudentID,
			CourseID:  CourseID,
		}
		if err := tx.Create(&enrollment).Error; err != nil {
			return errors.New("enrollment create failed")
		}

		// 更新选课人数
		if err := tx.Model(&database.Course{}).
			Where("course_id = ?", CourseID).
			Update("enroll", gorm.Expr("enroll + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
}

// CheckEnrollment 获取所有选课信息
func CheckEnrollment() ([]database.Enrollment, error) {
	var enrollment []database.Enrollment
	if err := database.DB.First(&enrollment).Error; err != nil {
		return nil, errors.New("enrollment select failed")
	}
	return enrollment, nil
}

// CheckInfo 获取所有课程&学生信息
func CheckInfo() ([]database.Student, []database.Course, error) {
	var student []database.Student
	var course []database.Course
	if err := database.DB.Find(&student).Error; err != nil {
		return nil, nil, errors.New("student select failed")
	}
	if err := database.DB.Find(&course).Error; err != nil {
		return nil, nil, errors.New("course select failed")
	}
	return student, course, nil
}

// AddStudent 新增学生
func AddStudent(name, grade, class string) error {
	student := database.Student{
		Name:  name,
		Grade: grade,
		Class: class,
	}
	if err := database.DB.Create(&student).Error; err != nil {
		return errors.New("student create failed")
	}
	return nil
}

// AddCourse 新增课程
func AddCourse(name string, capital int) error {
	course := database.Course{
		Name:    name,
		Capital: capital,
	}
	if err := database.DB.Create(&course).Error; err != nil {
		return errors.New("course create failed")
	}
	return nil
}
