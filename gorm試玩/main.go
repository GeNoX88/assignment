package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string
}

func main() {
	// // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "root:root1234@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Println("連上mysql報錯:", err)
	// 	return
	// }

	// // db.AutoMigrate(&UserInfo{})
	// // u1 := UserInfo{gorm.Model{}, "七米", "男", "蛙泳"}
	// // db.Create(&u1)
	// // var u UserInfo
	// // db.First(&u)
	// // fmt.Println(u)
	// // db.Model(&u).Update("hobby", "雙色球")
	// // fmt.Println(u)
	// // db.Delete(&u)
	// // fmt.Println(u)
	// db.Delete(&Todo{gorm.Model{ID: 2}})
	m := new(map[string]int)

	// m := map[string]int{}
	*m = map[string]int{
		"hi": 8,
	}
	// people := new(map[string]string)
	// fmt.Printf("%+v", people)
	// (*people)["hi"] := "string"
	fmt.Printf("%+v", m)
	// p := *people
	// p["name"] = "Kalan" // panic: assignment to entry in nil map
}
