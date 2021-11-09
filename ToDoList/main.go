package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func messageHandler(conn *websocket.Conn) {

	for {
		_, p, err := conn.ReadMessage()
		fmt.Println(string(p))
		if err != nil {
			log.Println("ReadMessage報錯", err)
			return
		}

		if string(p) == "Taking selfie" {

		} else if string(p) == "Saving selfie" {

		}
	}
}

func wsEndpoint(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("將protocol做Upgrade時報錯:", err)
		return
	}
	log.Println("客戶端連接")

	messageHandler(ws)
}

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
	r.Static("/ginCvWs", "./")
	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")
		i, err := strconv.Atoi(page)
		if err != nil {
			log.Println("page做Atoi時報錯", err)
			c.Request.URL.Path = "/1"
			r.HandleContext(c)
		}
		var A Todo
		var B Todo
		var C Todo
		var D Todo
		db.First(&A, i*4-3)
		db.First(&B, i*4-2)
		db.First(&C, i*4-1)
		db.First(&D, i*4)

		c.HTML(http.StatusOK, "index.html", gin.H{"one": A.Name, "two": B.Name, "three": C.Name, "four": D.Name})

	})
	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		db.Create(&Todo{gorm.Model{}, name})
		var A Todo
		var B Todo
		var C Todo
		var D Todo
		db.First(&A, 1)
		db.First(&B, 2)
		db.First(&C, 3)
		db.First(&D, 4)
		c.HTML(http.StatusOK, "index.html", gin.H{"one": A.Name, "two": B.Name, "three": C.Name, "four": D.Name})
	})
	r.NoRoute(func(c *gin.Context) {
		var A Todo
		var B Todo
		var C Todo
		var D Todo
		db.First(&A, 1)
		db.First(&B, 2)
		db.First(&C, 3)
		db.First(&D, 4)
		c.HTML(http.StatusOK, "index.html", gin.H{"one": A.Name, "two": B.Name, "three": C.Name, "four": D.Name})
	})

	r.Run(":8080")
}

type Todo struct {
	gorm.Model
	Name string
}
