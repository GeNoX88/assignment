package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("連上mysql報錯:", err)
		return
	}

	db.AutoMigrate(&Todo{})
	// u1 := UserInfo{1, "七米", "男", "蛙泳"}
	// db.Create(&u1)
	// db.First(&u)
	// db.Model(&u).Update("hobby", "雙色球")
	// db.Delete(&u)

	fmt.Println("伺服器開啟")
	r := gin.Default()

	r.LoadHTMLGlob("./*")
	r.Static("/ToDoList", "./")
	r.GET("/:page", func(c *gin.Context) {
		if c.Request.RequestURI == "/favicon.ico" {
			return
		}
		page := c.Param("page")
		fmt.Println("page=", page)

		i, err := strconv.Atoi(page)
		if err != nil {
			log.Println("page做Atoi時報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
		}
		var A Todo
		var B Todo
		var C Todo
		var D Todo
		// db.First(&A, i*4-3)
		// db.First(&B, i*4-2)
		// db.First(&C, i*4-1)
		// db.First(&D, i*4)
		db.Offset(i*4 - 4).First(&A)
		db.Offset(i*4 - 3).First(&B)
		db.Offset(i*4 - 2).First(&C)
		db.Offset(i*4 - 1).First(&D)

		c.HTML(http.StatusOK, "index.html",
			gin.H{"one": A.Name, "two": B.Name, "three": C.Name, "four": D.Name, "page": page})

	})

	r.POST("/", func(c *gin.Context) { // 添加事項
		name := c.PostForm("add")
		fmt.Println("新增", name)
		db.Create(&Todo{gorm.Model{}, name})
		c.Redirect(http.StatusMovedPermanently, "/1")
	})

	// r.POST("/:page", func(c *gin.Context) { //改名
	// 	var json UdJson
	// 	c.ShouldBindJSON(&json)
	// 	fmt.Printf("%+v\n", json)

	// 	var t Todo
	// 	WhichPage, err := strconv.Atoi(json.WhichPage)
	// 	if err != nil {
	// 		log.Panicln("json.WhichPage做Atoi時報錯", err)
	// 	}
	// 	Number, err := strconv.Atoi(json.Number)
	// 	if err != nil {
	// 		log.Panicln("json.Number做Atoi時報錯", err)
	// 	}

	// 	db.Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
	// 	db.Model(&t).Update("name", json.NewName)
	// })

	r.PUT("/:page", func(c *gin.Context) { //改名
		var json UdJson
		c.BindJSON(&json)
		fmt.Printf("%+v\n", json)
		WhichPage, err := strconv.Atoi(json.WhichPage)
		if err != nil {
			log.Println("json.WhichPage做Atoi時報錯", err)
		}
		Number, err := strconv.Atoi(json.Number)
		if err != nil {
			log.Println("json.Number做Atoi時報錯", err)
		}

		var t Todo
		db.Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
		db.Model(&t).Update("name", json.NewName)
	})

	r.DELETE("/:page", func(c *gin.Context) {
		var json UdJson
		c.BindJSON(&json)
		fmt.Printf("%+v\n", json)

		WhichPage, err := strconv.Atoi(json.WhichPage)
		if err != nil {
			log.Println("json.WhichPage做Atoi時報錯", err)
		}
		Number, err := strconv.Atoi(json.Number)
		if err != nil {
			log.Println("json.Number做Atoi時報錯", err)
		}
		var t Todo
		db.Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
		db.Debug().Delete(&t)
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/1")
	})

	r.Run(":8080")
}

type Todo struct {
	gorm.Model
	Name string
}

type UdJson struct {
	WhichPage string
	Number    string
	NewName   string
}
