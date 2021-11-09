package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("連上mysql報錯:", err)
		return
	}

	// db.AutoMigrate(&UserInfo{})
	// u1 := UserInfo{gorm.Model{}, "七米", "男", "蛙泳"}
	// db.Create(&u1)
	// var u UserInfo
	// db.First(&u)
	// fmt.Println(u)
	// db.Model(&u).Update("hobby", "雙色球")
	// fmt.Println(u)
	// db.Delete(&u)
	// fmt.Println(u)
	db.Delete(&Todo{gorm.Model{ID: 2}})

}
