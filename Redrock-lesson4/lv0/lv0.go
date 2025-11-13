package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	UserID    int       `gorm:"primaryKey;column:user_id;autoIncrement"`       //主键，自增
	Username  string    `gorm:"uniqueIndex;column:username;type:varchar(255)"` //唯一索引
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"uniqueIndex;column:email;type:varchar(255)"` //唯一索引
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "lv0_users"
}

func main() {
	//dsn，Data Source Name，翻译过来叫数据库源名称。DSN 定义了一个数据库的连接信息，包含用户名、密码、数据库 IP、数据库端口、数据库字符集、数据库时区等信息
	dsn := "claran:chr070309@tcp(localhost:3306)/Redrock_lesson3?charset=utf8&parseTime=True&loc=Local"

	//连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示日志
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//自动迁移，让gorm自动创建users表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed AutoMigrated", err)
	}

	// 增
	db.Create(&User{
		Username: "claran",
		Password: "123456",
		Email:    "big_fell_sans@163.com",
	})

	db.Create(&User{
		Username: "visitor_1",
		Password: "666666",
		Email:    "visitor_1@email.com",
	})

	db.Create(&User{
		Username: "visitor_1",
		Password: "666666",
		Email:    "visitor_1@email.com",
	})

	//查
	var user User
	db.Where("name = ?", "claran").First(&user)

	var weakPasswordUser []User
	db.Where("Password = ?", "666666").Find(&weakPasswordUser)

	//改
	db.Model(&User{}).Where("password = ?", "666666").Update("password", "123456")

	//删
	db.Where("name = ?", "claran").
		Where("password = ?", "123456").
		Where("email = ?", "big_fell_sans@163.com").
		Delete(&User{})
}
