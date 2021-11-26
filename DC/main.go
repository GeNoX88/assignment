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
	dsn := "root:root1234@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm打開mysql資料庫時報錯:", err)
		return
	}

	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Println("database做AuToMigrate報錯", err)
		return
	}
	r := gin.Default()
	r.LoadHTMLGlob("index.html")
	r.Static("/DC", "./")

	r.GET("/:cpPage/:page/:record", func(c *gin.Context) { //連到已完成事項頁面
		var count int64     //未完成數
		var completed int64 //已完成數

		var cpPage string = c.Param("cpPage")
		record, err := strconv.Atoi(c.Param("record"))
		if err != nil {
			log.Println("", err)
			return
		}
		if err := db.Debug().Table("todos").Where("completed", false).Where("deleted_at", nil).Count(&count).Error; err != nil {
			log.Println("未完成事項的頁面撈未完成數count報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/all/1/4")
			return
		}

		if err := db.Debug().Table("todos").Where("completed", true).Where("deleted_at", nil).Count(&completed).Error; err != nil {
			log.Println("未完成事項頁面 撈已完成事項數量報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/all/1/4")
			return
		}

		i, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			log.Println("page做Atoi時報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/all/1/4")
			return
		}
		T := []Todo{}
		if cpPage == "all" {
			if err = db.Table("todos").Offset(record * (i - 1)).Limit(record).Find(&T).Error; err != nil {
				log.Println("GET路由dB拿all資料報錯", err)
				c.Redirect(http.StatusMovedPermanently, "/all/1/4")
				return
			}
		} else {
			cpBool := cpPage == "cp"
			if err = db.Table("todos").Where("completed", cpBool).Offset(record * (i - 1)).Limit(record).Find(&T).Error; err != nil {
				log.Println("GET路由dB拿ncp或cp資料報錯", err)
				c.Redirect(http.StatusMovedPermanently, "/all/1/4")
				return
			}
		}
		c.HTML(http.StatusOK, "index.html",
			gin.H{"T": T, "page": i, "count": count, "completed": completed, "cpPage": cpPage, "record": record})
	})
	r.POST("/add", func(c *gin.Context) { // 添加事項
		name := c.PostForm("add")

		fmt.Println("新增:", name)
		if name == "" {
			return
		}
		if err := db.Create(&Todo{gorm.Model{}, name, false}).Error; err != nil {
			log.Println("建立事件db.Create時報錯", err)
			return
		}
	})
	r.PUT("/changeName", func(c *gin.Context) { //改名
		type J struct {
			Id      uint
			NewName string
		}
		var json J
		if err := c.BindJSON(&json); err != nil {
			log.Println("事項做改名的BindJSON報錯:", err)
			return
		}
		fmt.Printf("改名收到了json:%+v\n", json)
		if json.NewName == "" {
			return
		}
		var t Todo
		if err := db.First(&t, json.Id).Error; err != nil {
			log.Println("dB找出要改名之項目的First函數報錯:", err)
			return
		}
		if err := db.Model(&t).Update("name", json.NewName).Error; err != nil {
			log.Println("dB做項目改名的Update函數報錯:", err)
			return
		}
	})

	r.PUT("/changeState", func(c *gin.Context) { //改狀態
		var json Todo
		if err := c.BindJSON(&json); err != nil {
			log.Println("事項切換完成狀態的BindJSON報錯:", err)
			return
		}
		fmt.Printf("切換完成狀態 json:%+v\n", json)

		var t Todo
		if err := db.First(&t, json.Model.ID).Error; err != nil {
			log.Println("dB找出要改哪一筆事件的完成狀態時First函數報錯:", err)
			return
		}
		if err := db.Model(&t).Update("completed", !json.Completed).Error; err != nil {
			log.Println("dB改資料的完成狀態時Update函數報錯:", err)
			return
		}
	})

	r.DELETE("/deleteTodo", func(c *gin.Context) { //刪除項目
		var json Todo
		if err := c.BindJSON(&json); err != nil {
			log.Println("項目刪除的路由中BindJson報錯", err)
			return
		}
		fmt.Printf("事件刪除路由json:%+v\n", json)

		var t Todo
		if err := db.Debug().First(&t, json.Model.ID).Error; err != nil {
			log.Println("dB找出要刪除的todo時報錯:", err)
			return
		}
		if err := db.Debug().Delete(&t).Error; err != nil {
			log.Println("dB刪除todo時報錯:", err)
			return
		}
	})

	r.NoRoute(func(c *gin.Context) {
		fmt.Println("NoRoute已作動   重定向至 /all/1/4")
		c.Redirect(http.StatusMovedPermanently, "/all/1/4")
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
