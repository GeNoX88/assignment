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

	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Println("database做AuToMigrate報錯", err)
		return
	}

	fmt.Println("伺服器開啟")
	r := gin.Default()

	r.LoadHTMLGlob("./*")
	r.Static("/ToDoList", "./")
	r.GET("/:page", func(c *gin.Context) { //連到未完成事項頁面
		if c.Request.RequestURI == "/favicon.ico" {
			return
		}
		var count int64
		var completed int64

		if err := db.Debug().Table("todos").Where("completed", false).Count(&count).Error; err != nil {
			log.Println("未完成事項的頁面撈count報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		if err := db.Debug().Table("todos").Where("completed", true).Count(&completed).Error; err != nil {
			log.Println("已完成事項撈數量報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		page := c.Param("page")

		i, err := strconv.Atoi(page)
		if err != nil {
			log.Println("page做Atoi時報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		T := make([]Todo, 4)
		if (count%4 == 0) && count != 0 {
			if err = db.Table("todos").Where("completed", false).Offset(i*4 - 4).Limit(4).Find(&T).Error; err != nil {
				log.Println("未完成事項拿四筆整的資料報錯", err)
				return
			}
		} else if count%4 != 0 {
			if err = db.Table("todos").Where("completed", false).Offset(i*4 - 4).Limit(int(count % 4)).Find(&T).Error; err != nil {
				log.Println("未完成事項拿1~3筆的資料報錯", err)
				return
			}
		}

		c.HTML(http.StatusOK, "index.html",
			gin.H{"one": T[0].Name, "two": T[1].Name, "three": T[2].Name, "four": T[3].Name, "page": page, "count": count, "completed": completed, "cpPage": false})
	})
	r.GET("/completed/:page", func(c *gin.Context) { //連到已完成事項頁面
		if c.Request.RequestURI == "/favicon.ico" {
			return
		}
		var count int64
		var completed int64

		if err := db.Debug().Table("todos").Where("completed", false).Count(&count).Error; err != nil {
			log.Println("已完成事項的頁面撈count報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		if err := db.Debug().Table("todos").Where("completed", true).Count(&completed).Error; err != nil {
			log.Println("已完成事項的頁面撈數量報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		page := c.Param("page")

		i, err := strconv.Atoi(page)
		if err != nil {
			log.Println("page做Atoi時報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		T := make([]Todo, 4)
		if (count%4 == 0) && count != 0 {
			if err = db.Table("todos").Where("completed", true).Offset(i*4 - 4).Limit(4).Find(&T).Error; err != nil {
				log.Println("已完成事項拿四筆整的資料報錯", err)
				return
			}
		} else if count%4 != 0 {
			if err = db.Table("todos").Where("completed", true).Offset(i*4 - 4).Limit(int(count % 4)).Find(&T).Error; err != nil {
				log.Println("已完成事項拿1~3筆的資料報錯", err)
				return
			}
		}

		c.HTML(http.StatusOK, "index.html",
			gin.H{"one": T[0].Name, "two": T[1].Name, "three": T[2].Name, "four": T[3].Name, "page": page, "count": count, "completed": completed, "cpPage": true})
	})
	r.POST("/", func(c *gin.Context) { // 添加事項
		name := c.PostForm("add")
		fmt.Println("新增", name)
		if err := db.Create(&Todo{gorm.Model{}, name, false}).Error; err != nil {
			log.Println("建立事件db.Create時報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/1")
	})

	r.PUT("/:page", func(c *gin.Context) { //改名
		var json UdJson
		if err := c.BindJSON(&json); err != nil {
			log.Println("未完成事項做改名的BindJSON報錯", err)
			return
		}
		fmt.Printf("%+v\n", json)
		WhichPage, err := strconv.Atoi(json.WhichPage)
		if err != nil {
			log.Println("json.WhichPage做Atoi時報錯", err)
			return
		}
		Number, err := strconv.Atoi(json.Number)
		if err != nil {
			log.Println("json.Number做Atoi時報錯", err)
			return
		}

		var t Todo
		db.Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
		db.Model(&t).Update("name", json.NewName)
	})

	r.PUT("/changeState/:page", func(c *gin.Context) { //改狀態
		var json UdJson
		if err := c.BindJSON(&json); err != nil {
			log.Println("事項切換完成狀態的BindJSON報錯", err)
			return
		}
		fmt.Printf("%+v\n", json)
		WhichPage, err := strconv.Atoi(json.WhichPage)
		if err != nil {
			log.Println("json.WhichPage做Atoi時報錯", err)
			return
		}
		Number, err := strconv.Atoi(json.Number)
		if err != nil {
			log.Println("json.Number做Atoi時報錯", err)
			return
		}

		var cpPage bool
		if json.cpPage == "false" {
			cpPage = false
		} else if json.cpPage == "true" {
			cpPage = true
		}
		var t Todo
		db.Where("completed", cpPage).Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
		db.Model(&t).Update("completed", !cpPage)
	})

	r.DELETE("/:page", func(c *gin.Context) { //刪除項目
		var json UdJson
		if err := c.BindJSON(&json); err != nil {
			log.Println("刪除未完成項目的路由中BindJson報錯", err)
		}
		fmt.Printf("%#v\n", json)

		WhichPage, err := strconv.Atoi(json.WhichPage)
		if err != nil {
			log.Println("json.WhichPage做Atoi時報錯", err)
		}
		Number, err := strconv.Atoi(json.Number)
		if err != nil {
			log.Println("json.Number做Atoi時報錯", err)
		}

		var cpPage bool
		if json.cpPage == "false" {
			cpPage = false
		} else if json.cpPage == "true" {
			cpPage = true
		}

		var t Todo
		db.Where("completed", cpPage).Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
		db.Debug().Delete(&t)
	})

	r.NoRoute(func(c *gin.Context) {
		fmt.Println("NoRoute已作動")
		c.Redirect(http.StatusMovedPermanently, "/1")
	})

	if err := r.Run(":8080"); err != nil {
		log.Println(`r.Run(":8080")報錯`, err)
		return
	}
}

type Todo struct {
	gorm.Model
	Name      string
	Completed bool
}

type UdJson struct {
	WhichPage string
	Number    string
	NewName   string
	cpPage    string
}
