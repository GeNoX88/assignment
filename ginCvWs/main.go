package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gocv.io/x/gocv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// var t = template.Must(template.ParseFiles("index.html"))

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	if err := t.ExecuteTemplate(w, "index.html", nil); err != nil {
// 		log.Println(err)
// 	}
// }

func messageHandler(conn *websocket.Conn) {
	Img := []byte{}
	for {
		_, p, err := conn.ReadMessage()
		fmt.Println(string(p))
		if err != nil {
			log.Println("ReadMessage報錯", err)
			return
		}

		if string(p) == "Taking selfie" {
			webcam, err := gocv.OpenVideoCapture(0)
			time.Sleep(1 * time.Second)
			if err != nil {
				fmt.Printf("Error opening video capture device: %v\n", 0)
				return
			}
			img := gocv.NewMat()

			if ok := webcam.Read(&img); !ok {
				fmt.Printf("Device closed: %v\n", 0)
				return
			}
			if err := webcam.Close(); err != nil {
				log.Println("webcam.Close報錯", err)
				return
			}
			if img.Empty() {
				fmt.Println("img is empty")
				return
			}

			fmt.Println(img.Type()) //CV8UC3
			buf, err := gocv.IMEncode(".jpg", img)
			if err != nil {
				log.Println("img編碼成buf時報錯", err)
				return
			}
			if err := img.Close(); err != nil {
				log.Println("img.Close報錯", err)
				return
			}
			Img = buf.GetBytes()
			if err := conn.WriteMessage(websocket.BinaryMessage, buf.GetBytes()); err != nil {
				if err != nil {
					log.Println("WriteMessage報錯", err)
					return
				}
			}
			buf.Close()

		} else if string(p) == "Saving selfie" {
			fmt.Println("要存照片囉！！")

			err = os.WriteFile("self.jpg", Img, os.ModePerm)
			if err != nil {
				log.Println("將Img存成jpg檔時報錯", err)
				return
			}
		}
	}
}

func wsEndpoint(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("將request做upgrade時報錯", err)
		return
	}
	log.Println("客戶端連接")

	messageHandler(ws)
}

func main() {
	fmt.Println("伺服器開啟")
	r := gin.Default()

	r.LoadHTMLGlob("./public/*")
	r.Static("/ginCvWs", "./")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/ws", wsEndpoint)
	r.Run(":8080")
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/ws", wsEndpoint)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
