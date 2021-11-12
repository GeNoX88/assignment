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

	// r.GET("/:page", func(c *gin.Context) { //連到未完成事項頁面
	// 	var count int64     //未完成數
	// 	var completed int64 //已完成數

	// 	if err := db.Debug().Table("todos").Where("completed", false).Where("deleted_at", nil).Count(&count).Error; err != nil {
	// 		log.Println("未完成事項的頁面撈未完成數count報錯", err)
	// 		c.Redirect(http.StatusMovedPermanently, "/1")
	// 		return
	// 	}

	// 	if err := db.Debug().Table("todos").Where("completed", true).Where("deleted_at", nil).Count(&completed).Error; err != nil {
	// 		log.Println("未完成事項頁面 撈已完成事項數量報錯", err)
	// 		c.Redirect(http.StatusMovedPermanently, "/1")
	// 		return
	// 	}

	// 	page := c.Param("page")

	// 	i, err := strconv.Atoi(page)
	// 	if err != nil {
	// 		log.Println("page做Atoi時報錯", err)
	// 		c.Redirect(http.StatusMovedPermanently, "/1")
	// 		return
	// 	}
	// 	T := make([]Todo, 4)
	// 	fmt.Printf("T: %v\n", T)
	// 	if count%4 != 0 && i > int(count/4) {
	// 		// T2 := []Todo{}
	// 		if err = db.Table("todos").Where("completed", false).Offset(i*4 - 4).Limit(4).Find(&T).Error; err != nil {
	// 			log.Println("未完成事項拿1~3筆的資料報錯", err)
	// 			return
	// 		}

	// 		// for c, v := range T {
	// 		// 	fmt.Printf("v: %d = %v\n", c, v)
	// 		// }
	// 		// for c, v := range T2 {
	// 		// 	fmt.Printf("T2: %d = %v\n", c, v)
	// 		// }
	// 		// copy(T, T2)
	// 		// for c, v := range T {
	// 		// 	fmt.Printf("T1: %d = %v\n", c, v)
	// 		// }
	// 	} else {
	// 		if err = db.Table("todos").Where("completed", false).Offset(i*4 - 4).Limit(4).Find(&T).Error; err != nil {
	// 			log.Println("未完成事項拿四筆整的資料報錯", err)
	// 			return
	// 		}
	// 	}
	// 	//

	// 	c.HTML(http.StatusOK, "index.html",
	// 		gin.H{"T": T, "page": page, "count": count, "completed": completed, "cpPage": false})
	// })
	r.GET("/:cpPage/:page/:record", func(c *gin.Context) { //連到已完成事項頁面
		var count int64     //未完成數
		var completed int64 //已完成數

		var cpPage bool = c.Param("cpPage") == "cp"
		record, err := strconv.Atoi(c.Param("record"))
		if err != nil {
			log.Println("", err)
			return
		}
		if err := db.Debug().Table("todos").Where("completed", false).Where("deleted_at", nil).Count(&count).Error; err != nil {
			log.Println("未完成事項的頁面撈未完成數count報錯", err)
			c.Redirect(http.StatusMovedPermanently, "/1")
			return
		}

		if err := db.Debug().Table("todos").Where("completed", true).Where("deleted_at", nil).Count(&completed).Error; err != nil {
			log.Println("未完成事項頁面 撈已完成事項數量報錯", err)
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
		T := []Todo{}
		fmt.Printf("T: %v\n", T)

		if err = db.Table("todos").Where("completed", cpPage).Offset(record * (i - 1)).Limit(record).Find(&T).Error; err != nil {
			log.Println("getHandler中拿資料報錯", err)
			return
		}

		c.HTML(http.StatusOK, "index.tmpl",
			gin.H{"T": T, "page": i, "count": count, "completed": completed, "cpPage": cpPage, "record": record})
	})
	r.POST("/add", func(c *gin.Context) { // 添加事項
		var t Todo
		if err := c.BindJSON(&t); err != nil {
			log.Println("新增事項的BindJSON報錯", err)
			return
		}
		fmt.Printf("%+v", t)
		fmt.Println("新增:", t.Name)
		if err := db.Create(&Todo{gorm.Model{}, t.Name, false}).Error; err != nil {
			log.Println("建立事件db.Create時報錯", err)
			return
		}
	})
	r.PUT("/changeName", func(c *gin.Context) { //改名
		var json UdJson
		if err := c.BindJSON(&json); err != nil {
			log.Println("事項做改名的BindJSON報錯", err)
			return
		}
		fmt.Printf("%+v\n", json)
		// WhichPage, err := strconv.Atoi(json.WhichPage)
		// if err != nil {
		// 	log.Println("json.WhichPage做Atoi時報錯", err)
		// 	return
		// }
		// Number, err := strconv.Atoi(json.Number)
		// if err != nil {
		// 	log.Println("json.Number做Atoi時報錯", err)
		// 	return
		// }
		// var cpPage bool = json.CpPage == "true"
		// fmt.Println(cpPage)

		var t Todo
		db.First(&t, json.Id)
		if t.Name != "" {
			db.Model(&t).Update("name", json.NewName)
		}
	})

	r.PUT("/changeState", func(c *gin.Context) { //改狀態
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
		fmt.Println(cpPage)
		if json.CpPage == "false" {
			cpPage = false
			fmt.Println(cpPage)
		} else if json.CpPage == "true" {
			cpPage = true
			fmt.Println(cpPage)
		}
		fmt.Println(cpPage)
		var t Todo
		if err := db.Where("completed", cpPage).Offset(((WhichPage)-1)*4 + Number - 1).First(&t).Error; err != nil {
			log.Println("dB找出要改哪一筆資料的完成狀態時報錯", err)
			return
		}
		if err := db.Model(&t).Update("completed", !cpPage).Error; err != nil {
			log.Println("dB改資料的完成狀態時報錯", err)
			return
		}
	})

	r.DELETE("/", func(c *gin.Context) { //刪除項目
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
		if json.CpPage == "false" {
			cpPage = false
		} else if json.CpPage == "true" {
			cpPage = true
		}

		var t Todo
		db.Where("completed", cpPage).Offset(((WhichPage)-1)*4 + Number - 1).First(&t)
		db.Debug().Delete(&t)
	})

	r.NoRoute(func(c *gin.Context) {
		fmt.Println("NoRoute已作動")
		c.Redirect(http.StatusMovedPermanently, "/ncp/1/4")
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
	CpPage    string
	Id        uint
	Record    uint
}
